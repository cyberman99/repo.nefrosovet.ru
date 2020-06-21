package storage

import "gopkg.in/guregu/null.v3"

type Patient struct {
	BaseModel

	ID         null.String `db:"id"`
	StatusCode null.String `db:"status_code"`
}
