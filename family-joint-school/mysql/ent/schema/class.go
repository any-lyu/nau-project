package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"time"
)

// Class holds the schema definition for the Class entity.
type Class struct {
	ent.Schema
}

// Fields of the Class.
func (Class) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.String("name").Optional(),
		field.String("extra").Optional(),
		field.Time("created_at").Optional().Default(time.Now),
		field.Time("updated_at").Optional().Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Class.
func (Class) Edges() []ent.Edge {
	return nil
}

// Annotations of the User
func (Class) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "class"},
	}
}
