package apiv1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	log "github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"

	"wedding/ent"
	"wedding/ent/invitee"
	"wedding/ent/inviteeparty"
)

type RSVP struct {
	database *ent.Client
	tracer   trace.Tracer
}

func NewRSVP(database *ent.Client) *RSVP {
	return &RSVP{
		database: database,
		tracer:   otel.Tracer("paperfree/apiv1"),
	}
}

func (api *RSVP) Register(singular *gin.RouterGroup, plural *gin.RouterGroup) {
	singular.GET("/:code", api.getInviteeByCode)
	plural.GET("", api.queryByInviteeForParty)
	plural.POST("", api.updateInviteeInfos)
}

func (api *RSVP) queryByInviteeForParty(c *gin.Context) {
	ctx, span := api.tracer.Start(c.Request.Context(), "query-by-party")
	defer span.End()

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
		AllX(ctx)

	c.JSON(http.StatusOK, gin.H{
		"matches": matches,
	})
}

func (api *RSVP) getInviteeByCode(c *gin.Context) {
	ctx, span := api.tracer.Start(c.Request.Context(), "get-by-code")
	defer span.End()

	code := c.Param("code")

	result, _ := api.database.InviteeParty.Query().
		Where(inviteeparty.Code(code)).
		WithInvitees().
		Only(ctx)

	if result == nil {
		c.Status(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"party": result,
	})
}

func (api *RSVP) updateInviteeInfos(c *gin.Context) {
	ctx, span := api.tracer.Start(c.Request.Context(), "update-info")
	defer span.End()

	var invitees []*ent.Invitee

	err := c.ShouldBindBodyWith(&invitees, binding.JSON)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	tx, err := api.database.Tx(ctx)
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
				Save(ctx)
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
