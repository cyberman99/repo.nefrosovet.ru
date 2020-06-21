package storage

import (
	"fmt"
)

var (
	ErrBackendsOrderNotFound = fmt.Errorf("backends order: %w", ErrNotFound)
	ErrBackendsOrderAlreadyExists = fmt.Errorf("backends order: %w", ErrAlreadyExists)

	DefaultBackendsOrder = []string{"index"}
)

type BackendsOrderStorage interface {
	Store(in StoreBackendsOrder) (*BackendsOrder, error)
	Update(in UpdateBackendsOrder) (*BackendsOrder, error)
	Get() (*BackendsOrder, error)
	Delete(in DeleteBackendsOrder) error
}

type BackendsOrder struct {
	ID    string   `json:"id" bson:"id"`
	Order []string `json:"order" bson:"order"`
}

type StoreBackendsOrder struct {
	Order []string
}

type UpdateBackendsOrder struct {
	Order []string
}

type DeleteBackendsOrder struct {
	IDs []string
}
