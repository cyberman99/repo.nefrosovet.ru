package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

const (
	// PolicyCollectionName - is document collection name
	PolicyCollectionName = "policies"
)

type PolicyRepository interface {
	Insert(policy Policy) (string, error)
	GetPolicyByID(id string) (*Policy, error)
	GetPolicyByResourceMethodPath(resource, method, path string) (*Policy, error)
	Update(policy Policy) error
	Delete(id string) error
	GetLastPolicyTime() (time.Time, error)
	GetPolicies() ([]*Policy, error)
}

type policyRepo struct {
	ctx context.Context
	db  *mongo.Collection
}

func NewPolicyRepo(store Storer) PolicyRepository {
	pr := policyRepo{
		context.Background(),
		store.Collection(PolicyCollectionName),
	}
	err := pr.ensurePoliciesIndex()
	if err != nil {
		panic(err)
	}
	return &pr
}

// Policy is a struct with information
type Policy struct {
	ID                string    `bson:"id"`
	Stored            time.Time `bson:"stored"`
	Description       string    `bson:"description,omitempty"`
	Resource          string    `bson:"resource,omitempty"`
	Method            string    `bson:"method,omitempty"`
	Path              string    `bson:"path,omitempty"`
	BackendHost       string    `bson:"backendHost,omitempty"`
	BackendPath       string    `bson:"backendPath,omitempty"`
	Roles             []string  `bson:"roles,omitempty"`
	QueryStringParams []string  `bson:"querystring_params,omitempty"`
	HeadersToPass     []string  `bson:"headers_to_pass,omitempty"`
	KeyCache          int       `bson:"key_cache,omitempty"`
	Cache             bool      `bson:"cache,omitempty"`
}

// Insert inserts policy to DB
func (pr *policyRepo) Insert(policy Policy) (string, error) {
	// Generate UUID on insert
	if policy.ID == "" {
		policy.ID = uuid.New().String()
	}

	// Control save time
	policy.Stored = time.Now()
	_, err := pr.db.InsertOne(pr.ctx, policy)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"function": "Policy.Insert",
			"error":    err,
		}).Debug("mongo request failed")
		return "", err
	}
	return policy.ID, err
}

// Update rewrites policy on DB
func (pr *policyRepo) Update(policy Policy) error {
	// Control save time
	policy.Stored = time.Now()

	opts := new(options.FindOneAndUpdateOptions)
	opts.SetUpsert(false)
	opts.SetReturnDocument(options.After)

	err := pr.db.FindOneAndUpdate(
		pr.ctx,
		bson.M{"id": policy.ID},
		bson.M{"$set": policy},
		opts,
	).Err()

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"function": "Policy.Update",
			"error":    err,
		}).Debug("mongo request failed")
		return err
	}

	if len(policy.Roles) == 0 {
		return pr.db.FindOneAndUpdate(
			pr.ctx,
			bson.M{"id": policy.ID},
			bson.M{"$unset": bson.M{"roles": policy.Roles}},
		).Err()
	}

	return nil
}

// GetPolicyByID returns policy from DB
func (pr *policyRepo) GetPolicyByID(id string) (*Policy, error) {
	var policy Policy
	err := pr.db.FindOne(pr.ctx, bson.M{"id": id}).Decode(&policy)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"function": "GetPolicyByID",
			"error":    err,
		}).Debug("mongo request failed")

		return nil, err
	}

	return &policy, err
}

// GetPolicyByResourceMethodPath returns policy from DB
func (pr *policyRepo) GetPolicyByResourceMethodPath(resource, method, path string) (*Policy, error) {
	var policy Policy
	err := pr.db.FindOne(
		pr.ctx,
		bson.M{"resource": resource, "method": method, "path": path},
	).Decode(&policy)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"function": "GetPolicyByResourceMethodPath",
			"error":    err,
		}).Debug("mongo request failed")

		return nil, err
	}

	return &policy, err
}

// GetPolicies returns policies collection from DB
func (pr *policyRepo) GetPolicies() ([]*Policy, error) {
	var policies []*Policy
	cur, err := pr.db.Find(pr.ctx, bson.M{})

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"function": "PolicyCollection",
			"error":    err,
		}).Debug("mongo request failed")
		return nil, err
	}

	defer cur.Close(pr.ctx)

	for cur.Next(pr.ctx) {
		var policy Policy
		if err = cur.Decode(&policy); err != nil {
			return nil, err
		}

		policies = append(policies, &policy)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}

	if len(policies) == 0 {
		return policies, mongo.ErrNoDocuments
	}

	return policies, err
}

// Delete exterminates policy
func (pr *policyRepo) Delete(id string) error {
	err := pr.db.FindOneAndDelete(pr.ctx, bson.M{"id": id}).Err()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"function": "Policy.Delete",
			"error":    err,
		}).Debug("mongo request failed")
	}
	return err
}

// GetLastPolicyTime returns last policy write time
func (pr *policyRepo) GetLastPolicyTime() (time.Time, error) {
	opts := new(options.FindOptions)
	opts.SetSort(bson.D{{"stored", -1}})

	var lastUpdated Policy
	cur, err := pr.db.Find(pr.ctx, bson.M{}, opts)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"function": "GetLastPolicyTime",
			"error":    err,
		}).Debug("mongo request failed")

		return time.Unix(0, 0), err
	}
	defer cur.Close(pr.ctx)

	if cur.Next(pr.ctx) {
		err = cur.Decode(&lastUpdated)
		if err != nil {
			return time.Unix(0, 0), err
		}
		return lastUpdated.Stored, err
	}
	return time.Unix(0, 0), mongo.ErrNoDocuments
}

func (pr *policyRepo) ensurePoliciesIndex() error {
	var (
		noErr     = true
		index     mongo.IndexModel
		doc       bsonx.Doc
		err       error
		indexName = "unique_resource_method_path"
		cur       *mongo.Cursor
	)
	defer func() {
		if !noErr {
			pr.db.Indexes().DropAll(pr.ctx) // rollback
		}
	}()

	cur, err = pr.db.Indexes().List(pr.ctx, &options.ListIndexesOptions{})
	if err != nil {
		noErr = false
		return err
	}
	for cur.Next(pr.ctx) {
		err = cur.Decode(&doc)
		if err != nil {
			noErr = false
			return err
		}
		for _, elem := range doc {
			if elem.Value.String() == indexName {
				return nil
			}
		}
	}
	keys := bsonx.Doc{
		{Key: "resource", Value: bsonx.Int32(int32(1))},
		{Key: "method", Value: bsonx.Int32(int32(1))},
		{Key: "path", Value: bsonx.Int32(int32(1))},
	}
	index.Keys = keys
	index.Options = new(options.IndexOptions)
	index.Options.SetName(indexName)
	index.Options.SetUnique(true)
	if _, err := pr.db.Indexes().CreateOne(pr.ctx, index); err != nil {
		return err
	}
	return nil
}
