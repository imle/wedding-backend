package server

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"wedding/ent"
)

type AdminAPIv1 struct {
	database *ent.Client
	router   *gin.RouterGroup
}

func RegisterAdminAPIv1(database *ent.Client, g *gin.RouterGroup) *AdminAPIv1 {
	api := &AdminAPIv1{
		database: database,
		router:   g,
	}

	g.GET("/parties", api.getAllParties)

	return api
}

func (api *AdminAPIv1) getAllParties(c *gin.Context) {
	result, _ := api.database.InviteeParty.Query().
		WithInvitees().
		All(c)

	c.JSON(http.StatusOK, gin.H{
		"parties": result,
	})
}
