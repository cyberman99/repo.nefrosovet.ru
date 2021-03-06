// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"repo.nefrosovet.ru/go-lms/api-video/ent/subscriber"
	"repo.nefrosovet.ru/go-lms/api-video/ent/user"
)

// SubscriberCreate is the builder for creating a Subscriber entity.
type SubscriberCreate struct {
	config
	username *string
	domain   *string
	ha1      *string
	ha1b     *string
	user     map[int]struct{}
}

// SetUsername sets the username field.
func (sc *SubscriberCreate) SetUsername(s string) *SubscriberCreate {
	sc.username = &s
	return sc
}

// SetNillableUsername sets the username field if the given value is not nil.
func (sc *SubscriberCreate) SetNillableUsername(s *string) *SubscriberCreate {
	if s != nil {
		sc.SetUsername(*s)
	}
	return sc
}

// SetDomain sets the domain field.
func (sc *SubscriberCreate) SetDomain(s string) *SubscriberCreate {
	sc.domain = &s
	return sc
}

// SetNillableDomain sets the domain field if the given value is not nil.
func (sc *SubscriberCreate) SetNillableDomain(s *string) *SubscriberCreate {
	if s != nil {
		sc.SetDomain(*s)
	}
	return sc
}

// SetHa1 sets the ha1 field.
func (sc *SubscriberCreate) SetHa1(s string) *SubscriberCreate {
	sc.ha1 = &s
	return sc
}

// SetNillableHa1 sets the ha1 field if the given value is not nil.
func (sc *SubscriberCreate) SetNillableHa1(s *string) *SubscriberCreate {
	if s != nil {
		sc.SetHa1(*s)
	}
	return sc
}

// SetHa1b sets the ha1b field.
func (sc *SubscriberCreate) SetHa1b(s string) *SubscriberCreate {
	sc.ha1b = &s
	return sc
}

// SetNillableHa1b sets the ha1b field if the given value is not nil.
func (sc *SubscriberCreate) SetNillableHa1b(s *string) *SubscriberCreate {
	if s != nil {
		sc.SetHa1b(*s)
	}
	return sc
}

// SetUserID sets the user edge to User by id.
func (sc *SubscriberCreate) SetUserID(id int) *SubscriberCreate {
	if sc.user == nil {
		sc.user = make(map[int]struct{})
	}
	sc.user[id] = struct{}{}
	return sc
}

// SetUser sets the user edge to User.
func (sc *SubscriberCreate) SetUser(u *User) *SubscriberCreate {
	return sc.SetUserID(u.ID)
}

// Save creates the Subscriber in the database.
func (sc *SubscriberCreate) Save(ctx context.Context) (*Subscriber, error) {
	if sc.username == nil {
		v := subscriber.DefaultUsername
		sc.username = &v
	}
	if err := subscriber.UsernameValidator(*sc.username); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"username\": %v", err)
	}
	if sc.domain == nil {
		v := subscriber.DefaultDomain
		sc.domain = &v
	}
	if err := subscriber.DomainValidator(*sc.domain); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"domain\": %v", err)
	}
	if sc.ha1 == nil {
		v := subscriber.DefaultHa1
		sc.ha1 = &v
	}
	if err := subscriber.Ha1Validator(*sc.ha1); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"ha1\": %v", err)
	}
	if sc.ha1b == nil {
		v := subscriber.DefaultHa1b
		sc.ha1b = &v
	}
	if err := subscriber.Ha1bValidator(*sc.ha1b); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"ha1b\": %v", err)
	}
	if len(sc.user) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"user\"")
	}
	if sc.user == nil {
		return nil, errors.New("ent: missing required edge \"user\"")
	}
	return sc.sqlSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SubscriberCreate) SaveX(ctx context.Context) *Subscriber {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (sc *SubscriberCreate) sqlSave(ctx context.Context) (*Subscriber, error) {
	var (
		s     = &Subscriber{config: sc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: subscriber.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: subscriber.FieldID,
			},
		}
	)
	if value := sc.username; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: subscriber.FieldUsername,
		})
		s.Username = *value
	}
	if value := sc.domain; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: subscriber.FieldDomain,
		})
		s.Domain = *value
	}
	if value := sc.ha1; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: subscriber.FieldHa1,
		})
		s.Ha1 = *value
	}
	if value := sc.ha1b; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: subscriber.FieldHa1b,
		})
		s.Ha1b = *value
	}
	if nodes := sc.user; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   subscriber.UserTable,
			Columns: []string{subscriber.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	s.ID = int(id)
	return s, nil
}
