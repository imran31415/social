# Creating a social network

## Goal:
    
Create a social network application where users can have an account, make posts, see other users posts in a feed, and comment on those feed posts.  in addition users should be able to engage with each other in nested threads.   

## Stack:
- Go
- Mysql
- GRPC/ Protocol Buffers
- Vue
- Envoy
- Docker

## Plan

1. Create database schema and models
2. create API Interface
3. implement the API
3. Create javascript client to demonstrate integration


## Part 1: Create database schema:

We need the following initial schema to support our app:

```mysql
CREATE TABLE social_user
(
    id BIGINT auto_increment PRIMARY KEY,
    password VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL UNIQUE,
    ## implement schema later
    profile JSON NOT NULL
);

CREATE TABLE social_post
(
    id BIGINT auto_increment PRIMARY KEY,
    user_id BIGINT NOT NULL,
    FOREIGN KEY  (user_id) REFERENCES social_user(id),
    ## implement schema later
    content JSON NOT NULL
);

CREATE TABLE social_feed (
    id BIGINT auto_increment PRIMARY KEY,
    owner_id BIGINT NOT NULL,
    FOREIGN KEY  (owner_id) REFERENCES social_user(id),
    post_id BIGINT NOT NULL,
    FOREIGN KEY  (post_id) REFERENCES social_post(id)
);

CREATE TABLE social_comment (
    id BIGINT auto_increment PRIMARY KEY,
    post_id BIGINT NOT NULL,
    FOREIGN KEY  (post_id) REFERENCES social_post(id),
    parent_comment_id BIGINT NOT NULL,
    # implement schema later
    content JSON NOT NULL
);
```
The corresponding models code to represent these items in go is:

```go
package db

import "encoding/json"

type User struct {
	Id       int64           `db:"id"`
	Password string          `db:"password"`
	Username string          `db:"username"`
	Profile  json.RawMessage `db:"profile"`
}

type Post struct {
	Id      int64           `db:"id"`
	UserId  int64           `db:"user_id"`
	Content json.RawMessage `db:"content"`
}

type Feed struct {
	Id      int64 `db:"id"`
	OwnerId int64 `db:"owner_id"`
	PostId  int64 `db:"post_id"`
}

type Comment struct {
	Id              int64           `db:"id"`
	PostId          int64           `db:"post_id"`
	ParentCommentId int64           `db:"parent_comment_id"`
	Content         json.RawMessage `db:"content"`
}
```

## Part 2: Create APIs
With the db schema and go models in place, we can define our protobuf file API for the CRUD methods for the db:
```protobuf

service Social {
rpc GetUser(GetUserReq) returns (User){}
rpc CreateUser(CreateUserReq) returns(User){}

rpc GetFeed(GetFeedReq) returns (Feed){}
rpc CreateFeedItem(CreateFeedItemReq) returns(FeedItem){}

rpc GetPost(GetPostsReq) returns (Posts){}
rpc CreatePost(CreatePostReq) returns (Post){}

rpc CreateComment(CreateCommentReq) returns (Comment){}
rpc GetComments(GetCommentsReq) returns(Comments){}
}


message User {
  int64 id = 1;
  string user_name = 2;
  // profile is json string;
  string profile = 3;
}

message Post {
  int64 id = 1;
  // content is a string of JSON
  string content = 2;
}

message Posts {
  repeated Post items =1;
}

message FeedItem {
  int64 id = 1;
  // content is a string of JSON
  string content = 2;
}



message Feed {
  repeated FeedItem items =1;
}

message GetUserReq {
  oneof get_by {
    string user_name =1;
    int64 id = 2;
  }
}

message Comment {
  int64 id = 1;
  // content is a string of JSON
  string content = 2;
  int64 parent_id = 3;
  int64 post_id = 4;
}

message Comments {
  repeated Comments items = 1;
}

message GetPostsReq {
  repeated int64 ids = 2;
  enum GetPostsIdType {
    GetPostsIdType_user =0;
    GetPostsIdType_post =1;
  }
  GetPostsIdType get_by = 3;
}

message GetFeedReq {
  int64 owner_id = 1;
}

message GetCommentsReq {
  int64 post_id = 1;
}

message CreateUserReq {
  string user_name = 1;
  string password = 2;
  // profile is a string of JSON
  string profile = 3;
}

message CreatePostReq {
  int64 user_id = 1;
  // content is a string of JSON
  string content = 2;
}

message CreateFeedItemReq {
  int64 owner_id = 1;
  int64 post_id = 2;
}

message CreateCommentReq {
  int64 post_id = 1;
  // if parent_comment_id = 0, then the comment is a root comment on the post
  int64 parent_comment_id = 2;
  // content is a string of JSON
  string content = 3;
}
```
## Part 3: Implement the API

With our protofile in place, we can generate the GRPC and protobuf message types. 
In addition we can autogenerate documentation on our API
```bash
protoc -I protos/ protos/social.proto --go_out=plugins=grpc:protos
protoc --proto_path=protos --js_out=import_style=commonjs,binary:../http/static/js --grpc-web_out=import_style=commonjs,mode=grpcwebtext:../http/static/js protos/social.proto 
protoc --doc_out=protos  protos/social.proto 
```

This should generate a social.pb as well as a documentation HTML file that you can load in the browser.

Next we want to implement the Go grpc server code to use our database and protobuf definitions 

The minimal implementation of the grpc server is as follows:

