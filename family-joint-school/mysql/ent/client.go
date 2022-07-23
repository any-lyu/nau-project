// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"family-joint-school/mysql/ent/migrate"

	"family-joint-school/mysql/ent/account"
	"family-joint-school/mysql/ent/class"
	"family-joint-school/mysql/ent/homework"
	"family-joint-school/mysql/ent/notice"
	"family-joint-school/mysql/ent/user"
	"family-joint-school/mysql/ent/userhomework"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Account is the client for interacting with the Account builders.
	Account *AccountClient
	// Class is the client for interacting with the Class builders.
	Class *ClassClient
	// Homework is the client for interacting with the Homework builders.
	Homework *HomeworkClient
	// Notice is the client for interacting with the Notice builders.
	Notice *NoticeClient
	// User is the client for interacting with the User builders.
	User *UserClient
	// UserHomeWork is the client for interacting with the UserHomeWork builders.
	UserHomeWork *UserHomeWorkClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Account = NewAccountClient(c.config)
	c.Class = NewClassClient(c.config)
	c.Homework = NewHomeworkClient(c.config)
	c.Notice = NewNoticeClient(c.config)
	c.User = NewUserClient(c.config)
	c.UserHomeWork = NewUserHomeWorkClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
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

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:          ctx,
		config:       cfg,
		Account:      NewAccountClient(cfg),
		Class:        NewClassClient(cfg),
		Homework:     NewHomeworkClient(cfg),
		Notice:       NewNoticeClient(cfg),
		User:         NewUserClient(cfg),
		UserHomeWork: NewUserHomeWorkClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:          ctx,
		config:       cfg,
		Account:      NewAccountClient(cfg),
		Class:        NewClassClient(cfg),
		Homework:     NewHomeworkClient(cfg),
		Notice:       NewNoticeClient(cfg),
		User:         NewUserClient(cfg),
		UserHomeWork: NewUserHomeWorkClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Account.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Account.Use(hooks...)
	c.Class.Use(hooks...)
	c.Homework.Use(hooks...)
	c.Notice.Use(hooks...)
	c.User.Use(hooks...)
	c.UserHomeWork.Use(hooks...)
}

// AccountClient is a client for the Account schema.
type AccountClient struct {
	config
}

