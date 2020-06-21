package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// UserAccount holds the schema definition for the UserAccount entity.
type UserAccount struct {
	ent.Schema
}

// Fields of the Account.
func (UserAccount) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").MaxLen(50).Unique(),
		field.String("password").MaxLen(100),
		field.String("remember_token").MaxLen(100).Nillable().Optional(),
		field.Int16("active").Default(0),
		field.String("event_channel").MaxLen(255).Nillable().Optional(),
		field.String("did_prefix").MaxLen(45).Default("EMPTY"),
		field.Int16("use_kamalio").Default(1),
	}
}

// Edges of the Account.
func (UserAccount) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("useraccount").
			Unique().
			Required(),
		edge.To("accountkeys", AccountKey.Type),
	}
}
