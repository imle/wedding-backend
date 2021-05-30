package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
)

// EventRSVP holds the schema definition for the EventRSVP entity.
type EventRSVP struct {
	ent.Schema
}

// Fields of the EventRSVP.
func (EventRSVP) Fields() []ent.Field {
	return nil
}

// Edges of the EventRSVP.
func (EventRSVP) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("event", Event.Type).
			Ref("rsvps").
			Unique(),
		edge.From("invitee", Invitee.Type).
			Ref("events").
			Unique(),
	}
}

// Annotations of the EventRSVP.
func (EventRSVP) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "event_rsvps"},
	}
}
