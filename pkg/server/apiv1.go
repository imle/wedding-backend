package server

import (
	"github.com/gin-gonic/gin"

	"wedding/ent"
	"wedding/ent/invitee"
	"wedding/ent/inviteeparty"
)

type APIv1 struct {
	database *ent.Client
	router   *gin.RouterGroup
}

func RegisterAPIv1(database *ent.Client, g *gin.RouterGroup) *APIv1 {
	api := &APIv1{
		database: database,
		router:   g,
	}

	g.GET("", api.QueryByInviteeForParty)
	g.GET("/:code", api.GetInviteeByCode)

	return api
}

func (api *APIv1) QueryByInviteeForParty(c *gin.Context) {
	name := c.Query("query")

	if len(name) < 3 {
		c.JSON(200, gin.H{
			"error": "Unable to find your invite. Please try again or contact the Bride and Groom.",
		})
		return
	}

	// No need to return the related objects cause we only care about the code here.
	matches := api.database.Invitee.Query().
		Where(invitee.NameContainsFold(name)).
		QueryParty().
		WithInvitees().
		AllX(c)

	c.JSON(200, gin.H{
		"matches": matches,
	})
}

func (api *APIv1) GetInviteeByCode(c *gin.Context) {
	code := c.Param("code")

	result, _ := api.database.InviteeParty.Query().
		Where(inviteeparty.Code(code)).
		WithInvitees().
		Only(c)

	if result == nil {
		c.Status(404)
		return
	}

	c.JSON(200, gin.H{
		"party": result,
	})
}
