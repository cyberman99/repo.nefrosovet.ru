// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"repo.nefrosovet.ru/go-lms/api-video/ent/accountkey"
	"repo.nefrosovet.ru/go-lms/api-video/ent/predicate"
	"repo.nefrosovet.ru/go-lms/api-video/ent/user"
	"repo.nefrosovet.ru/go-lms/api-video/ent/useraccount"
)

// UserAccountQuery is the builder for querying UserAccount entities.
type UserAccountQuery struct {
	config
	limit      *int
	offset     *int
	order      []Order
	unique     []string
	predicates []predicate.UserAccount
	// eager-loading edges.
	withUser        *UserQuery
	withAccountkeys *AccountKeyQuery
	withFKs         bool
	// intermediate query.
	sql *sql.Selector
}

// Where adds a new predicate for the builder.
func (uaq *UserAccountQuery) Where(ps ...predicate.UserAccount) *UserAccountQuery {
	uaq.predicates = append(uaq.predicates, ps...)
	return uaq
}

// Limit adds a limit step to the query.
func (uaq *UserAccountQuery) Limit(limit int) *UserAccountQuery {
	uaq.limit = &limit
	return uaq
}

// Offset adds an offset step to the query.
func (uaq *UserAccountQuery) Offset(offset int) *UserAccountQuery {
	uaq.offset = &offset
	return uaq
}

// Order adds an order step to the query.
func (uaq *UserAccountQuery) Order(o ...Order) *UserAccountQuery {
	uaq.order = append(uaq.order, o...)
	return uaq
}

// QueryUser chains the current query on the user edge.
func (uaq *UserAccountQuery) QueryUser() *UserQuery {
	query := &UserQuery{config: uaq.config}
	step := sqlgraph.NewStep(
		sqlgraph.From(useraccount.Table, useraccount.FieldID, uaq.sqlQuery()),
		sqlgraph.To(user.Table, user.FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, useraccount.UserTable, useraccount.UserColumn),
	)
	query.sql = sqlgraph.SetNeighbors(uaq.driver.Dialect(), step)
	return query
}

// QueryAccountkeys chains the current query on the accountkeys edge.
func (uaq *UserAccountQuery) QueryAccountkeys() *AccountKeyQuery {
	query := &AccountKeyQuery{config: uaq.config}
	step := sqlgraph.NewStep(
		sqlgraph.From(useraccount.Table, useraccount.FieldID, uaq.sqlQuery()),
		sqlgraph.To(accountkey.Table, accountkey.FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, useraccount.AccountkeysTable, useraccount.AccountkeysColumn),
	)
	query.sql = sqlgraph.SetNeighbors(uaq.driver.Dialect(), step)
	return query
}