```go


type Server struct {
}

func (s *Server) GetUser(ctx context.Context, req *pb.GetUserReq)(*pb.User, error) {
    return nil, nil
}
func (s *Server) CreateUser(ctx context.Context, req *pb.CreateUserReq)(*pb.User, error) {
    return nil, nil
}

func (s *Server) GetPosts(ctx context.Context, req *pb.GetPostsReq)(*pb.Posts, error) {
    return nil, nil
}
func (s *Server) CreatePost(ctx context.Context, req *pb.CreatePostReq)(*pb.Post, error) {
    return nil, nil
}

func (s *Server) GetFeed(ctx context.Context, req *pb.GetFeedReq)(*pb.Feed, error) {
    return nil, nil
}
func (s *Server) CreateFeedItem(ctx context.Context, req *pb.CreateFeedItemReq)(*pb.FeedItem, error) {
    return nil, nil
}

func (s *Server) GetComments(ctx context.Context, req *pb.GetCommentsReq)(*pb.Comments, error) {
    return nil, nil
}
func (s *Server) CreateComment(ctx context.Context, req *pb.CreateCommentReq)(*pb.Comment, error) {
    return nil, nil
}


func Run(grpcPort string) error {
    lis, err := net.Listen("tcp", grpcPort)
    if err != nil {
        return fmt.Errorf("failed to listen, err: %s", err)
    }
    
    serv  := &Server{}
    s := grpc.NewServer()
    pb.RegisterSocialServer(s, serv)
    log.Println("Serving GRPC...")
    if err = s.Serve(lis); err != nil {
        log.Println("Err runnning GRPC server: ", err)
        return err
    }
    return nil
}



```

Now we can test that our grpc scaffolding is successful by running the GRPC server.
To do this we can create a new package called service and create a main.go file to run the Grpc server

```go
package main

import (
	"log"
	"social/app/grpc"
)

func main() {
	log.Println(grpc.Run("localhost:50053"))
}
```

This should output:

```bash
go run main.go 
2021/05/31 13:32:39 Serving GRPC...
```

This means our go GRPC server implementation is successful.  
Now we need to implement each RPC method to interact with the GRPC and return the protobuf messages defined.

To start with, implement Get and Create user methods:

This will look something like, note this is pseudo code:
```go
func (s *Server) CreateUser(ctx context.Context, req *pb.CreateUserReq)(*pb.User, error) {
	if err := repo.CreateUser(ctx, req); err != nil {
		return nil, err
	}
	user, err := repo.GetUser(ctx, TODO);
	if err != nil {
		return nil, err
	}
	return user.Proto(), err
}
```

However, we are unable to do this since we have no "repo" package to do DB queries.   

To create the repo logic first we need to setup a db user and password in our mysql db. 
For now we can just use a locally running MYSQL instance for development and create the credentials here
```mysql
create user 'social_user'@'localhost' identified by '13tg1t8bqfsa76u';

GRANT ALL PRIVILEGES ON *.* TO 'social_user'@'localhost';
```

Now we can create a repo/ package under db/ to hold some go code to connect to our database and run migrations according to our initial DB file.
Note this code will use some go mysql third party libraries to implement the migration

```go
package repo

import (
	"context"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	mysqldb "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
	"path/filepath"
	"time"
)
type Repo struct {
	Db     *sqlx.DB
	Config *RepoConfig
}

type RepoConfig struct {
	DbPass string
	DbUser string
	DbName string
	DbHost string
	DbPort string
}

func (c *RepoConfig) MigrationsPath()string {
	return filepath.FromSlash("../db/schema")
}


// NewRepo initializes the struct as well as connects to the database and performs the initial migrations.
func NewRepo(config *RepoConfig) (*Repo, error) {
	r := &Repo{
		Config: config,
	}
	if err := r.Migrate(); err != nil {
		return r, err
	}
	// if we get a migration err we should return r so we can cleanup if needed.
	return r, nil
}


func (r *Repo) ConnectDatabase() error {
	log.Printf("Connecting to DB Host: %s with user: %s, dbname: %s, port: %s", r.Config.DbHost, r.Config.DbUser, r.Config.DbName, r.Config.DbPort)
	var connectErr error
	// if 10149 (db does not exist) create it
	if r.Db, connectErr = sqlx.Connect("mysql", r.connectionString()); GetMysqlErrorNumber(connectErr) == 1049 {
		log.Println("DB does not exist, attempting to create it")
		createErr := r.createDatabase(r.Config.DbName)
		if createErr != nil {
			log.Printf("failed to connect to  database: %s, err: %s", r.Config.DbName, createErr)
			return createErr
		}
		log.Printf("successfully created database: %s", r.Config.DbName)
		r.Db, connectErr = sqlx.Connect("mysql", r.connectionString())
	}
	return connectErr
}


func (r *Repo) Ping() error {
	return r.Db.Ping()
}

func (r *Repo) connectionString() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&multiStatements=true", r.Config.DbUser, r.Config.DbPass, r.Config.DbHost, r.Config.DbPort, r.Config.DbName)
}

func (r *Repo) connectionStringNoDb() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/?parseTime=true&multiStatements=true", r.Config.DbUser, r.Config.DbPass, r.Config.DbHost)
}

func (r *Repo) createDatabase(dbName string) error {
	db, err := sqlx.Connect("mysql", r.connectionStringNoDb())
	if err != nil {
		log.Println(err)
		return err
	}
	db.MustExec("CREATE database if not exists " + dbName)
	return db.Close()
}

func GetMysqlErrorNumber(err error) uint16 {
	switch v := err.(type) {
	case *mysql.MySQLError:
		return v.Number
	}
	// not a mysql error
	return 0
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}


func (r *Repo) Migrate() error {
	if err := r.ConnectDatabase(); err != nil {
		log.Printf("Err connecting to the DB, err: %s", err)
		return err
	} else {
		log.Printf("Successfully connected to DB")
	}
	r.Db.SetMaxOpenConns(50)
	r.Db.SetMaxIdleConns(5)
	r.Db.SetConnMaxLifetime(2 * time.Minute)
	driver, err := mysqldb.WithInstance(r.Db.DB, &mysqldb.Config{})
	if err != nil {
		log.Println(err)
		return err
	}

	if exists, err := exists(r.Config.MigrationsPath()); err != nil {
		log.Println(err)
		return err

	} else if !exists {
		return fmt.Errorf("migrations path does not exist: %s", r.Config.MigrationsPath())
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://"+r.Config.MigrationsPath(),
		"mysql",
		driver,
	)

	if err != nil {
		log.Println(err)
		return err
	}

	// Migrate all the way up ...
	if err = m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Println(err)
		return err
	}

	return nil

}

func (r *Repo) BeginTx(ctx context.Context) (*sqlx.Tx, error) {
	return r.Db.BeginTxx(ctx, nil)
}


```

