package mongo

import (
	"context"
	"errors"
	"reflect"
	"sort"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	dbMongo "repo.nefrosovet.ru/maximus-platform/profile/db/mongo"
	"repo.nefrosovet.ru/maximus-platform/profile/storage"
)

const UserCollectionName = "users"

var (
	ErrInsertedIDNotFound = errors.New("insertedID not found")
)

type UserStorage struct {
	client *dbMongo.Client

	storage.UserStorage
}

func newUserStorage(client *dbMongo.Client) *UserStorage {
	return &UserStorage{
		client: client,
	}
}

func (s *UserStorage) StoreUser(in storage.StoreUser) (storage.User, error) {
	collection := s.client.Collection(UserCollectionName)
	ctx, _ := context.WithTimeout(background, time.Second*1)

	var err error

	if err = areContactsUnique(in.Contacts, collection, ctx); err != nil {
		return storage.User{}, err
	}

	for i, _ := range in.Contacts {
		in.Contacts[i].Created, err = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		if err != nil {
			return storage.User{}, err
		}
	}

	_, err = collection.InsertOne(ctx, in)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return storage.User{}, storage.ErrUserAlreadyExists
		}

		return storage.User{}, err
	}

	return storage.User{
		ID:           storage.UUID(in.ID),
		PasswordHash: in.PasswordHash,
		FirstName:    in.FirstName,
		LastName:     in.LastName,
		MiddleName:   in.MiddleName,
		Contacts:     in.Contacts,
	}, nil
}

func (s *UserStorage) UpdateUser(id storage.UUID, in storage.UpdateUser) (storage.User, error) {
	collection := s.client.Collection(UserCollectionName)
	ctx, _ := context.WithTimeout(background, time.Second*1)

	var err error

	if err = areContactsUnique(*in.Contacts, collection, ctx); err != nil {
		return storage.User{}, err
	}

	filter := bson.M{
		"_id": id,
	}
	update := bson.M{}

	if in.PasswordHash != nil {
		update["passwordHash"] = *in.PasswordHash
	}

	if in.FirstName != nil {
		update["firstName"] = *in.FirstName
	}

	if in.LastName != nil {
		update["lastName"] = *in.LastName
	}

	if in.MiddleName != nil {
		update["middleName"] = *in.MiddleName
	}

	contactsUpdate := bson.M{
		"contacts": bson.M{
			"$each": in.Contacts,
		},
	}

	result := collection.FindOneAndUpdate(ctx, filter, bson.M{
		"$set":      update,
		"$addToSet": contactsUpdate,
	}, options.FindOneAndUpdate().SetReturnDocument(options.After))
	if result.Err() != nil {
		switch {
		//case errors.Is(result.Err(), mongo.ErrNoDocuments):
		//	return storage.User{}, storage.ErrUserNotFound
		case strings.Contains(result.Err().Error(), "duplicate"):
			return storage.User{}, storage.ErrUserAlreadyExists
		}

		return storage.User{}, result.Err()
	}

	var user storage.User
	if err := result.Decode(&user); err != nil {
		return storage.User{}, err
	}

	return user, nil
}

func (s *UserStorage) GetUser(in storage.GetUser) (storage.User, error) {
	collection := s.client.Collection(UserCollectionName)
	ctx, _ := context.WithTimeout(background, time.Second*1)

	filter := bson.M{}

	if in.ID != nil {
		filter["_id"] = in.ID
	}

	if in.Value != nil {
		filter["contacts.value"] = in.Value
	}

	if len(filter) == 0 {
		return storage.User{}, storage.ErrBadInput
	}

	result := collection.FindOne(ctx, filter)
	if result.Err() != nil {
		switch result.Err() {
		case mongo.ErrNoDocuments:
			return storage.User{}, storage.ErrUserNotFound
		}

		return storage.User{}, result.Err()
	}

	var user storage.User
	if err := result.Decode(&user); err != nil {
		return storage.User{}, err
	}

	var contacts []storage.UserContact

	lastEmailContact := getLastContact("EMAIL", user, ctx)
	contacts = append(contacts, lastEmailContact)

	lastMobileContact:= getLastContact("MOBILE", user, ctx)
	contacts = append(contacts, lastMobileContact)

	user.Contacts = contacts

	return user, nil
}

func (s *UserStorage) GetUsers() ([]storage.User, error) {
	collection := s.client.Collection(UserCollectionName)
	ctx, _ := context.WithTimeout(background, time.Second*10)

	users := make([]storage.User, 0)
	if cur, err := collection.Find(ctx, bson.M{}, options.Find()); err != nil {
		return nil, err
	} else {
		if cur.Err() != nil {
			return nil, cur.Err()
		}
		defer cur.Close(ctx)

		for cur.Next(ctx) {
			var user storage.User
			if err := cur.Decode(&user); err != nil {
				return nil, err
			}

			var contacts []storage.UserContact

			lastEmailContact := getLastContact("EMAIL", user, ctx)
			contacts = append(contacts, lastEmailContact)

			lastMobileContact:= getLastContact("MOBILE", user, ctx)
			contacts = append(contacts, lastMobileContact)

			user.Contacts = contacts

			users = append(users, user)
		}
	}

	return users, nil
}

func areContactsUnique(contacts []storage.UserContact, coll *mongo.Collection, ctx context.Context) error {
	for _, contact := range contacts {
		filter := bson.M{
			"contacts.value": contact.Value,
			"contacts.verified": contact.Verified,
		}
		cur, err := coll.Find(ctx, filter, options.Find().SetLimit(1))
		if err != nil {
			return err
		}
		if cur.Err() != nil {
			return cur.Err()
		}

		for cur.Next(ctx) {
			var user storage.User
			if err := cur.Decode(&user); err != nil {
				return err
			}

			if !reflect.DeepEqual(user, storage.User{}) {
				return storage.ErrUserContactAlreadyExists
			}
		}

		cur.Close(ctx)
	}

	return nil
}

func getLastContact(contactType string, user storage.User, ctx context.Context) storage.UserContact{
	var contacts []storage.UserContact
	for _, c := range user.Contacts {
		if c.TypeCODE == contactType {
			contacts = append(contacts, c)
		}
	}

	if len(contacts) == 0 {
		return storage.UserContact{}
	}

	sort.SliceStable(contacts, func(i, j int) bool {
		return contacts[i].Created.Before(contacts[j].Created)
	})

	return contacts[0]
}

//func getLastContact(
//	contactType string,
//	userID storage.UUID,
//	coll *mongo.Collection,
//	ctx context.Context) (storage.UserContact, error) {
//	filter := bson.M{
//		"userID":   userID,
//		"typeCODE": contactType,
//	}
//	sort := bson.D{
//		{"created", -1},
//	}
//
//	result := coll.FindOne(ctx, filter, options.FindOne().SetSort(sort))
//	if result.Err() != nil {
//		switch result.Err() {
//		case mongo.ErrNoDocuments:
//			return storage.UserContact{}, storage.ErrUserContactNotFound
//		default:
//			return storage.UserContact{}, result.Err()
//		}
//	}
//
//	var userContact storage.UserContact
//	if err := result.Decode(&userContact); err != nil {
//		return storage.UserContact{}, err
//	}
//
//	return userContact, nil
//}
