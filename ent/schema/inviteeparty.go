package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// InviteeParty holds the schema definition for the InviteeParty entity.
type InviteeParty struct {
	ent.Schema
}

// Fields of the InviteeParty.
func (InviteeParty) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty(),
		field.String("code").
			Unique().
			MaxLen(10).
			NotEmpty(),
	}
}

// Edges of the InviteeParty.
func (InviteeParty) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("invitees", Invitee.Type),
	}
}