Now we can update our main service/main.go file to load our repo before running the GRPC server:

```go
func main() {
	// TODO: replace with env variables
	_, err := repo.NewRepo(&repo.RepoConfig{
		DbPass: "13tg1t8bqfsa76u",
		DbUser: "social_user",
		DbName: "social_test",
		DbHost: "localhost",
		DbPort: "3306",
	})
	if err != nil{
		log.Fatal(err)
		return
	}
	log.Println(grpc.Run("localhost:50053"))
}
```

when we run this with `go run main.go` from `/service` directory we should see:
```bash

go run main.go 
```
Output:
```
2021/05/31 13:56:51 Connecting to DB Host: localhost with user: social_user, dbname: social_test, port: 3306
2021/05/31 13:56:51 Successfully connected to DB
2021/05/31 13:56:51 Serving GRPC...

```

Now we can implement the repo Create and Get user methods.

First  create an Interface for the Repo methods so that we can pass the Repo interface to the grpc server:

in `repo.go`
```go
type Iface interface {
	InsertUser(p *User) (int64, error)
	GetUserById(id int64) (*User, error)
	GetUserByUserName(username string) (*User, error)
}
```
Next implement the user repo methods:

```go


func (r *Repo) GetUserByUserName(username string) (*User, error) {
	c := &User{}
	err := r.Db.Get(c, "SELECT * FROM social_user where username = ?", username)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *Repo) GetUserById(id int64) (*User, error) {
	c := &User{}
	err := r.Db.Get(c, "SELECT * FROM social_user where id = ?", id)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *Repo) InsertUser(p *User) (int64, error) {
	res, err := r.Db.NamedExec(`
		INSERT INTO social_user (
			password,
			username,
			profile
		) VALUES (
			:password,
			:username,
			:profile
		);`, p)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

```
Now we can include the repo interface as a member of the GRPC Server wrapper struct
```go
type Server struct {
	r repo.Iface
}

func NewServer(r repo.Iface)*Server {
	return &Server{r: r}
}
```

In our GRPC server, we can implement the User API requests as so:

```go
import (
	"context"
	pb "social/app/grpc/protos"
	"social/app/grpc/protos/serializers"
)

func (s *Server) CreateUser(ctx context.Context, req *pb.CreateUserReq)(*pb.User, error) {
	insertedId, err :=  s.r.InsertUser(serializers.CreateUserReq(req))
	if err != nil {
		return nil, err
	}
	user, err := s.r.GetUserById(insertedId)
	if err != nil {
		return nil, err
	}
	return serializers.User(user), nil
}
```

And GetUser, note: we leverage the protobuf oneof type to determine what field to search by in the repo. 

```go


func (s *Server) GetUser(ctx context.Context, req *pb.GetUserReq)(*pb.User, error) {
	switch req.GetGetBy().(type) {
	case *pb.GetUserReq_Id:
		id := req.GetId()
		user, err := s.r.GetUserById(id)
		if err != nil {
			return nil, err
		}
		return serializers.User(user), nil
	case *pb.GetUserReq_UserName:
		name := req.GetUserName()
		user, err := s.r.GetUserByUserName(name)
		if err != nil {
			return nil, err
		}
		return serializers.User(user), nil
	default:
		return nil, fmt.Errorf("unknown req.GetGetBy")
	}
}
```

We also define some repo -> api  serializer helper methods in protos/serializers 
```go
func User(user *repo.User)*pb.User {
	return &pb.User{
		Id:      	user.Id,
		UserName:  user.Username,
		Profile:  string(user.Profile),
	}
}

func CreateUserReq(req *pb.CreateUserReq) *repo.User {
	return &repo.User{
		Password: req.GetPassword(),
		Username: req.GetUserName(),
		Profile:  []byte(req.GetProfile()),
	}
}
```

Now our grpc server should be able to successfully integrate with the Database in the user Get and Create calls.
To validate this we can write a test in the grpc folder to run the GRPC server, and call the methods

Before writing the test itself we need a helper to be able to create isolated databases for test purposes
Define a grpc/test_helpers/helpers.go file with code to create a test database and clean up:

```go
package test_helpers

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"social/repo"
	"strings"
)

const (

	DbUser = "social_test"
	DbPass = "social_test"
)

func SetupDbForTest(dbName string, sqlFile string, migrationsPath string) (*repo.Repo, error) {
	if !strings.HasPrefix(dbName, "test") {
		return nil, fmt.Errorf("name: %s does not begin with test", dbName)
	}
	// cleanupAndErr ensures that any errs encountered result in an attempt to clean up the database if created.
	// this is helpful as callers of SetupDbForTest can assert this function returned NoError and if there is an err,
	// do not need to worry about calling defer() to cleanup the DB.
	cleanupAndErr := func(err error, r *repo.Repo) (*repo.Repo, error) {
		fmt.Printf("encountered unexpected error during test database setup, err %s:", err.Error())
		DeleteTestDatabase(dbName, r)
		return nil, err
	}
	type testEnv struct {
		config             *repo.RepoConfig
		migrationsPathEnv  string
		insertsToApplyPath string
	}

	tEnv := &testEnv{
		config: &repo.RepoConfig{
			DbPass: DbUser,
			DbUser: DbPass,
			DbName: dbName,
			DbHost: "localhost",
			DbPort: "3306",
		},
		migrationsPathEnv:migrationsPath,
		// some test data for a product
		insertsToApplyPath: sqlFile,
	}
	// create repo with input test config
	// Note if this fails with access issue ensure access privileges of the mysql user
	r, err := repo.NewRepo(tEnv.config)
	if err != nil {
		return cleanupAndErr(err, r)
	}

	err = r.Ping()
	if err != nil {
		return cleanupAndErr(err, r)
	}

	// insert test data if path specified
	if tEnv.insertsToApplyPath != "" {
		err = executeSqlFromFile(r, tEnv.insertsToApplyPath)
		if err != nil {
			return cleanupAndErr(err, r)
		}
	}

	return r, nil
}

func DeleteTestDatabase(dbName string, r *repo.Repo) error {
	// double check we are not deleting any dbs with the same name as production
	// tests should create their own database names for each test method
	if !strings.HasPrefix(dbName, "test") {
		return fmt.Errorf("can not delete db: %s, invalid name", dbName)
	}

	if _, err := r.Db.Exec("DROP DATABASE " + dbName); err != nil {
		fmt.Printf("err cleaning up test database, err: %s", err.Error())
		return err
	}

	return nil
}

// executeSqlFromFile: a utility so each database created for test can execute some sql to generate test data
// each test should use its own test data and database so each test is independent
func executeSqlFromFile(r *repo.Repo, p string) error {
	path := filepath.Join(p)
	c, ioErr := ioutil.ReadFile(path)
	if ioErr != nil {
		return ioErr
	}
	sql := string(c)
	_, err := r.Db.Exec(sql)
	return err
}
```

