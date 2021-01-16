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