// NewAccountClient returns a client for the Account from the given config.
func NewAccountClient(c config) *AccountClient {
	return &AccountClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `account.Hooks(f(g(h())))`.
func (c *AccountClient) Use(hooks ...Hook) {
	c.hooks.Account = append(c.hooks.Account, hooks...)
}

// Create returns a builder for creating a Account entity.
func (c *AccountClient) Create() *AccountCreate {
	mutation := newAccountMutation(c.config, OpCreate)
	return &AccountCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Account entities.
func (c *AccountClient) CreateBulk(builders ...*AccountCreate) *AccountCreateBulk {
	return &AccountCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Account.
func (c *AccountClient) Update() *AccountUpdate {
	mutation := newAccountMutation(c.config, OpUpdate)
	return &AccountUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AccountClient) UpdateOne(a *Account) *AccountUpdateOne {
	mutation := newAccountMutation(c.config, OpUpdateOne, withAccount(a))
	return &AccountUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AccountClient) UpdateOneID(id uint64) *AccountUpdateOne {
	mutation := newAccountMutation(c.config, OpUpdateOne, withAccountID(id))
	return &AccountUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Account.
func (c *AccountClient) Delete() *AccountDelete {
	mutation := newAccountMutation(c.config, OpDelete)
	return &AccountDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *AccountClient) DeleteOne(a *Account) *AccountDeleteOne {
	return c.DeleteOneID(a.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *AccountClient) DeleteOneID(id uint64) *AccountDeleteOne {
	builder := c.Delete().Where(account.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AccountDeleteOne{builder}
}

// Query returns a query builder for Account.
func (c *AccountClient) Query() *AccountQuery {
	return &AccountQuery{
		config: c.config,
	}
}

// Get returns a Account entity by its id.
func (c *AccountClient) Get(ctx context.Context, id uint64) (*Account, error) {
	return c.Query().Where(account.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AccountClient) GetX(ctx context.Context, id uint64) *Account {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *AccountClient) Hooks() []Hook {
	return c.hooks.Account
}

// ClassClient is a client for the Class schema.
type ClassClient struct {
	config
}

// NewClassClient returns a client for the Class from the given config.
func NewClassClient(c config) *ClassClient {
	return &ClassClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `class.Hooks(f(g(h())))`.
func (c *ClassClient) Use(hooks ...Hook) {
	c.hooks.Class = append(c.hooks.Class, hooks...)
}

// Create returns a builder for creating a Class entity.
func (c *ClassClient) Create() *ClassCreate {
	mutation := newClassMutation(c.config, OpCreate)
	return &ClassCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Class entities.
func (c *ClassClient) CreateBulk(builders ...*ClassCreate) *ClassCreateBulk {
	return &ClassCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Class.
func (c *ClassClient) Update() *ClassUpdate {
	mutation := newClassMutation(c.config, OpUpdate)
	return &ClassUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ClassClient) UpdateOne(cl *Class) *ClassUpdateOne {
	mutation := newClassMutation(c.config, OpUpdateOne, withClass(cl))
	return &ClassUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ClassClient) UpdateOneID(id uint64) *ClassUpdateOne {
	mutation := newClassMutation(c.config, OpUpdateOne, withClassID(id))
	return &ClassUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Class.
func (c *ClassClient) Delete() *ClassDelete {
	mutation := newClassMutation(c.config, OpDelete)
	return &ClassDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ClassClient) DeleteOne(cl *Class) *ClassDeleteOne {
	return c.DeleteOneID(cl.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *ClassClient) DeleteOneID(id uint64) *ClassDeleteOne {
	builder := c.Delete().Where(class.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ClassDeleteOne{builder}
}

// Query returns a query builder for Class.
func (c *ClassClient) Query() *ClassQuery {
	return &ClassQuery{
		config: c.config,
	}
}

// Get returns a Class entity by its id.
func (c *ClassClient) Get(ctx context.Context, id uint64) (*Class, error) {
	return c.Query().Where(class.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ClassClient) GetX(ctx context.Context, id uint64) *Class {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *ClassClient) Hooks() []Hook {
	return c.hooks.Class
}

// HomeworkClient is a client for the Homework schema.
type HomeworkClient struct {
	config
}

// NewHomeworkClient returns a client for the Homework from the given config.
func NewHomeworkClient(c config) *HomeworkClient {
	return &HomeworkClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `homework.Hooks(f(g(h())))`.
func (c *HomeworkClient) Use(hooks ...Hook) {
	c.hooks.Homework = append(c.hooks.Homework, hooks...)
}

// Create returns a builder for creating a Homework entity.
func (c *HomeworkClient) Create() *HomeworkCreate {
	mutation := newHomeworkMutation(c.config, OpCreate)
	return &HomeworkCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Homework entities.
func (c *HomeworkClient) CreateBulk(builders ...*HomeworkCreate) *HomeworkCreateBulk {
	return &HomeworkCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Homework.
func (c *HomeworkClient) Update() *HomeworkUpdate {
	mutation := newHomeworkMutation(c.config, OpUpdate)
	return &HomeworkUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *HomeworkClient) UpdateOne(h *Homework) *HomeworkUpdateOne {
	mutation := newHomeworkMutation(c.config, OpUpdateOne, withHomework(h))
	return &HomeworkUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *HomeworkClient) UpdateOneID(id uint64) *HomeworkUpdateOne {
	mutation := newHomeworkMutation(c.config, OpUpdateOne, withHomeworkID(id))
	return &HomeworkUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Homework.
func (c *HomeworkClient) Delete() *HomeworkDelete {
	mutation := newHomeworkMutation(c.config, OpDelete)
	return &HomeworkDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *HomeworkClient) DeleteOne(h *Homework) *HomeworkDeleteOne {
	return c.DeleteOneID(h.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *HomeworkClient) DeleteOneID(id uint64) *HomeworkDeleteOne {
	builder := c.Delete().Where(homework.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &HomeworkDeleteOne{builder}
}

// Query returns a query builder for Homework.
func (c *HomeworkClient) Query() *HomeworkQuery {
	return &HomeworkQuery{
		config: c.config,
	}
}

// Get returns a Homework entity by its id.
func (c *HomeworkClient) Get(ctx context.Context, id uint64) (*Homework, error) {
	return c.Query().Where(homework.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *HomeworkClient) GetX(ctx context.Context, id uint64) *Homework {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *HomeworkClient) Hooks() []Hook {
	return c.hooks.Homework
}

// NoticeClient is a client for the Notice schema.
type NoticeClient struct {
	config
}

// NewNoticeClient returns a client for the Notice from the given config.
func NewNoticeClient(c config) *NoticeClient {
	return &NoticeClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `notice.Hooks(f(g(h())))`.
func (c *NoticeClient) Use(hooks ...Hook) {
	c.hooks.Notice = append(c.hooks.Notice, hooks...)
}

// Create returns a builder for creating a Notice entity.
func (c *NoticeClient) Create() *NoticeCreate {
	mutation := newNoticeMutation(c.config, OpCreate)
	return &NoticeCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Notice entities.
func (c *NoticeClient) CreateBulk(builders ...*NoticeCreate) *NoticeCreateBulk {
	return &NoticeCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Notice.
func (c *NoticeClient) Update() *NoticeUpdate {
	mutation := newNoticeMutation(c.config, OpUpdate)
	return &NoticeUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *NoticeClient) UpdateOne(n *Notice) *NoticeUpdateOne {
	mutation := newNoticeMutation(c.config, OpUpdateOne, withNotice(n))
	return &NoticeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *NoticeClient) UpdateOneID(id uint64) *NoticeUpdateOne {
	mutation := newNoticeMutation(c.config, OpUpdateOne, withNoticeID(id))
	return &NoticeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Notice.
func (c *NoticeClient) Delete() *NoticeDelete {
	mutation := newNoticeMutation(c.config, OpDelete)
	return &NoticeDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *NoticeClient) DeleteOne(n *Notice) *NoticeDeleteOne {
	return c.DeleteOneID(n.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *NoticeClient) DeleteOneID(id uint64) *NoticeDeleteOne {
	builder := c.Delete().Where(notice.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &NoticeDeleteOne{builder}
}

// Query returns a query builder for Notice.
func (c *NoticeClient) Query() *NoticeQuery {
	return &NoticeQuery{
		config: c.config,
	}
}

// Get returns a Notice entity by its id.
func (c *NoticeClient) Get(ctx context.Context, id uint64) (*Notice, error) {
	return c.Query().Where(notice.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *NoticeClient) GetX(ctx context.Context, id uint64) *Notice {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *NoticeClient) Hooks() []Hook {
	return c.hooks.Notice
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Create returns a builder for creating a User entity.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id uint64) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *UserClient) DeleteOneID(id uint64) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id uint64) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id uint64) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}

// UserHomeWorkClient is a client for the UserHomeWork schema.
type UserHomeWorkClient struct {
	config
}

// NewUserHomeWorkClient returns a client for the UserHomeWork from the given config.
func NewUserHomeWorkClient(c config) *UserHomeWorkClient {
	return &UserHomeWorkClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `userhomework.Hooks(f(g(h())))`.
func (c *UserHomeWorkClient) Use(hooks ...Hook) {
	c.hooks.UserHomeWork = append(c.hooks.UserHomeWork, hooks...)
}

// Create returns a builder for creating a UserHomeWork entity.
func (c *UserHomeWorkClient) Create() *UserHomeWorkCreate {
	mutation := newUserHomeWorkMutation(c.config, OpCreate)
	return &UserHomeWorkCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of UserHomeWork entities.
func (c *UserHomeWorkClient) CreateBulk(builders ...*UserHomeWorkCreate) *UserHomeWorkCreateBulk {
	return &UserHomeWorkCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for UserHomeWork.
func (c *UserHomeWorkClient) Update() *UserHomeWorkUpdate {
	mutation := newUserHomeWorkMutation(c.config, OpUpdate)
	return &UserHomeWorkUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserHomeWorkClient) UpdateOne(uhw *UserHomeWork) *UserHomeWorkUpdateOne {
	mutation := newUserHomeWorkMutation(c.config, OpUpdateOne, withUserHomeWork(uhw))
	return &UserHomeWorkUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserHomeWorkClient) UpdateOneID(id uint64) *UserHomeWorkUpdateOne {
	mutation := newUserHomeWorkMutation(c.config, OpUpdateOne, withUserHomeWorkID(id))
	return &UserHomeWorkUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for UserHomeWork.
func (c *UserHomeWorkClient) Delete() *UserHomeWorkDelete {
	mutation := newUserHomeWorkMutation(c.config, OpDelete)
	return &UserHomeWorkDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserHomeWorkClient) DeleteOne(uhw *UserHomeWork) *UserHomeWorkDeleteOne {
	return c.DeleteOneID(uhw.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *UserHomeWorkClient) DeleteOneID(id uint64) *UserHomeWorkDeleteOne {
	builder := c.Delete().Where(userhomework.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserHomeWorkDeleteOne{builder}
}

// Query returns a query builder for UserHomeWork.
func (c *UserHomeWorkClient) Query() *UserHomeWorkQuery {
	return &UserHomeWorkQuery{
		config: c.config,
	}
}

// Get returns a UserHomeWork entity by its id.
func (c *UserHomeWorkClient) Get(ctx context.Context, id uint64) (*UserHomeWork, error) {
	return c.Query().Where(userhomework.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserHomeWorkClient) GetX(ctx context.Context, id uint64) *UserHomeWork {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *UserHomeWorkClient) Hooks() []Hook {
	return c.hooks.UserHomeWork
}
