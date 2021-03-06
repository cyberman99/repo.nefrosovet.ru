// Code generated by entc, DO NOT EDIT.

package webinaruser

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"repo.nefrosovet.ru/go-lms/api-video/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.WebinarUser {
	return predicate.WebinarUser(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldID), id))
		},
	)
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.WebinarUser {
	return predicate.WebinarUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	},
	)
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.WebinarUser {
	return predicate.WebinarUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	},
	)
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.WebinarUser {
	return predicate.WebinarUser(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	},
	)
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.WebinarUser {
	return predicate.WebinarUser(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	},
	)
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.WebinarUser {
	return predicate.WebinarUser(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	},
	)
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.WebinarUser {
	return predicate.WebinarUser(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	},
	)
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.WebinarUser {
	return predicate.WebinarUser(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	},
	)
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.WebinarUser {
	return predicate.WebinarUser(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	},
	)
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int) predicate.WebinarUser {
	return predicate.WebinarUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	},
	)
}

// WebinarID applies equality check predicate on the "webinar_id" field. It's identical to WebinarIDEQ.
func WebinarID(v int) predicate.WebinarUser {
	return predicate.WebinarUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWebinarID), v))
	},
	)
}

// MedoozeID applies equality check predicate on the "medooze_id" field. It's identical to MedoozeIDEQ.
func MedoozeID(v int) predicate.WebinarUser {
	return predicate.WebinarUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMedoozeID), v))
	},
	)
}

// OldMedoozeID applies equality check predicate on the "old_medooze_id" field. It's identical to OldMedoozeIDEQ.
func OldMedoozeID(v int) predicate.WebinarUser {
	return predicate.WebinarUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOldMedoozeID), v))
	},
	)
}

// Mic applies equality check predicate on the "mic" field. It's identical to MicEQ.
func Mic(v int16) predicate.WebinarUser {
	return predicate.WebinarUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMic), v))
	},
	)
}

// Sound applies equality check predicate on the "sound" field. It's identical to SoundEQ.
func Sound(v int16) predicate.WebinarUser {
	return predicate.WebinarUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSound), v))
	},
	)
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int) predicate.WebinarUser {
	return predicate.WebinarUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	},
	)
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int) predicate.WebinarUser {
	return predicate.WebinarUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUserID), v))
	},
	)
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int) predicate.WebinarUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WebinarUser(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUserID), v...))
	},
	)
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int) predicate.WebinarUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WebinarUser(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUserID), v...))
	},
	)
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v int) predicate.WebinarUser {
	return predicate.WebinarUser(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUserID), v))
	},
	)
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v int) predicate.WebinarUser {
	return predicate.WebinarUser(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUserID), v))
	},
	)
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v int) predicate.WebinarUser {
	return predicate.WebinarUser(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUserID), v))
	},
	)
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v int) predicate.WebinarUser {
	return predicate.WebinarUser(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUserID), v))
	},
	)
}

// WebinarIDEQ applies the EQ predicate on the "webinar_id" field.
func WebinarIDEQ(v int) predicate.WebinarUser {
	return predicate.WebinarUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWebinarID), v))
	},
	)
}

// WebinarIDNEQ applies the NEQ predicate on the "webinar_id" field.
func WebinarIDNEQ(v int) predicate.WebinarUser {
	return predicate.WebinarUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldWebinarID), v))
	},
	)
}

// WebinarIDIn applies the In predicate on the "webinar_id" field.
func WebinarIDIn(vs ...int) predicate.WebinarUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WebinarUser(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldWebinarID), v...))
	},
	)
}

// WebinarIDNotIn applies the NotIn predicate on the "webinar_id" field.
func WebinarIDNotIn(vs ...int) predicate.WebinarUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WebinarUser(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldWebinarID), v...))
	},
	)
}

// WebinarIDGT applies the GT predicate on the "webinar_id" field.
func WebinarIDGT(v int) predicate.WebinarUser {
	return predicate.WebinarUser(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldWebinarID), v))
	},
	)
}

// WebinarIDGTE applies the GTE predicate on the "webinar_id" field.
func WebinarIDGTE(v int) predicate.WebinarUser {
	return predicate.WebinarUser(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldWebinarID), v))
	},
	)
}

// WebinarIDLT applies the LT predicate on the "webinar_id" field.
func WebinarIDLT(v int) predicate.WebinarUser {
	return predicate.WebinarUser(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldWebinarID), v))
	},
	)
}

