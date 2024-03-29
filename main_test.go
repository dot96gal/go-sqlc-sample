package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

var db *sql.DB

func TestMain(m *testing.M) {
	mysqlDatabase := os.Getenv("TEST_MYSQL_DATABASE")
	mysqlRootPass := os.Getenv("TEST_MYSQL_ROOT_PASSWORD")
	mysqlUser := os.Getenv("TEST_MYSQL_USER")
	mysqlPass := os.Getenv("TEST_MYSQL_PASSWORD")
	mysqlHost := os.Getenv("TEST_MYSQL_HOST")
	mysqlPort := os.Getenv("TEST_MYSQL_TCP_PORT")

	pwd, _ := os.Getwd()
	schemaFiles := fmt.Sprintf("file://%s/db/migrations", pwd)

	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not construct pool: %s", err)
	}

	pool.MaxWait = 30 * time.Second

	err = pool.Client.Ping()
	if err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}

	runOptions := &dockertest.RunOptions{
		Repository: "mysql",
		Tag:        "8.3.0",
		Env: []string{
			fmt.Sprintf("MYSQL_DATABASE=%s", mysqlDatabase),
			fmt.Sprintf("MYSQL_ROOT_PASSWORD=%s", mysqlRootPass),
			fmt.Sprintf("MYSQL_USER=%s", mysqlUser),
			fmt.Sprintf("MYSQL_PASSWORD=%s", mysqlPass),
			fmt.Sprintf("MYSQL_HOST=%s", mysqlHost),
			fmt.Sprintf("MYSQL_TCP_PORT=%s", mysqlPort),
		},
	}

	resource, err := pool.RunWithOptions(
		runOptions,
		func(config *docker.HostConfig) {
			config.AutoRemove = true
			config.RestartPolicy = docker.RestartPolicy{
				Name: "no",
			}
		},
	)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	port := resource.GetPort(fmt.Sprintf("%s/tcp", mysqlPort))
	dataSource := fmt.Sprintf("%s:%s@(%s:%s)/%s", mysqlUser, mysqlPass, mysqlHost, port, mysqlDatabase)

	if err := pool.Retry(func() error {
		db, err = sql.Open("mysql", dataSource)
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to database: %s", err)
	}

	// database migration
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatalf("Could not instantiate driver: %s", err)
	}
	mig, err := migrate.NewWithDatabaseInstance(schemaFiles, "mysql", driver)
	if err != nil {
		log.Fatalf("Could not instantiate migrate: %s", err)
	}
	err = mig.Up()
	if err != nil {
		log.Fatalf("Could not migrate database: %s", err)
	}

	code := m.Run()

	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}
