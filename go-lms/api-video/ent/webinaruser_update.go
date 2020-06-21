// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"repo.nefrosovet.ru/go-lms/api-video/ent/predicate"
	"repo.nefrosovet.ru/go-lms/api-video/ent/webinaruser"
)

// WebinarUserUpdate is the builder for updating WebinarUser entities.
type WebinarUserUpdate struct {
	config
	user_id             *int
	adduser_id          *int
	webinar_id          *int
	addwebinar_id       *int
	status              *webinaruser.Status
	clearstatus         bool
	medooze_id          *int
	addmedooze_id       *int
	clearmedooze_id     bool
	old_medooze_id      *int
	addold_medooze_id   *int
	clearold_medooze_id bool
	mic                 *int16
	addmic              *int16
	sound               *int16
	addsound            *int16
	predicates          []predicate.WebinarUser
}

// Where adds a new predicate for the builder.
func (wuu *WebinarUserUpdate) Where(ps ...predicate.WebinarUser) *WebinarUserUpdate {
	wuu.predicates = append(wuu.predicates, ps...)
	return wuu
}

// SetUserID sets the user_id field.
func (wuu *WebinarUserUpdate) SetUserID(i int) *WebinarUserUpdate {
	wuu.user_id = &i
	wuu.adduser_id = nil
	return wuu
}

// AddUserID adds i to user_id.
func (wuu *WebinarUserUpdate) AddUserID(i int) *WebinarUserUpdate {
	if wuu.adduser_id == nil {
		wuu.adduser_id = &i
	} else {
		*wuu.adduser_id += i
	}
	return wuu
}

// SetWebinarID sets the webinar_id field.
func (wuu *WebinarUserUpdate) SetWebinarID(i int) *WebinarUserUpdate {
	wuu.webinar_id = &i
	wuu.addwebinar_id = nil
	return wuu
}

// AddWebinarID adds i to webinar_id.
func (wuu *WebinarUserUpdate) AddWebinarID(i int) *WebinarUserUpdate {
	if wuu.addwebinar_id == nil {
		wuu.addwebinar_id = &i
	} else {
		*wuu.addwebinar_id += i
	}
	return wuu
}

// SetStatus sets the status field.
func (wuu *WebinarUserUpdate) SetStatus(w webinaruser.Status) *WebinarUserUpdate {
	wuu.status = &w
	return wuu
}

// SetNillableStatus sets the status field if the given value is not nil.
func (wuu *WebinarUserUpdate) SetNillableStatus(w *webinaruser.Status) *WebinarUserUpdate {
	if w != nil {
		wuu.SetStatus(*w)
	}
	return wuu
}

// ClearStatus clears the value of status.
func (wuu *WebinarUserUpdate) ClearStatus() *WebinarUserUpdate {
	wuu.status = nil
	wuu.clearstatus = true
	return wuu
}

// SetMedoozeID sets the medooze_id field.
func (wuu *WebinarUserUpdate) SetMedoozeID(i int) *WebinarUserUpdate {
	wuu.medooze_id = &i
	wuu.addmedooze_id = nil
	return wuu
}

// SetNillableMedoozeID sets the medooze_id field if the given value is not nil.
func (wuu *WebinarUserUpdate) SetNillableMedoozeID(i *int) *WebinarUserUpdate {
	if i != nil {
		wuu.SetMedoozeID(*i)
	}
	return wuu
}

// AddMedoozeID adds i to medooze_id.
func (wuu *WebinarUserUpdate) AddMedoozeID(i int) *WebinarUserUpdate {
	if wuu.addmedooze_id == nil {
		wuu.addmedooze_id = &i
	} else {
		*wuu.addmedooze_id += i
	}
	return wuu
}

// ClearMedoozeID clears the value of medooze_id.
func (wuu *WebinarUserUpdate) ClearMedoozeID() *WebinarUserUpdate {
	wuu.medooze_id = nil
	wuu.clearmedooze_id = true
	return wuu
}

// SetOldMedoozeID sets the old_medooze_id field.
func (wuu *WebinarUserUpdate) SetOldMedoozeID(i int) *WebinarUserUpdate {
	wuu.old_medooze_id = &i
	wuu.addold_medooze_id = nil
	return wuu
}

// SetNillableOldMedoozeID sets the old_medooze_id field if the given value is not nil.
func (wuu *WebinarUserUpdate) SetNillableOldMedoozeID(i *int) *WebinarUserUpdate {
	if i != nil {
		wuu.SetOldMedoozeID(*i)
	}
	return wuu
}

// AddOldMedoozeID adds i to old_medooze_id.
func (wuu *WebinarUserUpdate) AddOldMedoozeID(i int) *WebinarUserUpdate {
	if wuu.addold_medooze_id == nil {
		wuu.addold_medooze_id = &i
	} else {
		*wuu.addold_medooze_id += i
	}
	return wuu
}

