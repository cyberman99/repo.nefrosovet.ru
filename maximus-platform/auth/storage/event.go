package storage

type EventStorage interface {
	Store(in StoreEvent) error
	Get(in GetEvent) (*Event, error)
	GetAll(in GetEvents) ([]*Event, error)
}

type Event struct {
	ID          string
	Type        string
	SourceIP    string
	EntityID    string
	EntityLogin string
	Status      string
	Time        string
}

type StoreEvent struct {
	EventType   string
	SourceIP    string
	EntityID    string
	EntityLogin string
	Status      string

	Data string
}

type GetEvent struct {
	ID string
}

type GetEvents struct {
	EventType   string
	SourceIP    string
	EntityID    string
	EntityLogin string
	Status      string

	Limit  int64
	Offset int64
}