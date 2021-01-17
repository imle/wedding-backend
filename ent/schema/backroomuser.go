package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// BackroomUser holds the schema definition for the BackroomUser entity.
type BackroomUser struct {
	ent.Schema
}

// Fields of the BackroomUser.
func (BackroomUser) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").
			Unique().
			NotEmpty().
			MaxLen(20),
		field.String("password").
			MaxLen(64).
			Sensitive().
			StructTag(""),
	}
}

// Edges of the BackroomUser.
func (BackroomUser) Edges() []ent.Edge {
	return nil
}
