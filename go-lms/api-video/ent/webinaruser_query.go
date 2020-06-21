// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"repo.nefrosovet.ru/go-lms/api-video/ent/predicate"
	"repo.nefrosovet.ru/go-lms/api-video/ent/webinaruser"
)

// WebinarUserQuery is the builder for querying WebinarUser entities.
type WebinarUserQuery struct {
	config
	limit      *int
	offset     *int
	order      []Order
	unique     []string
	predicates []predicate.WebinarUser
	// intermediate query.
	sql *sql.Selector
}

// Where adds a new predicate for the builder.
func (wuq *WebinarUserQuery) Where(ps ...predicate.WebinarUser) *WebinarUserQuery {
	wuq.predicates = append(wuq.predicates, ps...)
	return wuq
}

// Limit adds a limit step to the query.
func (wuq *WebinarUserQuery) Limit(limit int) *WebinarUserQuery {
	wuq.limit = &limit
	return wuq
}

// Offset adds an offset step to the query.
func (wuq *WebinarUserQuery) Offset(offset int) *WebinarUserQuery {
	wuq.offset = &offset
	return wuq
}

// Order adds an order step to the query.
func (wuq *WebinarUserQuery) Order(o ...Order) *WebinarUserQuery {
	wuq.order = append(wuq.order, o...)
	return wuq
}

