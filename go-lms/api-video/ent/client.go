// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"repo.nefrosovet.ru/go-lms/api-video/ent/migrate"

	"repo.nefrosovet.ru/go-lms/api-video/ent/accountkey"
	"repo.nefrosovet.ru/go-lms/api-video/ent/subscriber"
	"repo.nefrosovet.ru/go-lms/api-video/ent/user"
	"repo.nefrosovet.ru/go-lms/api-video/ent/useraccount"
	"repo.nefrosovet.ru/go-lms/api-video/ent/webinaruser"

	"github.com/facebookincubator/ent/dialect"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// AccountKey is the client for interacting with the AccountKey builders.
	AccountKey *AccountKeyClient
	// Subscriber is the client for interacting with the Subscriber builders.
	Subscriber *SubscriberClient
	// User is the client for interacting with the User builders.
	User *UserClient
	// UserAccount is the client for interacting with the UserAccount builders.
	UserAccount *UserAccountClient
	// WebinarUser is the client for interacting with the WebinarUser builders.
	WebinarUser *WebinarUserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	c := config{log: log.Println}
	c.options(opts...)
	return &Client{
		config:      c,
		Schema:      migrate.NewSchema(c.driver),
		AccountKey:  NewAccountKeyClient(c),
		Subscriber:  NewSubscriberClient(c),
		User:        NewUserClient(c),
		UserAccount: NewUserAccountClient(c),
		WebinarUser: NewWebinarUserClient(c),
	}
}