With the above code to enable us to create a database per test case and clean it up,
we can write the following test in `grpc/user_test`


```go
package grpc_test

import (
	"context"
	"github.com/stretchr/testify/require"
	"social/app/grpc"
	pb "social/app/grpc/protos"
	"social/app/grpc/test_helpers"
	"testing"
)

func Test_CreateUser(t *testing.T) {
	t.Parallel()
	type test struct {
		name       string
		testDbName string
		in         *pb.CreateUserReq
		exp        *pb.User
		expErr     error
	}
	tests := []test{
		{
			name:       "CreateUser successfully creates a user",
			testDbName: "test_CreateUser_1",
			in: &pb.CreateUserReq{
				UserName: "foo",
				Password: "foopw",
			},
			exp: &pb.User{
				Id:       1,
				UserName: "foo",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, err := test_helpers.SetupDbForTest(tt.testDbName, "", "../db/schema")
			defer test_helpers.DeleteTestDatabase(tt.testDbName, r)
			require.NoError(t, err)
			require.NoError(t, r.Ping())
			s := grpc.NewServer(r)
			user, err := s.CreateUser(context.TODO(), &pb.CreateUserReq{
				UserName: tt.in.GetUserName(),
				Password: tt.in.GetPassword(),
			})
			if tt.expErr != nil {
				require.Error(t, err)
				require.Equal(t, tt.expErr, err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.exp, user)

		})
	}
}

```

We have successfully tested the GRPC server and database integration for both `GetUser` and `CreateUser` rpc methods (well at least the happy path :)

Following the same flow to implement the user methods we can implement the other API methods for Posts, Feed, and Comments:

Starting with Post:

`repo/post.go`

```go
package repo

import (
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
)

const (

	FieldNameSocialPostUserId SocialPostFieldName = "user_id"
	FieldNameSocialPostId SocialPostFieldName = "id"
)

type SocialPostFieldName string
type Post struct {
	Id      int64            `db:"id"`
	Content *json.RawMessage `db:"content"`
	UserId  int64            `db:"user_id"`
}

type Posts struct {
	Items []*Post
}

func (r *Repo) GetPostsIds(ids []int64, fieldName SocialPostFieldName) (*Posts, error) {
	posts := []*Post{}
	query, args, err := sqlx.In(fmt.Sprintf("SELECT * FROM social_post WHERE %s IN (?)", fieldName), ids)
	if err != nil {
		return nil, err
	}
	query = r.Db.Rebind(query)
	err = r.Db.Select(&posts, query, args...)

	if err != nil {
		return nil, err
	}
	return &Posts{
		Items: posts,
	}, nil
}

func (r *Repo) InsertPost(p *Post) (int64, error) {
	if p.Content == nil {
		return 0, fmt.Errorf("err, content is empty")
	}
	res, err := r.Db.NamedExec(`
		INSERT INTO social_post (
			content,
			user_id
		) VALUES (
			:content,
			:user_id
		);`, p)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

```

And the GRPC layer `grpc/post.go`
```go
package grpc

import (
	"context"
	"fmt"
	pb "social/app/grpc/protos"
	"social/app/grpc/protos/serializers"
	"social/repo"
)

func (s *Server) CreatePost(ctx context.Context, req *pb.CreatePostReq)(*pb.Post, error) {
	insertedId, err :=  s.r.InsertPost(serializers.CreatePostReq(req))
	if err != nil {
		return nil, err
	}
	posts, err := s.r.GetPostsIds([]int64{insertedId}, repo.FieldNameSocialPostId)
	if err != nil {
		return nil, err
	}
	if posts.Items == nil || len(posts.Items) != 1 {
		return nil, fmt.Errorf("err unexp items")
	}
	return serializers.Post(posts.Items[0]), nil
}


func (s *Server) GetPosts(ctx context.Context, req *pb.GetPostsReq)(*pb.Posts, error) {
	switch req.GetGetBy() {
	case pb.GetPostsReq_GetPostsIdType_post:
		posts, err := s.r.GetPostsIds(req.GetIds(), repo.FieldNameSocialPostId)
		if err != nil {
			return nil, err
		}
		return serializers.Posts(posts), nil
	case pb.GetPostsReq_GetPostsIdType_user:
		posts, err := s.r.GetPostsIds(req.GetIds(), repo.FieldNameSocialPostUserId)
		if err != nil {
			return nil, err
		}
		return serializers.Posts(posts), nil
	default:
		return nil, fmt.Errorf("invalid req")
	}

}
```

And the test for the Post RPCs, Note, this test includes testing the json serialization, as well as using the User repo methods to create the Post.user:

