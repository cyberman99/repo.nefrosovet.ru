package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Subscriber holds the schema definition for the Subscriber entity.
type Subscriber struct {
	ent.Schema
}

// Fields of the Subscriber.
func (Subscriber) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").MaxLen(64).Default(""),
		field.String("domain").MaxLen(64).Default(""),
		field.String("ha1").MaxLen(64).Default(""),
		field.String("ha1b").MaxLen(64).Default(""),
	}
}

// Edges of the Subscriber.
func (Subscriber) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("subscriber").
			Unique().
			Required(),
	}
}
