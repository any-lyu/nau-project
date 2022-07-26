// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"family-joint-school/mysql/ent/homework"
	"family-joint-school/mysql/ent/predicate"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// HomeworkQuery is the builder for querying Homework entities.
type HomeworkQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Homework
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the HomeworkQuery builder.
func (hq *HomeworkQuery) Where(ps ...predicate.Homework) *HomeworkQuery {
	hq.predicates = append(hq.predicates, ps...)
	return hq
}

// Limit adds a limit step to the query.
func (hq *HomeworkQuery) Limit(limit int) *HomeworkQuery {
	hq.limit = &limit
	return hq
}

// Offset adds an offset step to the query.
func (hq *HomeworkQuery) Offset(offset int) *HomeworkQuery {
	hq.offset = &offset
	return hq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (hq *HomeworkQuery) Unique(unique bool) *HomeworkQuery {
	hq.unique = &unique
	return hq
}

// Order adds an order step to the query.
func (hq *HomeworkQuery) Order(o ...OrderFunc) *HomeworkQuery {
	hq.order = append(hq.order, o...)
	return hq
}

// First returns the first Homework entity from the query.
// Returns a *NotFoundError when no Homework was found.
func (hq *HomeworkQuery) First(ctx context.Context) (*Homework, error) {
	nodes, err := hq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{homework.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (hq *HomeworkQuery) FirstX(ctx context.Context) *Homework {
	node, err := hq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Homework ID from the query.
// Returns a *NotFoundError when no Homework ID was found.
func (hq *HomeworkQuery) FirstID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = hq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{homework.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (hq *HomeworkQuery) FirstIDX(ctx context.Context) uint64 {
	id, err := hq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Homework entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Homework entity is found.
// Returns a *NotFoundError when no Homework entities are found.
func (hq *HomeworkQuery) Only(ctx context.Context) (*Homework, error) {
	nodes, err := hq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{homework.Label}
	default:
		return nil, &NotSingularError{homework.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (hq *HomeworkQuery) OnlyX(ctx context.Context) *Homework {
	node, err := hq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Homework ID in the query.
// Returns a *NotSingularError when more than one Homework ID is found.
// Returns a *NotFoundError when no entities are found.
func (hq *HomeworkQuery) OnlyID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = hq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{homework.Label}
	default:
		err = &NotSingularError{homework.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (hq *HomeworkQuery) OnlyIDX(ctx context.Context) uint64 {
	id, err := hq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Homeworks.
func (hq *HomeworkQuery) All(ctx context.Context) ([]*Homework, error) {
	if err := hq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return hq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (hq *HomeworkQuery) AllX(ctx context.Context) []*Homework {
	nodes, err := hq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Homework IDs.
func (hq *HomeworkQuery) IDs(ctx context.Context) ([]uint64, error) {
	var ids []uint64
	if err := hq.Select(homework.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (hq *HomeworkQuery) IDsX(ctx context.Context) []uint64 {
	ids, err := hq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (hq *HomeworkQuery) Count(ctx context.Context) (int, error) {
	if err := hq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return hq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (hq *HomeworkQuery) CountX(ctx context.Context) int {
	count, err := hq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (hq *HomeworkQuery) Exist(ctx context.Context) (bool, error) {
	if err := hq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return hq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (hq *HomeworkQuery) ExistX(ctx context.Context) bool {
	exist, err := hq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the HomeworkQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (hq *HomeworkQuery) Clone() *HomeworkQuery {
	if hq == nil {
		return nil
	}
	return &HomeworkQuery{
		config:     hq.config,
		limit:      hq.limit,
		offset:     hq.offset,
		order:      append([]OrderFunc{}, hq.order...),
		predicates: append([]predicate.Homework{}, hq.predicates...),
		// clone intermediate query.
		sql:    hq.sql.Clone(),
		path:   hq.path,
		unique: hq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Title string `json:"title,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Homework.Query().
//		GroupBy(homework.FieldTitle).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (hq *HomeworkQuery) GroupBy(field string, fields ...string) *HomeworkGroupBy {
	grbuild := &HomeworkGroupBy{config: hq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := hq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return hq.sqlQuery(ctx), nil
	}
	grbuild.label = homework.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Title string `json:"title,omitempty"`
//	}
//
//	client.Homework.Query().
//		Select(homework.FieldTitle).
//		Scan(ctx, &v)
//
func (hq *HomeworkQuery) Select(fields ...string) *HomeworkSelect {
	hq.fields = append(hq.fields, fields...)
	selbuild := &HomeworkSelect{HomeworkQuery: hq}
	selbuild.label = homework.Label
	selbuild.flds, selbuild.scan = &hq.fields, selbuild.Scan
	return selbuild
}

func (hq *HomeworkQuery) prepareQuery(ctx context.Context) error {
	for _, f := range hq.fields {
		if !homework.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if hq.path != nil {
		prev, err := hq.path(ctx)
		if err != nil {
			return err
		}
		hq.sql = prev
	}
	return nil
}

func (hq *HomeworkQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Homework, error) {
	var (
		nodes = []*Homework{}
		_spec = hq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*Homework).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &Homework{config: hq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, hq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (hq *HomeworkQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := hq.querySpec()
	_spec.Node.Columns = hq.fields
	if len(hq.fields) > 0 {
		_spec.Unique = hq.unique != nil && *hq.unique
	}
	return sqlgraph.CountNodes(ctx, hq.driver, _spec)
}

func (hq *HomeworkQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := hq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (hq *HomeworkQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   homework.Table,
			Columns: homework.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: homework.FieldID,
			},
		},
		From:   hq.sql,
		Unique: true,
	}
	if unique := hq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := hq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, homework.FieldID)
		for i := range fields {
			if fields[i] != homework.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := hq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := hq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := hq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := hq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (hq *HomeworkQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(hq.driver.Dialect())
	t1 := builder.Table(homework.Table)
	columns := hq.fields
	if len(columns) == 0 {
		columns = homework.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if hq.sql != nil {
		selector = hq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if hq.unique != nil && *hq.unique {
		selector.Distinct()
	}
	for _, p := range hq.predicates {
		p(selector)
	}
	for _, p := range hq.order {
		p(selector)
	}
	if offset := hq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := hq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// HomeworkGroupBy is the group-by builder for Homework entities.
type HomeworkGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (hgb *HomeworkGroupBy) Aggregate(fns ...AggregateFunc) *HomeworkGroupBy {
	hgb.fns = append(hgb.fns, fns...)
	return hgb
}

// Scan applies the group-by query and scans the result into the given value.
func (hgb *HomeworkGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := hgb.path(ctx)
	if err != nil {
		return err
	}
	hgb.sql = query
	return hgb.sqlScan(ctx, v)
}

func (hgb *HomeworkGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range hgb.fields {
		if !homework.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := hgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := hgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (hgb *HomeworkGroupBy) sqlQuery() *sql.Selector {
	selector := hgb.sql.Select()
	aggregation := make([]string, 0, len(hgb.fns))
	for _, fn := range hgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(hgb.fields)+len(hgb.fns))
		for _, f := range hgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(hgb.fields...)...)
}

// HomeworkSelect is the builder for selecting fields of Homework entities.
type HomeworkSelect struct {
	*HomeworkQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (hs *HomeworkSelect) Scan(ctx context.Context, v interface{}) error {
	if err := hs.prepareQuery(ctx); err != nil {
		return err
	}
	hs.sql = hs.HomeworkQuery.sqlQuery(ctx)
	return hs.sqlScan(ctx, v)
}

func (hs *HomeworkSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := hs.sql.Query()
	if err := hs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
