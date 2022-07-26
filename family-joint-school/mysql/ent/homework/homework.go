// Code generated by ent, DO NOT EDIT.

package homework

import (
	"time"
)

const (
	// Label holds the string label denoting the homework type in the database.
	Label = "homework"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldContent holds the string denoting the content field in the database.
	FieldContent = "content"
	// FieldPics holds the string denoting the pics field in the database.
	FieldPics = "pics"
	// FieldClassID holds the string denoting the class_id field in the database.
	FieldClassID = "class_id"
	// FieldClassName holds the string denoting the class_name field in the database.
	FieldClassName = "class_name"
	// FieldTeacherID holds the string denoting the teacher_id field in the database.
	FieldTeacherID = "teacher_id"
	// FieldExtra holds the string denoting the extra field in the database.
	FieldExtra = "extra"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// Table holds the table name of the homework in the database.
	Table = "homework"
)

// Columns holds all SQL columns for homework fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldContent,
	FieldPics,
	FieldClassID,
	FieldClassName,
	FieldTeacherID,
	FieldExtra,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)
