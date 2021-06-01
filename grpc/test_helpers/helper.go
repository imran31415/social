package test_helpers

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"social/repo"
	"strings"
)

const (
	DbUser = "social_test"
	DbPass = "social_test"
)

func SetupDbForTest(dbName string, sqlFile string, migrationsPath string) (*repo.Repo, error) {
	os.Setenv("MIGRATIONS_PATH", "../db/schema")
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
		config             *repo.Config
		migrationsPathEnv  string
		insertsToApplyPath string
	}

	tEnv := &testEnv{
		config: &repo.Config{
			DbPass: DbUser,
			DbUser: DbPass,
			DbName: dbName,
			DbHost: "localhost",
			DbPort: "3306",
		},
		migrationsPathEnv: migrationsPath,
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
