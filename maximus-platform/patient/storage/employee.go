package storage

import "gopkg.in/guregu/null.v3"

type Employee struct {
	BaseModel

	ID null.String `db:"id"`

	FirstName    null.String `db:"first_name"`
	LastName     null.String `db:"last_name"`
	MiddleName   null.String `db:"middle_name"`
	PositionCode null.String `db:"position_code"`
	PhotoGUID    null.String `db:"photo_guid"`
}