// WebinarIDLTE applies the LTE predicate on the "webinar_id" field.
func WebinarIDLTE(v int) predicate.WebinarUser {
	return predicate.WebinarUser(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldWebinarID), v))
	},
	)
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.WebinarUser {
	return predicate.WebinarUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	},
	)
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.WebinarUser {
	return predicate.WebinarUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStatus), v))
	},
	)
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.WebinarUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WebinarUser(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStatus), v...))
	},
	)
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.WebinarUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WebinarUser(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStatus), v...))
	},
	)
}

// StatusIsNil applies the IsNil predicate on the "status" field.
func StatusIsNil() predicate.WebinarUser {
	return predicate.WebinarUser(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldStatus)))
	},
	)
}

// StatusNotNil applies the NotNil predicate on the "status" field.
func StatusNotNil() predicate.WebinarUser {
	return predicate.WebinarUser(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldStatus)))
	},
	)
}

// MedoozeIDEQ applies the EQ predicate on the "medooze_id" field.
func MedoozeIDEQ(v int) predicate.WebinarUser {
	return predicate.WebinarUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMedoozeID), v))
	},
	)
}

// MedoozeIDNEQ applies the NEQ predicate on the "medooze_id" field.
func MedoozeIDNEQ(v int) predicate.WebinarUser {
	return predicate.WebinarUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMedoozeID), v))
	},
	)
}

// MedoozeIDIn applies the In predicate on the "medooze_id" field.
func MedoozeIDIn(vs ...int) predicate.WebinarUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WebinarUser(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMedoozeID), v...))
	},
	)
}

// MedoozeIDNotIn applies the NotIn predicate on the "medooze_id" field.
func MedoozeIDNotIn(vs ...int) predicate.WebinarUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WebinarUser(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMedoozeID), v...))
	},
	)
}

// MedoozeIDGT applies the GT predicate on the "medooze_id" field.
func MedoozeIDGT(v int) predicate.WebinarUser {
	return predicate.WebinarUser(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMedoozeID), v))
	},
	)
}

// MedoozeIDGTE applies the GTE predicate on the "medooze_id" field.
func MedoozeIDGTE(v int) predicate.WebinarUser {
	return predicate.WebinarUser(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMedoozeID), v))
	},
	)
}

// MedoozeIDLT applies the LT predicate on the "medooze_id" field.
func MedoozeIDLT(v int) predicate.WebinarUser {
	return predicate.WebinarUser(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMedoozeID), v))
	},
	)
}

// MedoozeIDLTE applies the LTE predicate on the "medooze_id" field.
func MedoozeIDLTE(v int) predicate.WebinarUser {
	return predicate.WebinarUser(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMedoozeID), v))
	},
	)
}

// MedoozeIDIsNil applies the IsNil predicate on the "medooze_id" field.
func MedoozeIDIsNil() predicate.WebinarUser {
	return predicate.WebinarUser(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldMedoozeID)))
	},
	)
}

// MedoozeIDNotNil applies the NotNil predicate on the "medooze_id" field.
func MedoozeIDNotNil() predicate.WebinarUser {
	return predicate.WebinarUser(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldMedoozeID)))
	},
	)
}

// OldMedoozeIDEQ applies the EQ predicate on the "old_medooze_id" field.
func OldMedoozeIDEQ(v int) predicate.WebinarUser {
	return predicate.WebinarUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOldMedoozeID), v))
	},
	)
}

// OldMedoozeIDNEQ applies the NEQ predicate on the "old_medooze_id" field.
func OldMedoozeIDNEQ(v int) predicate.WebinarUser {
	return predicate.WebinarUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOldMedoozeID), v))
	},
	)
}

// OldMedoozeIDIn applies the In predicate on the "old_medooze_id" field.
func OldMedoozeIDIn(vs ...int) predicate.WebinarUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WebinarUser(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOldMedoozeID), v...))
	},
	)
}

// OldMedoozeIDNotIn applies the NotIn predicate on the "old_medooze_id" field.
func OldMedoozeIDNotIn(vs ...int) predicate.WebinarUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WebinarUser(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOldMedoozeID), v...))
	},
	)
}

// OldMedoozeIDGT applies the GT predicate on the "old_medooze_id" field.
func OldMedoozeIDGT(v int) predicate.WebinarUser {
	return predicate.WebinarUser(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOldMedoozeID), v))
	},
	)
}

// OldMedoozeIDGTE applies the GTE predicate on the "old_medooze_id" field.
func OldMedoozeIDGTE(v int) predicate.WebinarUser {
	return predicate.WebinarUser(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOldMedoozeID), v))
	},
	)
}

// OldMedoozeIDLT applies the LT predicate on the "old_medooze_id" field.
func OldMedoozeIDLT(v int) predicate.WebinarUser {
	return predicate.WebinarUser(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOldMedoozeID), v))
	},
	)
}