```go

package grpc_test

import (
	"context"
	"encoding/json"
	"github.com/stretchr/testify/require"
	"social/app/grpc"
	pb "social/app/grpc/protos"
	"social/app/grpc/test_helpers"
	"social/repo"
	"testing"
)

type ContentJson struct {
	PostBody string
}

func Test_CreatePost(t *testing.T) {
	t.Parallel()
	testContentStruct := &ContentJson{
		PostBody: "A test post",
	}
	testContent, err := json.Marshal(testContentStruct)
	require.NoError(t, err)
	testContentString := string(testContent)

	testUser := &repo.User{Username: "foo", Password:"foo_pw"}
	type test struct {
		name       string
		testDbName string
		in         *pb.CreatePostReq
		exp        *pb.Post
		expErr     error
	}
	tests := []test{
		{
			name:       "CreatePost successfully creates a post",
			testDbName: "test_CreatePost_1",
			in: &pb.CreatePostReq{
				Content:testContentString,
			},
			exp: &pb.Post{
				Id:       1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, err := test_helpers.SetupDbForTest(tt.testDbName, "", "../db/schema")
			defer test_helpers.DeleteTestDatabase(tt.testDbName, r)
			require.NoError(t, err)
			require.NoError(t, r.Ping())
			s := grpc.NewServer(r)
			// create a test user for the test
			user, err := s.CreateUser(context.TODO(), &pb.CreateUserReq{
				UserName: testUser.Username,
				Password: testUser.Password,
			})
			require.NoError(t, err)
			postIn := tt.in
			// set the post User id to the user that was created
			postIn.UserId = user.Id
			post, err := s.CreatePost(context.TODO(), postIn)
			if tt.expErr != nil {
				require.Error(t, err)
				require.Equal(t, tt.expErr, err)
				return
			}
			require.NoError(t, err)

			// validate the content was serialized correctly
			postContent := post.GetContent()
			post.Content = ""
			require.Equal(t, tt.exp, post)
			postContentSerialized := &ContentJson{}
			require.NoError(t, json.Unmarshal([]byte(postContent), postContentSerialized))
			require.Equal(t, testContentStruct, postContentSerialized)

		})
	}
}

```

Repeat the above with the Feed and Comment RPCs. 

Note: this code is omitted as it's very similar to the above, see code repo for the implementation. 

After implementation we should have the following repo and grpc folder files:

Repo:

```go
repo
├── comment.go
├── feed.go
├── go.mod
├── go.sum
├── models.go
├── post.go
├── repo.go
└── user.go

```

Grpc:

```go
grpc
├── comment.go
├── comment_test.go
├── feed.go
├── feed_test.go
├── go.mod
├── go.sum
├── grpc.go
├── post.go
├── post_test.go
├── protos
│   ├── index.html
│   ├── serializers
│   │   └── serializer.go
│   ├── social.pb.go
│   └── social.proto
├── test_helpers
│   └── helper.go
├── user.go
└── user_test.go


```

Now we should have a functioning and tested GRPC server. 

We can determine test coverage for the grpc and repo modules by running the following within each of the dirs:

```bash
go tool cover -func=coverage.out
```

```
social/app/grpc/comment.go:9:   CreateComment   71.4%
social/app/grpc/comment.go:21:  GetComments     0.0%
social/app/grpc/feed.go:9:      CreateFeedItem  71.4%
social/app/grpc/feed.go:21:     GetFeed         0.0%
social/app/grpc/grpc.go:16:     NewServer       100.0%
social/app/grpc/grpc.go:20:     Run             0.0%
social/app/grpc/post.go:11:     CreatePost      66.7%
social/app/grpc/post.go:26:     GetPosts        0.0%
social/app/grpc/user.go:10:     CreateUser      71.4%
social/app/grpc/user.go:22:     GetUser         0.0%
total:                          (statements)    31.0%

```

Our coverage is pretty average however,  we have greater than 50% coverage for all DB write methods.



## Part 4: Create Frontend

Armed with our fully implemented social GRPC server, we can now implement the front end API to use the grpc api and manage user input.

To Create our Front end app we are going to add our third module: `http`

`http` will include our vue app as well as a go http server we can use to run the vue app. 

I wont go to deep into the specific code implementation but it involves the following steps

1. installing vue, grpc-web node dependencies
2. creating an `http/frontend` folder and starting the vue app in there with `vue create`
3. updating the protoc command to complie javascrip/proto files to the destination file:
```
protoc --proto_path=protos --js_out=import_style=commonjs,binary:../http/static/js --grpc-web_out=import_style=commonjs,mode=grpcwebtext:../http/static/js protos/social.proto
```
4. adding code to the generated App.vue file to login a user using the GRPC api:
```vue
<template>

  <div id="app" class="container" >
    <div class="row">
      <div class="col-md-6 offset-md-3 py-5">
        <h1>Welcome to the social network</h1>
      </div>
    </div>
    <div class="row" v-if="loggedIn">
      <div class="col-md-6 offset-md-3 py-5">
        <h2>You are logged in, welcome {{this.username}}</h2>
      </div>
    </div>
    <div class="row" v-else>
      <div class="col-md-6 offset-md-3 py-5">
        <h1>Login to Social Network</h1>
        <p>If you dont have a login, enter a username and password and one will be created</p>
        <form v-on:submit.prevent="createUser">
          <div class="form-group">
            <input v-model="username" type="text" id="username-input" placeholder="Enter a username" class="form-control">
            <input v-model="password" type="password" id="password-input" placeholder="Enter a password" class="form-control">
          </div>
          <div class="form-group">
            <button class="btn btn-primary">Create!</button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script>
import { BootstrapVue, IconsPlugin } from 'bootstrap-vue'
import {SocialPromiseClient} from './static/js/social_grpc_web_pb'
import {CreateUserReq, User} from './static/js/social_pb'

import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'
export default {
  name: 'App',

  created: function() {
    this.grpcClient = new SocialPromiseClient("http://localhost:8083", null, null);
  },

  data() { return {
    username: '',
    password: '',
    userId: 0,
    loggedIn: false,
  } },

  methods: {
    async createUser() {
      console.log("Creating user: " + this.username)

      try {
        const user = new CreateUserReq()
        user.setUserName(this.username)
        user.setPassword(this.password)
        const res = await this.grpcClient.createUser(user, {})
        console.log("Successfully received GRPC response, object returned is: " + JSON.stringify(res))
        this.loggedIn = true
        this.$forceUpdate()

      } catch (err) {
        console.error(err.message)
        console.log("err in grpc response: ", err.message);
        throw err
      }
    }
  }
}
</script>
```

