package storage

import "gopkg.in/guregu/null.v3"

type Clinic struct {
	BaseModel

	ID null.String `db:"id"`

	Title null.String `db:"title"`
}