// ClearOldMedoozeID clears the value of old_medooze_id.
func (wuu *WebinarUserUpdate) ClearOldMedoozeID() *WebinarUserUpdate {
	wuu.old_medooze_id = nil
	wuu.clearold_medooze_id = true
	return wuu
}

// SetMic sets the mic field.
func (wuu *WebinarUserUpdate) SetMic(i int16) *WebinarUserUpdate {
	wuu.mic = &i
	wuu.addmic = nil
	return wuu
}

// SetNillableMic sets the mic field if the given value is not nil.
func (wuu *WebinarUserUpdate) SetNillableMic(i *int16) *WebinarUserUpdate {
	if i != nil {
		wuu.SetMic(*i)
	}
	return wuu
}

// AddMic adds i to mic.
func (wuu *WebinarUserUpdate) AddMic(i int16) *WebinarUserUpdate {
	if wuu.addmic == nil {
		wuu.addmic = &i
	} else {
		*wuu.addmic += i
	}
	return wuu
}

// SetSound sets the sound field.
func (wuu *WebinarUserUpdate) SetSound(i int16) *WebinarUserUpdate {
	wuu.sound = &i
	wuu.addsound = nil
	return wuu
}

// SetNillableSound sets the sound field if the given value is not nil.
func (wuu *WebinarUserUpdate) SetNillableSound(i *int16) *WebinarUserUpdate {
	if i != nil {
		wuu.SetSound(*i)
	}
	return wuu
}

