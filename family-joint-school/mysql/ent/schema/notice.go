package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"time"
)

// Notice holds the schema definition for the Notice entity.
type Notice struct {
	ent.Schema
}

// Fields of the Notice.
func (Notice) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.String("title").Optional(),
		field.String("content").Optional(),
		field.Uint8("level").Optional(),
		field.String("extra").Optional(),
		field.Time("created_at").Optional().Default(time.Now),
		field.Time("updated_at").Optional().Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Notice.
func (Notice) Edges() []ent.Edge {
	return nil
}

// Annotations of the Notice
func (Notice) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "notice"},
	}
}