// First returns the first UserAccount entity in the query. Returns *NotFoundError when no useraccount was found.
func (uaq *UserAccountQuery) First(ctx context.Context) (*UserAccount, error) {
	uas, err := uaq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(uas) == 0 {
		return nil, &NotFoundError{useraccount.Label}
	}
	return uas[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (uaq *UserAccountQuery) FirstX(ctx context.Context) *UserAccount {
	ua, err := uaq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return ua
}

// FirstID returns the first UserAccount id in the query. Returns *NotFoundError when no id was found.
func (uaq *UserAccountQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = uaq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{useraccount.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (uaq *UserAccountQuery) FirstXID(ctx context.Context) int {
	id, err := uaq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only UserAccount entity in the query, returns an error if not exactly one entity was returned.
func (uaq *UserAccountQuery) Only(ctx context.Context) (*UserAccount, error) {
	uas, err := uaq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(uas) {
	case 1:
		return uas[0], nil
	case 0:
		return nil, &NotFoundError{useraccount.Label}
	default:
		return nil, &NotSingularError{useraccount.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (uaq *UserAccountQuery) OnlyX(ctx context.Context) *UserAccount {
	ua, err := uaq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return ua
}

// OnlyID returns the only UserAccount id in the query, returns an error if not exactly one id was returned.
func (uaq *UserAccountQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = uaq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{useraccount.Label}
	default:
		err = &NotSingularError{useraccount.Label}
	}
	return
}

// OnlyXID is like OnlyID, but panics if an error occurs.
func (uaq *UserAccountQuery) OnlyXID(ctx context.Context) int {
	id, err := uaq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UserAccounts.
func (uaq *UserAccountQuery) All(ctx context.Context) ([]*UserAccount, error) {
	return uaq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (uaq *UserAccountQuery) AllX(ctx context.Context) []*UserAccount {
	uas, err := uaq.All(ctx)
	if err != nil {
		panic(err)
	}
	return uas
}

// IDs executes the query and returns a list of UserAccount ids.
func (uaq *UserAccountQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := uaq.Select(useraccount.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (uaq *UserAccountQuery) IDsX(ctx context.Context) []int {
	ids, err := uaq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (uaq *UserAccountQuery) Count(ctx context.Context) (int, error) {
	return uaq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (uaq *UserAccountQuery) CountX(ctx context.Context) int {
	count, err := uaq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (uaq *UserAccountQuery) Exist(ctx context.Context) (bool, error) {
	return uaq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (uaq *UserAccountQuery) ExistX(ctx context.Context) bool {
	exist, err := uaq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (uaq *UserAccountQuery) Clone() *UserAccountQuery {
	return &UserAccountQuery{
		config:     uaq.config,
		limit:      uaq.limit,
		offset:     uaq.offset,
		order:      append([]Order{}, uaq.order...),
		unique:     append([]string{}, uaq.unique...),
		predicates: append([]predicate.UserAccount{}, uaq.predicates...),
		// clone intermediate query.
		sql: uaq.sql.Clone(),
	}
}

//  WithUser tells the query-builder to eager-loads the nodes that are connected to
// the "user" edge. The optional arguments used to configure the query builder of the edge.
func (uaq *UserAccountQuery) WithUser(opts ...func(*UserQuery)) *UserAccountQuery {
	query := &UserQuery{config: uaq.config}
	for _, opt := range opts {
		opt(query)
	}
	uaq.withUser = query
	return uaq
}

//  WithAccountkeys tells the query-builder to eager-loads the nodes that are connected to
// the "accountkeys" edge. The optional arguments used to configure the query builder of the edge.
func (uaq *UserAccountQuery) WithAccountkeys(opts ...func(*AccountKeyQuery)) *UserAccountQuery {
	query := &AccountKeyQuery{config: uaq.config}
	for _, opt := range opts {
		opt(query)
	}
	uaq.withAccountkeys = query
	return uaq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Username string `json:"username,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.UserAccount.Query().
//		GroupBy(useraccount.FieldUsername).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (uaq *UserAccountQuery) GroupBy(field string, fields ...string) *UserAccountGroupBy {
	group := &UserAccountGroupBy{config: uaq.config}
	group.fields = append([]string{field}, fields...)
	group.sql = uaq.sqlQuery()
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		Username string `json:"username,omitempty"`
//	}
//
//	client.UserAccount.Query().
//		Select(useraccount.FieldUsername).
//		Scan(ctx, &v)
//
func (uaq *UserAccountQuery) Select(field string, fields ...string) *UserAccountSelect {
	selector := &UserAccountSelect{config: uaq.config}
	selector.fields = append([]string{field}, fields...)
	selector.sql = uaq.sqlQuery()
	return selector
}

func (uaq *UserAccountQuery) sqlAll(ctx context.Context) ([]*UserAccount, error) {
	var (
		nodes       = []*UserAccount{}
		withFKs     = uaq.withFKs
		_spec       = uaq.querySpec()
		loadedTypes = [2]bool{
			uaq.withUser != nil,
			uaq.withAccountkeys != nil,
		}
	)
	if uaq.withUser != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, useraccount.ForeignKeys...)
	}
	_spec.ScanValues = func() []interface{} {
		node := &UserAccount{config: uaq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		if withFKs {
			values = append(values, node.fkValues()...)
		}
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, uaq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := uaq.withUser; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*UserAccount)
		for i := range nodes {
			if fk := nodes[i].user_useraccount; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(user.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "user_useraccount" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.User = n
			}
		}
	}

	if query := uaq.withAccountkeys; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*UserAccount)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.AccountKey(func(s *sql.Selector) {
			s.Where(sql.InValues(useraccount.AccountkeysColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.user_account_accountkeys
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "user_account_accountkeys" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "user_account_accountkeys" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Accountkeys = append(node.Edges.Accountkeys, n)
		}
	}

	return nodes, nil
}

func (uaq *UserAccountQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := uaq.querySpec()
	return sqlgraph.CountNodes(ctx, uaq.driver, _spec)
}

func (uaq *UserAccountQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := uaq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (uaq *UserAccountQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   useraccount.Table,
			Columns: useraccount.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: useraccount.FieldID,
			},
		},
		From:   uaq.sql,
		Unique: true,
	}
	if ps := uaq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := uaq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := uaq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := uaq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (uaq *UserAccountQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(uaq.driver.Dialect())
	t1 := builder.Table(useraccount.Table)
	selector := builder.Select(t1.Columns(useraccount.Columns...)...).From(t1)
	if uaq.sql != nil {
		selector = uaq.sql
		selector.Select(selector.Columns(useraccount.Columns...)...)
	}
	for _, p := range uaq.predicates {
		p(selector)
	}
	for _, p := range uaq.order {
		p(selector)
	}
	if offset := uaq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := uaq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UserAccountGroupBy is the builder for group-by UserAccount entities.
type UserAccountGroupBy struct {
	config
	fields []string
	fns    []Aggregate
	// intermediate query.
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the group-by query.
func (uagb *UserAccountGroupBy) Aggregate(fns ...Aggregate) *UserAccountGroupBy {
	uagb.fns = append(uagb.fns, fns...)
	return uagb
}

// Scan applies the group-by query and scan the result into the given value.
func (uagb *UserAccountGroupBy) Scan(ctx context.Context, v interface{}) error {
	return uagb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (uagb *UserAccountGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := uagb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (uagb *UserAccountGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(uagb.fields) > 1 {
		return nil, errors.New("ent: UserAccountGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := uagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (uagb *UserAccountGroupBy) StringsX(ctx context.Context) []string {
	v, err := uagb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (uagb *UserAccountGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(uagb.fields) > 1 {
		return nil, errors.New("ent: UserAccountGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := uagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (uagb *UserAccountGroupBy) IntsX(ctx context.Context) []int {
	v, err := uagb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (uagb *UserAccountGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(uagb.fields) > 1 {
		return nil, errors.New("ent: UserAccountGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := uagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (uagb *UserAccountGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := uagb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (uagb *UserAccountGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(uagb.fields) > 1 {
		return nil, errors.New("ent: UserAccountGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := uagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (uagb *UserAccountGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := uagb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (uagb *UserAccountGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := uagb.sqlQuery().Query()
	if err := uagb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (uagb *UserAccountGroupBy) sqlQuery() *sql.Selector {
	selector := uagb.sql
	columns := make([]string, 0, len(uagb.fields)+len(uagb.fns))
	columns = append(columns, uagb.fields...)
	for _, fn := range uagb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(uagb.fields...)
}

// UserAccountSelect is the builder for select fields of UserAccount entities.
type UserAccountSelect struct {
	config
	fields []string
	// intermediate queries.
	sql *sql.Selector
}

// Scan applies the selector query and scan the result into the given value.
func (uas *UserAccountSelect) Scan(ctx context.Context, v interface{}) error {
	return uas.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (uas *UserAccountSelect) ScanX(ctx context.Context, v interface{}) {
	if err := uas.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (uas *UserAccountSelect) Strings(ctx context.Context) ([]string, error) {
	if len(uas.fields) > 1 {
		return nil, errors.New("ent: UserAccountSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := uas.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (uas *UserAccountSelect) StringsX(ctx context.Context) []string {
	v, err := uas.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (uas *UserAccountSelect) Ints(ctx context.Context) ([]int, error) {
	if len(uas.fields) > 1 {
		return nil, errors.New("ent: UserAccountSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := uas.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (uas *UserAccountSelect) IntsX(ctx context.Context) []int {
	v, err := uas.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (uas *UserAccountSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(uas.fields) > 1 {
		return nil, errors.New("ent: UserAccountSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := uas.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (uas *UserAccountSelect) Float64sX(ctx context.Context) []float64 {
	v, err := uas.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (uas *UserAccountSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(uas.fields) > 1 {
		return nil, errors.New("ent: UserAccountSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := uas.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (uas *UserAccountSelect) BoolsX(ctx context.Context) []bool {
	v, err := uas.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (uas *UserAccountSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := uas.sqlQuery().Query()
	if err := uas.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (uas *UserAccountSelect) sqlQuery() sql.Querier {
	selector := uas.sql
	selector.Select(selector.Columns(uas.fields...)...)
	return selector
}
