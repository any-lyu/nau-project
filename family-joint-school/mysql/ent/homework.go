// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"family-joint-school/model"
	"family-joint-school/mysql/ent/homework"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Homework is the model entity for the Homework schema.
type Homework struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Content holds the value of the "content" field.
	Content string `json:"content,omitempty"`
	// Pics holds the value of the "pics" field.
	Pics []model.Pic `json:"pics,omitempty"`
	// ClassID holds the value of the "class_id" field.
	ClassID uint8 `json:"class_id,omitempty"`
	// ClassName holds the value of the "class_name" field.
	ClassName string `json:"class_name,omitempty"`
	// TeacherID holds the value of the "teacher_id" field.
	TeacherID uint8 `json:"teacher_id,omitempty"`
	// Extra holds the value of the "extra" field.
	Extra string `json:"extra,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Homework) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case homework.FieldPics:
			values[i] = new([]byte)
		case homework.FieldID, homework.FieldClassID, homework.FieldTeacherID:
			values[i] = new(sql.NullInt64)
		case homework.FieldTitle, homework.FieldContent, homework.FieldClassName, homework.FieldExtra:
			values[i] = new(sql.NullString)
		case homework.FieldCreatedAt, homework.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Homework", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Homework fields.
func (h *Homework) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case homework.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			h.ID = uint64(value.Int64)
		case homework.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				h.Title = value.String
			}
		case homework.FieldContent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content", values[i])
			} else if value.Valid {
				h.Content = value.String
			}
		case homework.FieldPics:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field pics", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &h.Pics); err != nil {
					return fmt.Errorf("unmarshal field pics: %w", err)
				}
			}
		case homework.FieldClassID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field class_id", values[i])
			} else if value.Valid {
				h.ClassID = uint8(value.Int64)
			}
		case homework.FieldClassName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field class_name", values[i])
			} else if value.Valid {
				h.ClassName = value.String
			}
		case homework.FieldTeacherID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field teacher_id", values[i])
			} else if value.Valid {
				h.TeacherID = uint8(value.Int64)
			}
		case homework.FieldExtra:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field extra", values[i])
			} else if value.Valid {
				h.Extra = value.String
			}
		case homework.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				h.CreatedAt = value.Time
			}
		case homework.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				h.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Homework.
// Note that you need to call Homework.Unwrap() before calling this method if this Homework
// was returned from a transaction, and the transaction was committed or rolled back.
func (h *Homework) Update() *HomeworkUpdateOne {
	return (&HomeworkClient{config: h.config}).UpdateOne(h)
}

// Unwrap unwraps the Homework entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (h *Homework) Unwrap() *Homework {
	_tx, ok := h.config.driver.(*txDriver)
	if !ok {
		panic("ent: Homework is not a transactional entity")
	}
	h.config.driver = _tx.drv
	return h
}

// String implements the fmt.Stringer.
func (h *Homework) String() string {
	var builder strings.Builder
	builder.WriteString("Homework(")
	builder.WriteString(fmt.Sprintf("id=%v, ", h.ID))
	builder.WriteString("title=")
	builder.WriteString(h.Title)
	builder.WriteString(", ")
	builder.WriteString("content=")
	builder.WriteString(h.Content)
	builder.WriteString(", ")
	builder.WriteString("pics=")
	builder.WriteString(fmt.Sprintf("%v", h.Pics))
	builder.WriteString(", ")
	builder.WriteString("class_id=")
	builder.WriteString(fmt.Sprintf("%v", h.ClassID))
	builder.WriteString(", ")
	builder.WriteString("class_name=")
	builder.WriteString(h.ClassName)
	builder.WriteString(", ")
	builder.WriteString("teacher_id=")
	builder.WriteString(fmt.Sprintf("%v", h.TeacherID))
	builder.WriteString(", ")
	builder.WriteString("extra=")
	builder.WriteString(h.Extra)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(h.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(h.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Homeworks is a parsable slice of Homework.
type Homeworks []*Homework

func (h Homeworks) config(cfg config) {
	for _i := range h {
		h[_i].config = cfg
	}
}
