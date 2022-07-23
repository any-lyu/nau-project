package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"time"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.String("nickname").Optional(),
		field.String("avatar").Optional(),
		field.String("mobile").Optional(),
		field.String("extra").Optional(),
		field.Uint8("class_id").Optional(),
		field.Time("created_at").Optional().Default(time.Now),
		field.Time("updated_at").Optional().Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

// Annotations of the User
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "user"},
	}
}
