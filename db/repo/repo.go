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

type Iface interface {
	InsertUser(p *User) (int64, error)
	GetUserById(id int64) (*User, error)
	GetUserByUserName(username string) (*User, error)
	GetAllUsers() (*Users, error)

	InsertPost(p *Post) (int64, error)
	GetPostsIds(ids []int64, fieldName SocialPostFieldName) (*Posts, error)

	GetCommentsByPostId(id int64) (*Comments, error)
	InsertComment(p *Comment) (int64, error)
	GetCommentById(id int64) (*Comment, error)

	InsertFeedItem(p *FeedItem) (int64, error)
	GetFeedByOwnerId(id int64) (*Feed, error)
	GetFeedItemById(id int64) (*FeedItem, error)
}
type Repo struct {
	Db     *sqlx.DB
	Config *Config
}

type Config struct {
	DbPass string
	DbUser string
	DbName string
	DbHost string
	DbPort string
}

func (c *Config) MigrationsPath() string {
	return filepath.FromSlash(os.Getenv("MIGRATIONS_PATH"))
}

// NewRepo initializes the struct as well as connects to the database and performs the initial migrations.
func NewRepo(config *Config) (*Repo, error) {
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
		path, err := os.Getwd()
		if err != nil {
			log.Println(err)
		}
		fmt.Println(path) // for example /home/user
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
