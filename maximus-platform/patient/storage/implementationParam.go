package storage

import "gopkg.in/guregu/null.v3"

type ImplementationParam struct {
	BaseModel

	ID null.String `db:"id"`

	ImplementationID null.String `db:"implementation_id"`

	TypeCode null.String `db:"type_code"`
	Value    null.String `db:"value"`
}