// AddSound adds i to sound.
func (wuu *WebinarUserUpdate) AddSound(i int16) *WebinarUserUpdate {
	if wuu.addsound == nil {
		wuu.addsound = &i
	} else {
		*wuu.addsound += i
	}
	return wuu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (wuu *WebinarUserUpdate) Save(ctx context.Context) (int, error) {
	if wuu.status != nil {
		if err := webinaruser.StatusValidator(*wuu.status); err != nil {
			return 0, fmt.Errorf("ent: validator failed for field \"status\": %v", err)
		}
	}
	return wuu.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (wuu *WebinarUserUpdate) SaveX(ctx context.Context) int {
	affected, err := wuu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (wuu *WebinarUserUpdate) Exec(ctx context.Context) error {
	_, err := wuu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wuu *WebinarUserUpdate) ExecX(ctx context.Context) {
	if err := wuu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (wuu *WebinarUserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   webinaruser.Table,
			Columns: webinaruser.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: webinaruser.FieldID,
			},
		},
	}
	if ps := wuu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value := wuu.user_id; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  *value,
			Column: webinaruser.FieldUserID,
		})
	}
	if value := wuu.adduser_id; value != nil {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  *value,
			Column: webinaruser.FieldUserID,
		})
	}
	if value := wuu.webinar_id; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  *value,
			Column: webinaruser.FieldWebinarID,
		})
	}
	if value := wuu.addwebinar_id; value != nil {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  *value,
			Column: webinaruser.FieldWebinarID,
		})
	}
	if value := wuu.status; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  *value,
			Column: webinaruser.FieldStatus,
		})
	}
	if wuu.clearstatus {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Column: webinaruser.FieldStatus,
		})
	}
	if value := wuu.medooze_id; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  *value,
			Column: webinaruser.FieldMedoozeID,
		})
	}
	if value := wuu.addmedooze_id; value != nil {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  *value,
			Column: webinaruser.FieldMedoozeID,
		})
	}
	if wuu.clearmedooze_id {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: webinaruser.FieldMedoozeID,
		})
	}
	if value := wuu.old_medooze_id; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  *value,
			Column: webinaruser.FieldOldMedoozeID,
		})
	}
	if value := wuu.addold_medooze_id; value != nil {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  *value,
			Column: webinaruser.FieldOldMedoozeID,
		})
	}
	if wuu.clearold_medooze_id {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: webinaruser.FieldOldMedoozeID,
		})
	}
	if value := wuu.mic; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt16,
			Value:  *value,
			Column: webinaruser.FieldMic,
		})
	}
	if value := wuu.addmic; value != nil {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt16,
			Value:  *value,
			Column: webinaruser.FieldMic,
		})
	}
	if value := wuu.sound; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt16,
			Value:  *value,
			Column: webinaruser.FieldSound,
		})
	}
	if value := wuu.addsound; value != nil {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt16,
			Value:  *value,
			Column: webinaruser.FieldSound,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, wuu.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// WebinarUserUpdateOne is the builder for updating a single WebinarUser entity.
type WebinarUserUpdateOne struct {
	config
	id                  int
	user_id             *int
	adduser_id          *int
	webinar_id          *int
	addwebinar_id       *int
	status              *webinaruser.Status
	clearstatus         bool
	medooze_id          *int
	addmedooze_id       *int
	clearmedooze_id     bool
	old_medooze_id      *int
	addold_medooze_id   *int
	clearold_medooze_id bool
	mic                 *int16
	addmic              *int16
	sound               *int16
	addsound            *int16
}

// SetUserID sets the user_id field.
func (wuuo *WebinarUserUpdateOne) SetUserID(i int) *WebinarUserUpdateOne {
	wuuo.user_id = &i
	wuuo.adduser_id = nil
	return wuuo
}

// AddUserID adds i to user_id.
func (wuuo *WebinarUserUpdateOne) AddUserID(i int) *WebinarUserUpdateOne {
	if wuuo.adduser_id == nil {
		wuuo.adduser_id = &i
	} else {
		*wuuo.adduser_id += i
	}
	return wuuo
}

// SetWebinarID sets the webinar_id field.
func (wuuo *WebinarUserUpdateOne) SetWebinarID(i int) *WebinarUserUpdateOne {
	wuuo.webinar_id = &i
	wuuo.addwebinar_id = nil
	return wuuo
}

// AddWebinarID adds i to webinar_id.
func (wuuo *WebinarUserUpdateOne) AddWebinarID(i int) *WebinarUserUpdateOne {
	if wuuo.addwebinar_id == nil {
		wuuo.addwebinar_id = &i
	} else {
		*wuuo.addwebinar_id += i
	}
	return wuuo
}

// SetStatus sets the status field.
func (wuuo *WebinarUserUpdateOne) SetStatus(w webinaruser.Status) *WebinarUserUpdateOne {
	wuuo.status = &w
	return wuuo
}

// SetNillableStatus sets the status field if the given value is not nil.
func (wuuo *WebinarUserUpdateOne) SetNillableStatus(w *webinaruser.Status) *WebinarUserUpdateOne {
	if w != nil {
		wuuo.SetStatus(*w)
	}
	return wuuo
}

// ClearStatus clears the value of status.
func (wuuo *WebinarUserUpdateOne) ClearStatus() *WebinarUserUpdateOne {
	wuuo.status = nil
	wuuo.clearstatus = true
	return wuuo
}

// SetMedoozeID sets the medooze_id field.
func (wuuo *WebinarUserUpdateOne) SetMedoozeID(i int) *WebinarUserUpdateOne {
	wuuo.medooze_id = &i
	wuuo.addmedooze_id = nil
	return wuuo
}

// SetNillableMedoozeID sets the medooze_id field if the given value is not nil.
func (wuuo *WebinarUserUpdateOne) SetNillableMedoozeID(i *int) *WebinarUserUpdateOne {
	if i != nil {
		wuuo.SetMedoozeID(*i)
	}
	return wuuo
}

// AddMedoozeID adds i to medooze_id.
func (wuuo *WebinarUserUpdateOne) AddMedoozeID(i int) *WebinarUserUpdateOne {
	if wuuo.addmedooze_id == nil {
		wuuo.addmedooze_id = &i
	} else {
		*wuuo.addmedooze_id += i
	}
	return wuuo
}

// ClearMedoozeID clears the value of medooze_id.
func (wuuo *WebinarUserUpdateOne) ClearMedoozeID() *WebinarUserUpdateOne {
	wuuo.medooze_id = nil
	wuuo.clearmedooze_id = true
	return wuuo
}

// SetOldMedoozeID sets the old_medooze_id field.
func (wuuo *WebinarUserUpdateOne) SetOldMedoozeID(i int) *WebinarUserUpdateOne {
	wuuo.old_medooze_id = &i
	wuuo.addold_medooze_id = nil
	return wuuo
}

// SetNillableOldMedoozeID sets the old_medooze_id field if the given value is not nil.
func (wuuo *WebinarUserUpdateOne) SetNillableOldMedoozeID(i *int) *WebinarUserUpdateOne {
	if i != nil {
		wuuo.SetOldMedoozeID(*i)
	}
	return wuuo
}

// AddOldMedoozeID adds i to old_medooze_id.
func (wuuo *WebinarUserUpdateOne) AddOldMedoozeID(i int) *WebinarUserUpdateOne {
	if wuuo.addold_medooze_id == nil {
		wuuo.addold_medooze_id = &i
	} else {
		*wuuo.addold_medooze_id += i
	}
	return wuuo
}

// ClearOldMedoozeID clears the value of old_medooze_id.
func (wuuo *WebinarUserUpdateOne) ClearOldMedoozeID() *WebinarUserUpdateOne {
	wuuo.old_medooze_id = nil
	wuuo.clearold_medooze_id = true
	return wuuo
}

// SetMic sets the mic field.
func (wuuo *WebinarUserUpdateOne) SetMic(i int16) *WebinarUserUpdateOne {
	wuuo.mic = &i
	wuuo.addmic = nil
	return wuuo
}

// SetNillableMic sets the mic field if the given value is not nil.
func (wuuo *WebinarUserUpdateOne) SetNillableMic(i *int16) *WebinarUserUpdateOne {
	if i != nil {
		wuuo.SetMic(*i)
	}
	return wuuo
}

// AddMic adds i to mic.
func (wuuo *WebinarUserUpdateOne) AddMic(i int16) *WebinarUserUpdateOne {
	if wuuo.addmic == nil {
		wuuo.addmic = &i
	} else {
		*wuuo.addmic += i
	}
	return wuuo
}

// SetSound sets the sound field.
func (wuuo *WebinarUserUpdateOne) SetSound(i int16) *WebinarUserUpdateOne {
	wuuo.sound = &i
	wuuo.addsound = nil
	return wuuo
}

// SetNillableSound sets the sound field if the given value is not nil.
func (wuuo *WebinarUserUpdateOne) SetNillableSound(i *int16) *WebinarUserUpdateOne {
	if i != nil {
		wuuo.SetSound(*i)
	}
	return wuuo
}

// AddSound adds i to sound.
func (wuuo *WebinarUserUpdateOne) AddSound(i int16) *WebinarUserUpdateOne {
	if wuuo.addsound == nil {
		wuuo.addsound = &i
	} else {
		*wuuo.addsound += i
	}
	return wuuo
}

// Save executes the query and returns the updated entity.
func (wuuo *WebinarUserUpdateOne) Save(ctx context.Context) (*WebinarUser, error) {
	if wuuo.status != nil {
		if err := webinaruser.StatusValidator(*wuuo.status); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"status\": %v", err)
		}
	}
	return wuuo.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (wuuo *WebinarUserUpdateOne) SaveX(ctx context.Context) *WebinarUser {
	wu, err := wuuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return wu
}

