package storage

type RoleStorage interface {
	Store(role Role) error
	Update(role Role) error
	Get(id string) (*Role, error)
	GetAll() ([]*Role, error)
	Delete(id string) error
}

const (
	// Default roles
	RoleDefaultAdmin    = "ADMIN"
	RoleDefaultClient   = "CLIENT"
	RoleDefaultEmployee = "EMPLOYEE"
	RoleDefaultPatient  = "PATIENT"
)

// Role is a struct with information
type Role struct {
	ID          string `bson:"id" json:"id"`
	Description string `bson:"description,omitempty" json:"description,omitempty"`
}

// IsDefaultRole returns true if role is reserved.
func IsDefaultRole(id string) bool {
	switch id {
	default:
		return false
	case RoleDefaultAdmin:
		return true
	case RoleDefaultEmployee:
		return true
	case RoleDefaultPatient:
		return true
	}
}