package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// AccountKey holds the schema definition for the AccountKey entity.
type AccountKey struct {
	ent.Schema
}

// Fields of the AccountKey.
func (AccountKey) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").Nillable().Optional(),
		field.Time("updated_at").Nillable().Optional(),
		field.Int16("account_id"),
		field.String("key").MaxLen(255),
		field.String("options").Nillable().Optional(),
		field.String("meta_data"),
	}
}

// Edges of the AccountKey.
func (AccountKey) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("useraccount", UserAccount.Type).
			Ref("accountkeys").
			Unique(),
	}
}
