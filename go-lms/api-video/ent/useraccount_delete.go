// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"repo.nefrosovet.ru/go-lms/api-video/ent/predicate"
	"repo.nefrosovet.ru/go-lms/api-video/ent/useraccount"
)

// UserAccountDelete is the builder for deleting a UserAccount entity.
type UserAccountDelete struct {
	config
	predicates []predicate.UserAccount
}

// Where adds a new predicate to the delete builder.
func (uad *UserAccountDelete) Where(ps ...predicate.UserAccount) *UserAccountDelete {
	uad.predicates = append(uad.predicates, ps...)
	return uad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (uad *UserAccountDelete) Exec(ctx context.Context) (int, error) {
	return uad.sqlExec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (uad *UserAccountDelete) ExecX(ctx context.Context) int {
	n, err := uad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (uad *UserAccountDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: useraccount.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: useraccount.FieldID,
			},
		},
	}
	if ps := uad.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, uad.driver, _spec)
}

// UserAccountDeleteOne is the builder for deleting a single UserAccount entity.
type UserAccountDeleteOne struct {
	uad *UserAccountDelete
}

// Exec executes the deletion query.
func (uado *UserAccountDeleteOne) Exec(ctx context.Context) error {
	n, err := uado.uad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{useraccount.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (uado *UserAccountDeleteOne) ExecX(ctx context.Context) {
	uado.uad.ExecX(ctx)
}
