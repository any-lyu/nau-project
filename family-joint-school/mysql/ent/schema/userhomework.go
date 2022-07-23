package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"family-joint-school/model"
	"time"
)

// UserHomeWork holds the schema definition for the UserHomeWork entity.
type UserHomeWork struct {
	ent.Schema
}

// Fields of the UserHomeWork.
func (UserHomeWork) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id"),
		field.Uint64("homework_id").Optional(),
		field.Uint64("user_id").Optional(),
		field.String("title").Optional(),
		field.String("content").Optional(),
		field.JSON("pics", []model.Pic{}).Optional(),
		field.String("extra").Optional(),
		field.Time("created_at").Optional().Default(time.Now),
		field.Time("updated_at").Optional().Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the UserHomeWork.
func (UserHomeWork) Edges() []ent.Edge {
	return nil
}

// Annotations of the User
func (UserHomeWork) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "stu_homework"},
	}
}
