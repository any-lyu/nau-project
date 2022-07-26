// Code generated by ent, DO NOT EDIT.

package class

import (
	"family-joint-school/mysql/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Extra applies equality check predicate on the "extra" field. It's identical to ExtraEQ.
func Extra(v string) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExtra), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Class {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Class(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Class {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Class(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameIsNil applies the IsNil predicate on the "name" field.
func NameIsNil() predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldName)))
	})
}

// NameNotNil applies the NotNil predicate on the "name" field.
func NameNotNil() predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldName)))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// ExtraEQ applies the EQ predicate on the "extra" field.
func ExtraEQ(v string) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExtra), v))
	})
}

// ExtraNEQ applies the NEQ predicate on the "extra" field.
func ExtraNEQ(v string) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldExtra), v))
	})
}

// ExtraIn applies the In predicate on the "extra" field.
func ExtraIn(vs ...string) predicate.Class {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Class(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldExtra), v...))
	})
}

// ExtraNotIn applies the NotIn predicate on the "extra" field.
func ExtraNotIn(vs ...string) predicate.Class {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Class(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldExtra), v...))
	})
}

// ExtraGT applies the GT predicate on the "extra" field.
func ExtraGT(v string) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldExtra), v))
	})
}

// ExtraGTE applies the GTE predicate on the "extra" field.
func ExtraGTE(v string) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldExtra), v))
	})
}

// ExtraLT applies the LT predicate on the "extra" field.
func ExtraLT(v string) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldExtra), v))
	})
}

// ExtraLTE applies the LTE predicate on the "extra" field.
func ExtraLTE(v string) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldExtra), v))
	})
}

// ExtraContains applies the Contains predicate on the "extra" field.
func ExtraContains(v string) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldExtra), v))
	})
}

// ExtraHasPrefix applies the HasPrefix predicate on the "extra" field.
func ExtraHasPrefix(v string) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldExtra), v))
	})
}

// ExtraHasSuffix applies the HasSuffix predicate on the "extra" field.
func ExtraHasSuffix(v string) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldExtra), v))
	})
}

// ExtraIsNil applies the IsNil predicate on the "extra" field.
func ExtraIsNil() predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldExtra)))
	})
}

// ExtraNotNil applies the NotNil predicate on the "extra" field.
func ExtraNotNil() predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldExtra)))
	})
}

// ExtraEqualFold applies the EqualFold predicate on the "extra" field.
func ExtraEqualFold(v string) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldExtra), v))
	})
}

// ExtraContainsFold applies the ContainsFold predicate on the "extra" field.
func ExtraContainsFold(v string) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldExtra), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Class {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Class(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Class {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Class(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIsNil applies the IsNil predicate on the "created_at" field.
func CreatedAtIsNil() predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCreatedAt)))
	})
}

// CreatedAtNotNil applies the NotNil predicate on the "created_at" field.
func CreatedAtNotNil() predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCreatedAt)))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Class {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Class(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Class {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Class(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUpdatedAt)))
	})
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUpdatedAt)))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Class) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Class) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Class) predicate.Class {
	return predicate.Class(func(s *sql.Selector) {
		p(s.Not())
	})
}
