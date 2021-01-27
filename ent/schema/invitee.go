package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Invitee holds the schema definition for the Invitee entity.
type Invitee struct {
	ent.Schema
}

// Fields of the Invitee.
func (Invitee) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty(),
		field.Bool("is_child").
			Default(false).
			Optional().Nillable(),
		field.Bool("has_plus_one").
			Default(false),
		field.Bool("is_bridesmaid").
			Default(false),
		field.Bool("is_groomsman").
			Default(false),
		field.String("plus_one_name").
			Optional().Nillable(),
		field.String("phone").
			Optional().Nillable(),
		field.String("email").
			Optional().Nillable(),
		field.String("address_line_1").
			Optional().Nillable(),
		field.String("address_line_2").
			Optional().Nillable(),
		field.String("address_city").
			Optional().Nillable(),
		field.String("address_state").
			Optional().Nillable(),
		field.String("address_postal_code").
			Optional().Nillable(),
		field.String("address_country").
			Optional().Nillable(),
		field.Bool("rsvp_response").
			Default(false),
	}
}

// Edges of the Invitee.
func (Invitee) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("party", InviteeParty.Type).
			Ref("invitees").
			Unique(),
	}
}
