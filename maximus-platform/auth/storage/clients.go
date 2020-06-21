package storage

type ClientStorage interface {
	Store(client ClientStorer) error
	Update(clientId string, updater ClientUpdater) error
	Get(filter ClientFilter) ([]*ClientStorer, error)
	Delete(clientId string) error
}

type ClientFilter struct {
	ID *string
}

type ClientUpdater struct {
	Descriptions *string
	Password     *string
}

type ClientStorer struct {
	ID           string `json:"id" bson:"id"`
	Descriptions string `json:"description" bson:"description"`
	Password     string `json:"password" bson:"password"`
}