Note this vue app is assuming a GRPC client is available at localhost:8083

Next, we need to slightly modify the CreateUserApi to support login by updating the CreateUser rpc to perform a get and match password before returning:

```go

func (s *Server) CreateUser(ctx context.Context, req *pb.CreateUserReq) (*pb.User, error) {
	if user, err := s.r.GetUserByUserName(req.GetUserName()); err == nil {
		if user.Password == req.GetPassword() {
			return serializers.User(user), nil
		} else {
			return nil, fmt.Errorf("invalid user")
		}
	}
	
}

```


## Part 5: building / deploying with docker-compose

At this point we have all the pieces to build and deploy the social grpc api and http server serving the vuejs files. 
If deployed correctly a user should be able to login via the CreateUserAPI frontend/backend integration

To enable the GRPC connection from our VUEjs app and GRPC api we need to use an envoy proxy between the 2 http server and go routine runinng the GRPC api.
 
We can plugin an envoy proxy to do this.  In a production system, this proxy would usually be at a layer upstream from the machine receiving the proxy requests.  However, for local development we can simmulate this with an envoy docker deployment.

See the `envoy/` package which consists of the envoy config YAML and Dockerfile.  These settings are default for grpc/http proxy for envoy.  

The one thing to note is the listener port address of 8083:
```
listeners:
- name: listener_0
address:
socket_address: { address: 0.0.0.0, port_value: 8083 }

```

This is the address the HTTP server will attempt to connect to reach the GRPC server. 

The echo service points to the port that is running the actual GRPC server. 

When a user performs an action on our app that triggers a grpc call the http server will call port 8083 which will proxy the GRPC server at 50053. 
```
 hosts: [{ socket_address: { address: host.docker.internal, port_value: 50053 }}]
```

To deploy we need 3 things: MYSQL database container, GO Binary, envoy proxy.

The go binary will run 2 programs at port 50053 and 3000, the grpc and http server respectively. 

To facilitate this,  we can create a main.go single binary to deploy at `http/server/main.go with the following:

```go
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"social/app/grpc"
	"social/repo"
	"time"
)

const (
	tenmplatesDir = "./frontend/src/static"
)

func main() {
	// Serve static files from the frontend/dist directory.
	var err error
	go func() {
		// TODO: replace with env
		host := os.Getenv("DB_HOST")
		dbname := os.Getenv("DB_NAME")
		pw := os.Getenv("DB_PASS")
		port := os.Getenv("DB_PORT")
		user := os.Getenv("DB_USER")
		r, rErr := repo.NewRepo(&repo.Config{
			DbPass: pw,
			DbUser: user,
			DbName: dbname,
			DbHost: host,
			DbPort: port,
		})
		if rErr != nil {
			log.Fatal(rErr)
			return
		}
		err = grpc.Run(":50053", grpc.NewServer(r))
	}()
	time.Sleep(time.Second * 2)
	if err != nil {
		log.Fatal(err)
		return
	}
	fs := http.FileServer(http.Dir("./dist"))
	http.Handle("/", fs)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(tenmplatesDir))))

	// Start the server.
	fmt.Println("Server listening on port 3000")
	log.Panic(
		http.ListenAndServe(":3000", nil),
	)
}

```

This also gives us the opportunity to swap out hardcoded DB params to environment variables we can manage outside the application. 

Note we are also serving /static/ directory to be used by the vue app

This line serves the compiled vue app /dist directory:
```go
fs := http.FileServer(http.Dir("./dist"))
http.Handle("/", fs)
```

We can also see that the GRPC server is run in the go routine. This is something that can be refactored to be run as a separate deployment.


To build the binary we can do the following dockerfile:
```dockerfile
# Build and bundle the Vue.js frontend SPA
#
FROM node:14-alpine AS vue-build
WORKDIR /build

COPY http/frontend/package*.json ./
RUN npm install

COPY http/frontend .


RUN npm run build

#
# Build the Go server backend
#
FROM golang:1.16-alpine as go-build

WORKDIR /build/src/

RUN apk update && apk add git gcc musl-dev

COPY http/ ./http
COPY db/ ./db
COPY grpc/ ./grpc

ENV GO111MODULE=on
WORKDIR /build/src/http/server
# Disabling cgo results in a fully static binary that can run without C libs
RUN CGO_ENABLED=0 GOOS=linux go build -o main.go

#
# Assemble the server binary and Vue bundle into a single app
#
FROM alpine
WORKDIR /app

COPY --from=vue-build /build/dist ./dist
COPY --from=go-build /build/src/http/server .
COPY --from=go-build /build/src/db/schema ./db/schema



ENV PORT 8080
EXPOSE 8080
CMD ["/app/main.go"]
```

This Dockerfile builds go and vue code into one http server binary.

One thing to note is we have to build the directory tree in the docker file to match whatever is expected in the go.mod file replace statements,

6.Next we create a docker-compose.yaml file that can run 3 dockefiles:  
 - mysql/Dockerfile [localhost, port:3306]
 - http/Dockerfile  [localhost, ports:3000, 50053]
 - envoy/Dockerfile[localhost, port:8083]

```yaml
version: "3"
services:
  # This is the MYSQL server,  this will/should only be using for development.
  # In production we will set an Env variable indicating it is prod and we will
  db:
    restart: always
    build:
      context: ./mysql
    environment:
      MYSQL_ROOT_PASSWORD: tauhisgasgknga
    container_name: db
    volumes:
      - dbdata:/var/lib/mysql
    ports:
      - "3306:3306"
    tty: true
    healthcheck:
      test: ["CMD-SHELL", 'mysqladmin -u root -ptauhisgasgknga ping']
      interval: 10s
      timeout: 2s
      retries: 10
    security_opt:
      - seccomp:unconfined
  backend:
    links:
      - db
    depends_on:
      db:
        condition: service_healthy
    build:
      context: .
      dockerfile: "http/Dockerfile"
    container_name: backend
    ports:
      - "50053:50053" # GRPC server // clients can connect to this via GRPC (expirimental)
      - "3000:3000" # HTTP server // clients can connect to this via HTTP
    tty: true
    environment:
      - "DB_PORT=3306"
      - "DB_HOST=db"
      - "DB_NAME=test"
      - "DB_PASS=tauhisgasgknga"
      - "DB_USER=root"
      - "MIGRATIONS_PATH=./db/schema"
  envoy:
    links:
      - backend
    build:
      context: ./envoy
    container_name: envoy
    ports:
      - "8083:8083"
    tty: true

