package schema

import (
	"time"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").Nillable().Optional().Default(time.Now),
		field.Time("updated_at").Nillable().Optional(),
		field.Text("meta_data"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("subscriber", Subscriber.Type).
			Unique(),
		edge.To("useraccount", UserAccount.Type).
			Unique(),
	}
}