// OldMedoozeIDLTE applies the LTE predicate on the "old_medooze_id" field.
func OldMedoozeIDLTE(v int) predicate.WebinarUser {
	return predicate.WebinarUser(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOldMedoozeID), v))
	},
	)
}

// OldMedoozeIDIsNil applies the IsNil predicate on the "old_medooze_id" field.
func OldMedoozeIDIsNil() predicate.WebinarUser {
	return predicate.WebinarUser(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldOldMedoozeID)))
	},
	)
}

// OldMedoozeIDNotNil applies the NotNil predicate on the "old_medooze_id" field.
func OldMedoozeIDNotNil() predicate.WebinarUser {
	return predicate.WebinarUser(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldOldMedoozeID)))
	},
	)
}

// MicEQ applies the EQ predicate on the "mic" field.
func MicEQ(v int16) predicate.WebinarUser {
	return predicate.WebinarUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMic), v))
	},
	)
}

// MicNEQ applies the NEQ predicate on the "mic" field.
func MicNEQ(v int16) predicate.WebinarUser {
	return predicate.WebinarUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMic), v))
	},
	)
}

// MicIn applies the In predicate on the "mic" field.
func MicIn(vs ...int16) predicate.WebinarUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WebinarUser(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMic), v...))
	},
	)
}

// MicNotIn applies the NotIn predicate on the "mic" field.
func MicNotIn(vs ...int16) predicate.WebinarUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WebinarUser(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMic), v...))
	},
	)
}

// MicGT applies the GT predicate on the "mic" field.
func MicGT(v int16) predicate.WebinarUser {
	return predicate.WebinarUser(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMic), v))
	},
	)
}

// MicGTE applies the GTE predicate on the "mic" field.
func MicGTE(v int16) predicate.WebinarUser {
	return predicate.WebinarUser(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMic), v))
	},
	)
}

// MicLT applies the LT predicate on the "mic" field.
func MicLT(v int16) predicate.WebinarUser {
	return predicate.WebinarUser(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMic), v))
	},
	)
}

// MicLTE applies the LTE predicate on the "mic" field.
func MicLTE(v int16) predicate.WebinarUser {
	return predicate.WebinarUser(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMic), v))
	},
	)
}

// SoundEQ applies the EQ predicate on the "sound" field.
func SoundEQ(v int16) predicate.WebinarUser {
	return predicate.WebinarUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSound), v))
	},
	)
}

// SoundNEQ applies the NEQ predicate on the "sound" field.
func SoundNEQ(v int16) predicate.WebinarUser {
	return predicate.WebinarUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSound), v))
	},
	)
}

// SoundIn applies the In predicate on the "sound" field.
func SoundIn(vs ...int16) predicate.WebinarUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WebinarUser(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSound), v...))
	},
	)
}

// SoundNotIn applies the NotIn predicate on the "sound" field.
func SoundNotIn(vs ...int16) predicate.WebinarUser {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.WebinarUser(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSound), v...))
	},
	)
}

// SoundGT applies the GT predicate on the "sound" field.
func SoundGT(v int16) predicate.WebinarUser {
	return predicate.WebinarUser(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSound), v))
	},
	)
}

// SoundGTE applies the GTE predicate on the "sound" field.
func SoundGTE(v int16) predicate.WebinarUser {
	return predicate.WebinarUser(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSound), v))
	},
	)
}

// SoundLT applies the LT predicate on the "sound" field.
func SoundLT(v int16) predicate.WebinarUser {
	return predicate.WebinarUser(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSound), v))
	},
	)
}

// SoundLTE applies the LTE predicate on the "sound" field.
func SoundLTE(v int16) predicate.WebinarUser {
	return predicate.WebinarUser(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSound), v))
	},
	)
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.WebinarUser) predicate.WebinarUser {
	return predicate.WebinarUser(
		func(s *sql.Selector) {
			s1 := s.Clone().SetP(nil)
			for _, p := range predicates {
				p(s1)
			}
			s.Where(s1.P())
		},
	)
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.WebinarUser) predicate.WebinarUser {
	return predicate.WebinarUser(
		func(s *sql.Selector) {
			s1 := s.Clone().SetP(nil)
			for i, p := range predicates {
				if i > 0 {
					s1.Or()
				}
				p(s1)
			}
			s.Where(s1.P())
		},
	)
}

// Not applies the not operator on the given predicate.
func Not(p predicate.WebinarUser) predicate.WebinarUser {
	return predicate.WebinarUser(
		func(s *sql.Selector) {
			p(s.Not())
		},
	)
}
