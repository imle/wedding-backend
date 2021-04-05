package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	log "github.com/sirupsen/logrus"

	"wedding/ent"
	"wedding/ent/invitee"
	"wedding/ent/inviteeparty"
)

type APIv1 struct {
	database *ent.Client
}

func RegisterAPIv1(database *ent.Client, singular *gin.RouterGroup, plural *gin.RouterGroup) *APIv1 {
	api := &APIv1{
		database: database,
	}

	singular.GET("/:code", api.getInviteeByCode)
	plural.GET("", api.queryByInviteeForParty)
	plural.POST("", api.updateInviteeInfos)

	return api
}

func (api *APIv1) queryByInviteeForParty(c *gin.Context) {
	name := c.Query("query")

	if len(name) < 3 {
		c.JSON(http.StatusOK, gin.H{
			"error": "Unable to find your invite. Please try again or contact the Bride and Groom.",
		})
		return
	}

	// No need to return the related objects cause we only care about the code here.
	matches := api.database.Invitee.Query().
		Where(invitee.NameContainsFold(name)).
		QueryParty().
		WithInvitees(func(query *ent.InviteeQuery) {
			query.Order(ent.Asc(invitee.FieldIsChild, invitee.FieldID))
		}).
		AllX(c)

	c.JSON(http.StatusOK, gin.H{
		"matches": matches,
	})
}

func (api *APIv1) getInviteeByCode(c *gin.Context) {
	code := c.Param("code")

	result, _ := api.database.InviteeParty.Query().
		Where(inviteeparty.Code(code)).
		WithInvitees().
		Only(c)

	if result == nil {
		c.Status(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"party": result,
	})
}

func (api *APIv1) updateInviteeInfos(c *gin.Context) {
	var invitees []*ent.Invitee

	err := c.ShouldBindBodyWith(&invitees, binding.JSON)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	tx, err := api.database.Tx(c)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		log.Error(err)
		return
	}

	for _, e := range invitees {
		if e.ID != 0 {
			_, err = tx.Invitee.Update().
				Where(invitee.ID(e.ID)).
				SetNillablePlusOneName(e.PlusOneName).
				SetNillablePhone(e.Phone).
				SetNillableEmail(e.Email).
				SetNillableAddressLine1(e.AddressLine1).
				SetNillableAddressLine2(e.AddressLine2).
				SetNillableAddressCity(e.AddressCity).
				SetNillableAddressState(e.AddressState).
				SetNillableAddressPostalCode(e.AddressPostalCode).
				SetNillableAddressCountry(e.AddressCountry).
				SetNillableRsvpResponse(e.RsvpResponse).
				Save(c)
			if err != nil {
				c.Status(http.StatusInternalServerError)
				log.Error(err)
				return
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		c.Status(http.StatusInternalServerError)
		log.Error(err)
		return
	}

	c.Status(http.StatusOK)
}