// Open opens a connection to the database specified by the driver name and a
// driver-specific data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: tx, log: c.log, debug: c.debug}
	return &Tx{
		config:      cfg,
		AccountKey:  NewAccountKeyClient(cfg),
		Subscriber:  NewSubscriberClient(cfg),
		User:        NewUserClient(cfg),
		UserAccount: NewUserAccountClient(cfg),
		WebinarUser: NewWebinarUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		AccountKey.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := config{driver: dialect.Debug(c.driver, c.log), log: c.log, debug: true}
	return &Client{
		config:      cfg,
		Schema:      migrate.NewSchema(cfg.driver),
		AccountKey:  NewAccountKeyClient(cfg),
		Subscriber:  NewSubscriberClient(cfg),
		User:        NewUserClient(cfg),
		UserAccount: NewUserAccountClient(cfg),
		WebinarUser: NewWebinarUserClient(cfg),
	}
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// AccountKeyClient is a client for the AccountKey schema.
type AccountKeyClient struct {
	config
}

// NewAccountKeyClient returns a client for the AccountKey from the given config.
func NewAccountKeyClient(c config) *AccountKeyClient {
	return &AccountKeyClient{config: c}
}

// Create returns a create builder for AccountKey.
func (c *AccountKeyClient) Create() *AccountKeyCreate {
	return &AccountKeyCreate{config: c.config}
}

// Update returns an update builder for AccountKey.
func (c *AccountKeyClient) Update() *AccountKeyUpdate {
	return &AccountKeyUpdate{config: c.config}
}

// UpdateOne returns an update builder for the given entity.
func (c *AccountKeyClient) UpdateOne(ak *AccountKey) *AccountKeyUpdateOne {
	return c.UpdateOneID(ak.ID)
}

// UpdateOneID returns an update builder for the given id.
func (c *AccountKeyClient) UpdateOneID(id int) *AccountKeyUpdateOne {
	return &AccountKeyUpdateOne{config: c.config, id: id}
}

// Delete returns a delete builder for AccountKey.
func (c *AccountKeyClient) Delete() *AccountKeyDelete {
	return &AccountKeyDelete{config: c.config}
}

// DeleteOne returns a delete builder for the given entity.
func (c *AccountKeyClient) DeleteOne(ak *AccountKey) *AccountKeyDeleteOne {
	return c.DeleteOneID(ak.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *AccountKeyClient) DeleteOneID(id int) *AccountKeyDeleteOne {
	return &AccountKeyDeleteOne{c.Delete().Where(accountkey.ID(id))}
}

// Create returns a query builder for AccountKey.
func (c *AccountKeyClient) Query() *AccountKeyQuery {
	return &AccountKeyQuery{config: c.config}
}

// Get returns a AccountKey entity by its id.
func (c *AccountKeyClient) Get(ctx context.Context, id int) (*AccountKey, error) {
	return c.Query().Where(accountkey.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AccountKeyClient) GetX(ctx context.Context, id int) *AccountKey {
	ak, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return ak
}

// QueryUseraccount queries the useraccount edge of a AccountKey.
func (c *AccountKeyClient) QueryUseraccount(ak *AccountKey) *UserAccountQuery {
	query := &UserAccountQuery{config: c.config}
	id := ak.ID
	step := sqlgraph.NewStep(
		sqlgraph.From(accountkey.Table, accountkey.FieldID, id),
		sqlgraph.To(useraccount.Table, useraccount.FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, accountkey.UseraccountTable, accountkey.UseraccountColumn),
	)
	query.sql = sqlgraph.Neighbors(ak.driver.Dialect(), step)

	return query
}

// SubscriberClient is a client for the Subscriber schema.
type SubscriberClient struct {
	config
}

// NewSubscriberClient returns a client for the Subscriber from the given config.
func NewSubscriberClient(c config) *SubscriberClient {
	return &SubscriberClient{config: c}
}

// Create returns a create builder for Subscriber.
func (c *SubscriberClient) Create() *SubscriberCreate {
	return &SubscriberCreate{config: c.config}
}

// Update returns an update builder for Subscriber.
func (c *SubscriberClient) Update() *SubscriberUpdate {
	return &SubscriberUpdate{config: c.config}
}

// UpdateOne returns an update builder for the given entity.
func (c *SubscriberClient) UpdateOne(s *Subscriber) *SubscriberUpdateOne {
	return c.UpdateOneID(s.ID)
}

// UpdateOneID returns an update builder for the given id.
func (c *SubscriberClient) UpdateOneID(id int) *SubscriberUpdateOne {
	return &SubscriberUpdateOne{config: c.config, id: id}
}

// Delete returns a delete builder for Subscriber.
func (c *SubscriberClient) Delete() *SubscriberDelete {
	return &SubscriberDelete{config: c.config}
}

// DeleteOne returns a delete builder for the given entity.
func (c *SubscriberClient) DeleteOne(s *Subscriber) *SubscriberDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *SubscriberClient) DeleteOneID(id int) *SubscriberDeleteOne {
	return &SubscriberDeleteOne{c.Delete().Where(subscriber.ID(id))}
}

// Create returns a query builder for Subscriber.
func (c *SubscriberClient) Query() *SubscriberQuery {
	return &SubscriberQuery{config: c.config}
}

// Get returns a Subscriber entity by its id.
func (c *SubscriberClient) Get(ctx context.Context, id int) (*Subscriber, error) {
	return c.Query().Where(subscriber.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *SubscriberClient) GetX(ctx context.Context, id int) *Subscriber {
	s, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return s
}

// QueryUser queries the user edge of a Subscriber.
func (c *SubscriberClient) QueryUser(s *Subscriber) *UserQuery {
	query := &UserQuery{config: c.config}
	id := s.ID
	step := sqlgraph.NewStep(
		sqlgraph.From(subscriber.Table, subscriber.FieldID, id),
		sqlgraph.To(user.Table, user.FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, subscriber.UserTable, subscriber.UserColumn),
	)
	query.sql = sqlgraph.Neighbors(s.driver.Dialect(), step)

	return query
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Create returns a create builder for User.
func (c *UserClient) Create() *UserCreate {
	return &UserCreate{config: c.config}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	return &UserUpdate{config: c.config}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	return c.UpdateOneID(u.ID)
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	return &UserUpdateOne{config: c.config, id: id}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	return &UserDelete{config: c.config}
}

// DeleteOne returns a delete builder for the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	return &UserDeleteOne{c.Delete().Where(user.ID(id))}
}

// Create returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{config: c.config}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	u, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return u
}

// QuerySubscriber queries the subscriber edge of a User.
func (c *UserClient) QuerySubscriber(u *User) *SubscriberQuery {
	query := &SubscriberQuery{config: c.config}
	id := u.ID
	step := sqlgraph.NewStep(
		sqlgraph.From(user.Table, user.FieldID, id),
		sqlgraph.To(subscriber.Table, subscriber.FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, user.SubscriberTable, user.SubscriberColumn),
	)
	query.sql = sqlgraph.Neighbors(u.driver.Dialect(), step)

	return query
}

// QueryUseraccount queries the useraccount edge of a User.
func (c *UserClient) QueryUseraccount(u *User) *UserAccountQuery {
	query := &UserAccountQuery{config: c.config}
	id := u.ID
	step := sqlgraph.NewStep(
		sqlgraph.From(user.Table, user.FieldID, id),
		sqlgraph.To(useraccount.Table, useraccount.FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, user.UseraccountTable, user.UseraccountColumn),
	)
	query.sql = sqlgraph.Neighbors(u.driver.Dialect(), step)

	return query
}

// UserAccountClient is a client for the UserAccount schema.
type UserAccountClient struct {
	config
}

// NewUserAccountClient returns a client for the UserAccount from the given config.
func NewUserAccountClient(c config) *UserAccountClient {
	return &UserAccountClient{config: c}
}

// Create returns a create builder for UserAccount.
func (c *UserAccountClient) Create() *UserAccountCreate {
	return &UserAccountCreate{config: c.config}
}

// Update returns an update builder for UserAccount.
func (c *UserAccountClient) Update() *UserAccountUpdate {
	return &UserAccountUpdate{config: c.config}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserAccountClient) UpdateOne(ua *UserAccount) *UserAccountUpdateOne {
	return c.UpdateOneID(ua.ID)
}

// UpdateOneID returns an update builder for the given id.
func (c *UserAccountClient) UpdateOneID(id int) *UserAccountUpdateOne {
	return &UserAccountUpdateOne{config: c.config, id: id}
}

// Delete returns a delete builder for UserAccount.
func (c *UserAccountClient) Delete() *UserAccountDelete {
	return &UserAccountDelete{config: c.config}
}

// DeleteOne returns a delete builder for the given entity.
func (c *UserAccountClient) DeleteOne(ua *UserAccount) *UserAccountDeleteOne {
	return c.DeleteOneID(ua.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *UserAccountClient) DeleteOneID(id int) *UserAccountDeleteOne {
	return &UserAccountDeleteOne{c.Delete().Where(useraccount.ID(id))}
}

// Create returns a query builder for UserAccount.
func (c *UserAccountClient) Query() *UserAccountQuery {
	return &UserAccountQuery{config: c.config}
}

// Get returns a UserAccount entity by its id.
func (c *UserAccountClient) Get(ctx context.Context, id int) (*UserAccount, error) {
	return c.Query().Where(useraccount.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserAccountClient) GetX(ctx context.Context, id int) *UserAccount {
	ua, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return ua
}

// QueryUser queries the user edge of a UserAccount.
func (c *UserAccountClient) QueryUser(ua *UserAccount) *UserQuery {
	query := &UserQuery{config: c.config}
	id := ua.ID
	step := sqlgraph.NewStep(
		sqlgraph.From(useraccount.Table, useraccount.FieldID, id),
		sqlgraph.To(user.Table, user.FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, useraccount.UserTable, useraccount.UserColumn),
	)
	query.sql = sqlgraph.Neighbors(ua.driver.Dialect(), step)

	return query
}

// QueryAccountkeys queries the accountkeys edge of a UserAccount.
func (c *UserAccountClient) QueryAccountkeys(ua *UserAccount) *AccountKeyQuery {
	query := &AccountKeyQuery{config: c.config}
	id := ua.ID
	step := sqlgraph.NewStep(
		sqlgraph.From(useraccount.Table, useraccount.FieldID, id),
		sqlgraph.To(accountkey.Table, accountkey.FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, useraccount.AccountkeysTable, useraccount.AccountkeysColumn),
	)
	query.sql = sqlgraph.Neighbors(ua.driver.Dialect(), step)

	return query
}

// WebinarUserClient is a client for the WebinarUser schema.
type WebinarUserClient struct {
	config
}

// NewWebinarUserClient returns a client for the WebinarUser from the given config.
func NewWebinarUserClient(c config) *WebinarUserClient {
	return &WebinarUserClient{config: c}
}

// Create returns a create builder for WebinarUser.
func (c *WebinarUserClient) Create() *WebinarUserCreate {
	return &WebinarUserCreate{config: c.config}
}

// Update returns an update builder for WebinarUser.
func (c *WebinarUserClient) Update() *WebinarUserUpdate {
	return &WebinarUserUpdate{config: c.config}
}

// UpdateOne returns an update builder for the given entity.
func (c *WebinarUserClient) UpdateOne(wu *WebinarUser) *WebinarUserUpdateOne {
	return c.UpdateOneID(wu.ID)
}

// UpdateOneID returns an update builder for the given id.
func (c *WebinarUserClient) UpdateOneID(id int) *WebinarUserUpdateOne {
	return &WebinarUserUpdateOne{config: c.config, id: id}
}

// Delete returns a delete builder for WebinarUser.
func (c *WebinarUserClient) Delete() *WebinarUserDelete {
	return &WebinarUserDelete{config: c.config}
}

// DeleteOne returns a delete builder for the given entity.
func (c *WebinarUserClient) DeleteOne(wu *WebinarUser) *WebinarUserDeleteOne {
	return c.DeleteOneID(wu.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *WebinarUserClient) DeleteOneID(id int) *WebinarUserDeleteOne {
	return &WebinarUserDeleteOne{c.Delete().Where(webinaruser.ID(id))}
}

// Create returns a query builder for WebinarUser.
func (c *WebinarUserClient) Query() *WebinarUserQuery {
	return &WebinarUserQuery{config: c.config}
}

// Get returns a WebinarUser entity by its id.
func (c *WebinarUserClient) Get(ctx context.Context, id int) (*WebinarUser, error) {
	return c.Query().Where(webinaruser.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *WebinarUserClient) GetX(ctx context.Context, id int) *WebinarUser {
	wu, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return wu
}