volumes:
  dbdata:
```


```bash
docker-compose up --build
```

Should run go server binary, mysql, and envoy container:

You should see the login screen at localhost:3000

![image](https://user-images.githubusercontent.com/2126188/120285542-33707380-c272-11eb-8813-5791b5810265.png)

And after logging in you should see:
![image](https://user-images.githubusercontent.com/2126188/120285582-3e2b0880-c272-11eb-998c-d792a530f2f1.png)

Backend logs:

`backend    | 2021/06/01 07:26:37 Successfully created user
`

## Part 6: Adding Create/Get Post user functionality

We want the user to be able to see posts they have created as well as create new posts. 

We can do this by adding code to our Vue app to call the GRPC posts api methods.

```vue
<template>

  <div id="app" class="container" >
    <div class="row">
      <div class="col-md-6 offset-md-3 py-5">
        <h1>Welcome to the social network</h1>
      </div>
    </div>
    <div class="row" v-if="loggedIn">
      <div class="col-md-6 offset-md-3 py-5">
        <h2>You are logged in, welcome {{this.username}}</h2>
      </div>
    </div>
    <div class="row" v-else>
      <div class="col-md-6 offset-md-3 py-5">
        <h1>Login to Social Network</h1>
        <p>If you dont have a login, enter a username and password and one will be created</p>
        <form v-on:submit.prevent="login">
          <div class="form-group">
            <input v-model="username" type="text" id="username-input" placeholder="Enter a username" class="form-control">
            <input v-model="password" type="password" id="password-input" placeholder="Enter a password" class="form-control">
          </div>
          <div class="form-group">
            <button class="btn btn-primary">Login!</button>
          </div>
        </form>
      </div>
    </div>
    <div v-if="loggedIn">
      <div class="row">
        <div class="col-md-6 offset-md-3 py-5">
          <h3>{{this.username}} Create a Post: </h3>
          <p>Enter some text to create a post</p>
          <form v-on:submit.prevent="createPost">
            <div class="form-group">
              <input v-model="newPost" type="text" id="newPost-input" placeholder="Your Post" class="form-control">
            </div>
            <div class="form-group">
              <button class="btn btn-primary">Post!</button>
            </div>
          </form>
        </div>
      </div>
      <div class="row" v-if="posts">
        <div class="col-md-6 offset-md-3 py-5">
          <h2>{{this.username}}'s Posts: </h2>
          <div class="row" v-for="post in posts" :key="post.id">
            {{post.id}}: {{post.content}}
          </div>
        </div>
      </div>
      <div class="row" v-else>
        <div class="col-md-6 offset-md-3 py-5">
          <h2> {{this.username}} has no posts in the social network </h2>
        </div>
      </div>

    </div>
  </div>
</template>

<script>
import { BootstrapVue, IconsPlugin } from 'bootstrap-vue'
import {SocialPromiseClient} from './static/js/social_grpc_web_pb'
import {CreateUserReq, GetPostsReq, GETPOSTSIDTYPE_USER, CreatePostReq} from './static/js/social_pb'

import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'
export default {
  name: 'App',

  created: function() {
    this.grpcClient = new SocialPromiseClient("http://localhost:8083", null, null)
  },

  data() { return {
    username: '',
    password: '',
    newPost: '',
    userId: 0,
    loggedIn: false,
    posts : [],
  } },

  methods: {
     async login() {
       try {
         const user = new CreateUserReq()
         user.setUserName(this.username)
         user.setPassword(this.password)
         const res = await this.grpcClient.createUser(user, {})
         console.log("Login Successful")
         this.loggedIn = true
         this.userId = res.toObject().id
         await this.getUserPosts()
         this.$forceUpdate();

       } catch (err) {
         console.error(err.message)
         console.log("err in grpc response: ", err.message)
         this.username = ""
         this.password = ""
         this.userId = 0
         throw err
       }
     },
    async createPost() {
      if (this.userId !== 0 && this.loggedIn && this.newPost !== '') {
        try {
          const post = new CreatePostReq()
          post.setContent("{\"body\":\""+  this.newPost +"\"}")
          post.setUserId(this.userId)
          const res = await this.grpcClient.createPost(post, {})
          console.log("createPost Successful")
          await this.getUserPosts();

        } catch (err) {
          console.error(err.message)
          console.log("err in grpc response: ", err.message)
          throw err
        }
      }
    },
    async getUserPosts() {
      if (this.userId !== 0 && this.loggedIn) {
        try {
          const postReq = new GetPostsReq()
          postReq.setGetBy(GETPOSTSIDTYPE_USER)
          let idList = []
          idList.push (this.userId)
          postReq.setIdsList(idList)
          const res = await this.grpcClient.getPosts(postReq, {})
          console.log("getUserPosts successful")
          this.posts = res.toObject().itemsList

        } catch (err) {
          console.error(err.message)
          console.log("err in grpc response: ", err.message)
          throw err
        }
      }
    }

  }
}
</script>