// First returns the first WebinarUser entity in the query. Returns *NotFoundError when no webinaruser was found.
func (wuq *WebinarUserQuery) First(ctx context.Context) (*WebinarUser, error) {
	wus, err := wuq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(wus) == 0 {
		return nil, &NotFoundError{webinaruser.Label}
	}
	return wus[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (wuq *WebinarUserQuery) FirstX(ctx context.Context) *WebinarUser {
	wu, err := wuq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return wu
}

// FirstID returns the first WebinarUser id in the query. Returns *NotFoundError when no id was found.
func (wuq *WebinarUserQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = wuq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{webinaruser.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (wuq *WebinarUserQuery) FirstXID(ctx context.Context) int {
	id, err := wuq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only WebinarUser entity in the query, returns an error if not exactly one entity was returned.
func (wuq *WebinarUserQuery) Only(ctx context.Context) (*WebinarUser, error) {
	wus, err := wuq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(wus) {
	case 1:
		return wus[0], nil
	case 0:
		return nil, &NotFoundError{webinaruser.Label}
	default:
		return nil, &NotSingularError{webinaruser.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (wuq *WebinarUserQuery) OnlyX(ctx context.Context) *WebinarUser {
	wu, err := wuq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return wu
}

// OnlyID returns the only WebinarUser id in the query, returns an error if not exactly one id was returned.
func (wuq *WebinarUserQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = wuq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{webinaruser.Label}
	default:
		err = &NotSingularError{webinaruser.Label}
	}
	return
}

// OnlyXID is like OnlyID, but panics if an error occurs.
func (wuq *WebinarUserQuery) OnlyXID(ctx context.Context) int {
	id, err := wuq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of WebinarUsers.
func (wuq *WebinarUserQuery) All(ctx context.Context) ([]*WebinarUser, error) {
	return wuq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (wuq *WebinarUserQuery) AllX(ctx context.Context) []*WebinarUser {
	wus, err := wuq.All(ctx)
	if err != nil {
		panic(err)
	}
	return wus
}

// IDs executes the query and returns a list of WebinarUser ids.
func (wuq *WebinarUserQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := wuq.Select(webinaruser.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (wuq *WebinarUserQuery) IDsX(ctx context.Context) []int {
	ids, err := wuq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (wuq *WebinarUserQuery) Count(ctx context.Context) (int, error) {
	return wuq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (wuq *WebinarUserQuery) CountX(ctx context.Context) int {
	count, err := wuq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (wuq *WebinarUserQuery) Exist(ctx context.Context) (bool, error) {
	return wuq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (wuq *WebinarUserQuery) ExistX(ctx context.Context) bool {
	exist, err := wuq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (wuq *WebinarUserQuery) Clone() *WebinarUserQuery {
	return &WebinarUserQuery{
		config:     wuq.config,
		limit:      wuq.limit,
		offset:     wuq.offset,
		order:      append([]Order{}, wuq.order...),
		unique:     append([]string{}, wuq.unique...),
		predicates: append([]predicate.WebinarUser{}, wuq.predicates...),
		// clone intermediate query.
		sql: wuq.sql.Clone(),
	}
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		UserID int `json:"user_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.WebinarUser.Query().
//		GroupBy(webinaruser.FieldUserID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (wuq *WebinarUserQuery) GroupBy(field string, fields ...string) *WebinarUserGroupBy {
	group := &WebinarUserGroupBy{config: wuq.config}
	group.fields = append([]string{field}, fields...)
	group.sql = wuq.sqlQuery()
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		UserID int `json:"user_id,omitempty"`
//	}
//
//	client.WebinarUser.Query().
//		Select(webinaruser.FieldUserID).
//		Scan(ctx, &v)
//
func (wuq *WebinarUserQuery) Select(field string, fields ...string) *WebinarUserSelect {
	selector := &WebinarUserSelect{config: wuq.config}
	selector.fields = append([]string{field}, fields...)
	selector.sql = wuq.sqlQuery()
	return selector
}

func (wuq *WebinarUserQuery) sqlAll(ctx context.Context) ([]*WebinarUser, error) {
	var (
		nodes = []*WebinarUser{}
		_spec = wuq.querySpec()
	)
	_spec.ScanValues = func() []interface{} {
		node := &WebinarUser{config: wuq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, wuq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (wuq *WebinarUserQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := wuq.querySpec()
	return sqlgraph.CountNodes(ctx, wuq.driver, _spec)
}

func (wuq *WebinarUserQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := wuq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (wuq *WebinarUserQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   webinaruser.Table,
			Columns: webinaruser.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: webinaruser.FieldID,
			},
		},
		From:   wuq.sql,
		Unique: true,
	}
	if ps := wuq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := wuq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := wuq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := wuq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (wuq *WebinarUserQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(wuq.driver.Dialect())
	t1 := builder.Table(webinaruser.Table)
	selector := builder.Select(t1.Columns(webinaruser.Columns...)...).From(t1)
	if wuq.sql != nil {
		selector = wuq.sql
		selector.Select(selector.Columns(webinaruser.Columns...)...)
	}
	for _, p := range wuq.predicates {
		p(selector)
	}
	for _, p := range wuq.order {
		p(selector)
	}
	if offset := wuq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := wuq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WebinarUserGroupBy is the builder for group-by WebinarUser entities.
type WebinarUserGroupBy struct {
	config
	fields []string
	fns    []Aggregate
	// intermediate query.
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the group-by query.
func (wugb *WebinarUserGroupBy) Aggregate(fns ...Aggregate) *WebinarUserGroupBy {
	wugb.fns = append(wugb.fns, fns...)
	return wugb
}

// Scan applies the group-by query and scan the result into the given value.
func (wugb *WebinarUserGroupBy) Scan(ctx context.Context, v interface{}) error {
	return wugb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (wugb *WebinarUserGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := wugb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (wugb *WebinarUserGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(wugb.fields) > 1 {
		return nil, errors.New("ent: WebinarUserGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := wugb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (wugb *WebinarUserGroupBy) StringsX(ctx context.Context) []string {
	v, err := wugb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (wugb *WebinarUserGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(wugb.fields) > 1 {
		return nil, errors.New("ent: WebinarUserGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := wugb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (wugb *WebinarUserGroupBy) IntsX(ctx context.Context) []int {
	v, err := wugb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (wugb *WebinarUserGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(wugb.fields) > 1 {
		return nil, errors.New("ent: WebinarUserGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := wugb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (wugb *WebinarUserGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := wugb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (wugb *WebinarUserGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(wugb.fields) > 1 {
		return nil, errors.New("ent: WebinarUserGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := wugb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (wugb *WebinarUserGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := wugb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (wugb *WebinarUserGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := wugb.sqlQuery().Query()
	if err := wugb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (wugb *WebinarUserGroupBy) sqlQuery() *sql.Selector {
	selector := wugb.sql
	columns := make([]string, 0, len(wugb.fields)+len(wugb.fns))
	columns = append(columns, wugb.fields...)
	for _, fn := range wugb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(wugb.fields...)
}

// WebinarUserSelect is the builder for select fields of WebinarUser entities.
type WebinarUserSelect struct {
	config
	fields []string
	// intermediate queries.
	sql *sql.Selector
}

// Scan applies the selector query and scan the result into the given value.
func (wus *WebinarUserSelect) Scan(ctx context.Context, v interface{}) error {
	return wus.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (wus *WebinarUserSelect) ScanX(ctx context.Context, v interface{}) {
	if err := wus.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (wus *WebinarUserSelect) Strings(ctx context.Context) ([]string, error) {
	if len(wus.fields) > 1 {
		return nil, errors.New("ent: WebinarUserSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := wus.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (wus *WebinarUserSelect) StringsX(ctx context.Context) []string {
	v, err := wus.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (wus *WebinarUserSelect) Ints(ctx context.Context) ([]int, error) {
	if len(wus.fields) > 1 {
		return nil, errors.New("ent: WebinarUserSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := wus.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (wus *WebinarUserSelect) IntsX(ctx context.Context) []int {
	v, err := wus.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (wus *WebinarUserSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(wus.fields) > 1 {
		return nil, errors.New("ent: WebinarUserSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := wus.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (wus *WebinarUserSelect) Float64sX(ctx context.Context) []float64 {
	v, err := wus.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (wus *WebinarUserSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(wus.fields) > 1 {
		return nil, errors.New("ent: WebinarUserSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := wus.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (wus *WebinarUserSelect) BoolsX(ctx context.Context) []bool {
	v, err := wus.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (wus *WebinarUserSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := wus.sqlQuery().Query()
	if err := wus.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (wus *WebinarUserSelect) sqlQuery() sql.Querier {
	selector := wus.sql
	selector.Select(selector.Columns(wus.fields...)...)
	return selector
}