// Exec executes the query on the entity.
func (wuuo *WebinarUserUpdateOne) Exec(ctx context.Context) error {
	_, err := wuuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wuuo *WebinarUserUpdateOne) ExecX(ctx context.Context) {
	if err := wuuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (wuuo *WebinarUserUpdateOne) sqlSave(ctx context.Context) (wu *WebinarUser, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   webinaruser.Table,
			Columns: webinaruser.Columns,
			ID: &sqlgraph.FieldSpec{
				Value:  wuuo.id,
				Type:   field.TypeInt,
				Column: webinaruser.FieldID,
			},
		},
	}
	if value := wuuo.user_id; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  *value,
			Column: webinaruser.FieldUserID,
		})
	}
	if value := wuuo.adduser_id; value != nil {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  *value,
			Column: webinaruser.FieldUserID,
		})
	}
	if value := wuuo.webinar_id; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  *value,
			Column: webinaruser.FieldWebinarID,
		})
	}
	if value := wuuo.addwebinar_id; value != nil {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  *value,
			Column: webinaruser.FieldWebinarID,
		})
	}
	if value := wuuo.status; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  *value,
			Column: webinaruser.FieldStatus,
		})
	}
	if wuuo.clearstatus {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Column: webinaruser.FieldStatus,
		})
	}
	if value := wuuo.medooze_id; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  *value,
			Column: webinaruser.FieldMedoozeID,
		})
	}
	if value := wuuo.addmedooze_id; value != nil {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  *value,
			Column: webinaruser.FieldMedoozeID,
		})
	}
	if wuuo.clearmedooze_id {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: webinaruser.FieldMedoozeID,
		})
	}
	if value := wuuo.old_medooze_id; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  *value,
			Column: webinaruser.FieldOldMedoozeID,
		})
	}
	if value := wuuo.addold_medooze_id; value != nil {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  *value,
			Column: webinaruser.FieldOldMedoozeID,
		})
	}
	if wuuo.clearold_medooze_id {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: webinaruser.FieldOldMedoozeID,
		})
	}
	if value := wuuo.mic; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt16,
			Value:  *value,
			Column: webinaruser.FieldMic,
		})
	}
	if value := wuuo.addmic; value != nil {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt16,
			Value:  *value,
			Column: webinaruser.FieldMic,
		})
	}
	if value := wuuo.sound; value != nil {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt16,
			Value:  *value,
			Column: webinaruser.FieldSound,
		})
	}
	if value := wuuo.addsound; value != nil {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt16,
			Value:  *value,
			Column: webinaruser.FieldSound,
		})
	}
	wu = &WebinarUser{config: wuuo.config}
	_spec.Assign = wu.assignValues
	_spec.ScanValues = wu.scanValues()
	if err = sqlgraph.UpdateNode(ctx, wuuo.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return wu, nil
}
