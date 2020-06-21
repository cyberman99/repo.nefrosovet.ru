package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
)

// TODO: relations

// WebinarUser holds the schema definition for the WebinarUser entity.
type WebinarUser struct {
	ent.Schema
}

// Fields of the WebinarUser.
func (WebinarUser) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id"),
		field.Int("webinar_id"),
		field.Enum("status").Nillable().Optional().Values(
			"WAIT",
			"OFFLINE", "ONLINE",
			"BANNED",
		),
		field.Int("medooze_id").Nillable().Optional(),
		field.Int("old_medooze_id").Nillable().Optional(),
		field.Int16("mic").Default(0),   // should be int8, but crashes on migration
		field.Int16("sound").Default(1), // should be int8, but crashes on migration
	}
}

// Edges of the WebinarUser.
func (WebinarUser) Edges() []ent.Edge {
	return nil
}
