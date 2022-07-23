package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"family-joint-school/model"
	"time"
)

// Homework holds the schema definition for the Homework entity.
type Homework struct {
	ent.Schema
}

// Fields of the Homework.
func (Homework) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.String("title").Optional(),
		field.String("content").Optional(),
		field.JSON("pics", []model.Pic{}).Optional(),
		field.Uint8("class_id").Optional(),
		field.String("class_name").Optional(),
		field.Uint8("teacher_id").Optional(),
		field.String("extra").Optional(),
		field.Time("created_at").Optional().Default(time.Now),
		field.Time("updated_at").Optional().Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Homework.
func (Homework) Edges() []ent.Edge {
	return nil
}

// Annotations of the Notice
func (Homework) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "homework"},
	}
}
