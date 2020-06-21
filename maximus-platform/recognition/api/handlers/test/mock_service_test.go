package test

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/emirpasic/gods/maps/treemap"
	"github.com/go-openapi/strfmt"
	"github.com/gofrs/uuid"
	"log"
	"reflect"
	"repo.nefrosovet.ru/maximus-platform/recognition/services/AWS"
)

type storedPhoto struct {
	personId string
	body []byte
	format string
}

type mockService struct {
	*treemap.Map
}

func NewMockService() mockService {
	return mockService {
		Map: treemap.NewWithStringComparator(),
	}
}

func (m *mockService) Set(personID string, format string, body []byte) (_ *strfmt.UUID, err error) {
	var photoID = new(strfmt.UUID)
	err = photoID.Scan(uuid.Must(uuid.NewV1()).String())
	if err != nil {
		log.Fatal("Check uuid correct")
	}

	m.Map.Put(photoID.String(), storedPhoto{
		personId: personID,
		format: format,
		body: body,
	})

	return photoID, nil
}
func (m *mockService) Get(photoID string) (personID string, err error) {
	value, ok := m.Map.Get(photoID)
	if !ok {
		return "", AWS.ErrObjectNotFound
	}

	castedValue := value.(storedPhoto)

	return castedValue.personId, nil
}

func (m *mockService) Delete(photoID string) (err error) {
	if _, ok := m.Map.Get(photoID); !ok {
		return AWS.ErrObjectNotFound
	}

	m.Map.Remove(photoID)
	return nil
}
func (m *mockService) List(limit *int64) ([]*strfmt.UUID, error) {
	var keys []*strfmt.UUID

	lim := *limit
	it := m.Map.Iterator()
	for it.Next() && lim > 0 {
		k := it.Key().(string)
		var v = strfmt.UUID(k)
		keys = append(keys, &v)
		lim--
	}

	return keys, nil
}
func (m *mockService) Rekognize(sourceImage []byte) (_ AWS.MatchedFaces, err error) {
	it := m.Map.Iterator()
	for it.Next() {
		value := it.Value().(storedPhoto)
		if reflect.DeepEqual(value.body, sourceImage) {
			k := it.Key().(string)
			var v = strfmt.UUID(k)

			sim := 100.0
			return AWS.MatchedFaces{AWS.FaceMatch{Key: &v, Similarity: &sim}}, nil
		}
	}
	return AWS.MatchedFaces{AWS.FaceMatch{nil, nil}}, AWS.ErrObjectNotFound
}

func (m *mockService) ServiceName() string { return aws.SDKName }