``` 

In addition we have a create post form in our VUE app that is connected to a createPost vue method to call the GRPC api. 

We trigger a getPosts method call after login or a new post is submitted.

After re-running `docker-compose up --build` we can see the app running with post functionality working:

Once logged in and after entering a post, the page should show the post dynamically as updated:

![image](https://user-images.githubusercontent.com/2126188/120374709-bcb59380-c2ce-11eb-91c0-fcc97a6a91b0.png)


## Part 7: Adding User Feed. 

We will use a "fan-out" approach to creating the user feed.  What this means is that any time an item is posted, we will "fan-out" the item to every other user. 

For example, we have users A, B and C

A makes a post:
    - create Post item in DB POST ID 100 (owned by USER A)
user A post gets Fanned out to B and C
    - Create a FeedItem in (OWNED by user B, points to POST 100)
    - Create a FeedItem in (OWNED by user C, points to POST 100)

When loading user B (and C's feed), we can query for all FeedItems owned by the respective user.  
Next we can query the posts by the feed.post_id to get the post content and return in the user's feed. 

To support this anytime a user makes a post through the GRPC create post endpoint, we need to trigger a process to populate all other user's feed table rows with the post and user ID.

We eventually will want to refactor this to create some notion of following so users only see posts they care about but for now we will fan out all posts to all users.
The implementation for this is a function to populate the feed based on a postId and the post's owner ID:

```go
func (s *Server) populateFeed(postId, postOwnerId int64) error {
	users, err := s.r.GetUsersOtherThanId(postOwnerId)
	if err != nil {
		return err
	}
	for _, user := range users.Items {
		toCreate := &pb.CreateFeedItemReq{
			OwnerId: user.Id,
			PostId:  postId,
		}
		if _, err = s.CreateFeedItem(context.TODO(), toCreate); err != nil {
			log.Println("err inserting feed item: ", err)

		}
	}
	if err != nil {
		return err
	}
	return nil

}

```

For now we can call this in-line as a go routine in CreatePost This will ensure the feed will get populated but not block the user path:

```go
func (s *Server) CreatePost(ctx context.Context, req *pb.CreatePostReq) (*pb.Post, error) {
	insertedId, err := s.r.InsertPost(serializers.CreatePostReq(req))
	if err != nil {
		return nil, err
	}
	posts, err := s.r.GetPostsIds([]int64{insertedId}, repo.FieldNameSocialPostId)
	if err != nil {
		return nil, err
	}
	if posts.Items == nil || len(posts.Items) != 1 {
		return nil, fmt.Errorf("err unexp items")
	}
	created := posts.Items[0]
	go s.populateFeed(created.Id, created.UserId)
	return serializers.Post(posts.Items[0]), nil
}
```

Note we also add a repo method to search for all users that dont match the content owner to populate the feed for that item

Note, it is not ideal that this just spawns a go routine that does N * number of users inserts on the DB.  If 1000 users were posting at the same time, that is 1000*1000 go routines spawning doing inserts to the DB at the same time.
A way to refactor this is either to instead push a msg with the post_id and user_id to a pubsub and have 1 worker process doing the fanout logic.  This way we can throttle inserts through this one point.
Another option is to use a go routine worker pool to ensure only X amount of go routines are executing inserts at one time.

We can come back to this and refactor it later.

Now that we have this logic in place we can re-spin up our server with `docker-compose up --build`

If we create 2 browser tabs and login with 2 different users and make a post with userA, that post should show up on user Bs feed:

![image](https://user-images.githubusercontent.com/2126188/120383731-1a031200-c2da-11eb-8a4c-a8f2c445490c.png)

We can also see that in the backend the feed items were inserted: 
```
backend    | 2021/06/01 20:38:19 successfully inserted Feed item
backend    | 2021/06/01 20:38:19 successfully inserted Feed item
backend    | 2021/06/01 20:38:19 successfully inserted Feed item
backend    | 2021/06/01 20:38:19 successfully inserted Feed item
...
```
You probably noticed it doesn't look right, we see "TODO" instead of "USER A says hi" which what we expect.  

This is because the returned object from Feed only has the PostId not the actual post.content. We can fix this by updating the feed RPC to query and returned post_ids against our GetPost api and update the response object on Feed to include Feed.Post.Content

To fix this we can add a query in the GetFeed call to make another call to get the corresponding posts, and update the serializer for Feed to display the content:

Updated GetFeed:

```go
func (s *Server) GetFeed(ctx context.Context, req *pb.GetFeedReq) (*pb.Feed, error) {
	feed, err := s.r.GetFeedByOwnerId(req.GetOwnerId())
	if err != nil {
		return nil, err
	}
	postIds := make([]int64, 0, len(feed.Items))

	for _, f := range feed.Items {
		postIds = append(postIds, f.PostId)
	}
	posts, err := s.GetPosts(ctx, &pb.GetPostsReq{
		Ids:   postIds,
		GetBy: pb.GetPostsReq_GetPostsIdType_post,
	})
	if err != nil {
		return nil, err
	}
	return serializers.Feed(feed, posts), nil
}

```

update serializer:

```go
func Feed(items *repo.Feed, posts *pb.Posts) *pb.Feed {
	if len(items.Items) == 0 || len(posts.GetItems()) == 0 {
		return nil
	}
	postMap := map[int64]string{}
	for _, p := range posts.Items {
		if p.GetContent() != "{}" {
			postMap[p.Id] = p.GetContent()
		}
	}
	out := make([]*pb.FeedItem, 0, len(items.Items))
	for _, item := range items.Items {
		fi := FeedItem(item)
		if v, ok := postMap[item.PostId]; ok {
			fi.PostContent = v
			out = append(out, fi)
		}
	}
	return &pb.Feed{Items: out}
}
```

And finally update to render the content in vue:

```vue 
 <div class="col-md-4 offset-md-1 py-5" v-if="hasFeed()">
          <h2>{{this.username}}'s Feed: </h2>
          <div class="row" v-for="feedItem in feed" :key="feedItem.id">
            {{feedItem.id}}: {{feedItem.postContent}}}
          </div>
        </div>
        <div class="col-md-4 offset-md-1 py-5" v-else>
          <h2> {{this.username}} has no posts in their Feed </h2>
        </div>
      </div>

  ```

Now when loading userBs profile, we can see the actual feed content:
