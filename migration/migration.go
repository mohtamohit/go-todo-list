package migration

import (
	"fmt"
	"os"

	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/golang-migrate/migrate/source/github"
	_ "github.com/lib/pq"
	"github.com/mohtamohit/go-todo-list/config"
)

var migrationPath = "file://" + os.Getenv("GOPATH") + "/src/github.com/mohtamohit/go-todo-list/migration/queries/"

var runner *migrate.Migrate

func Init() {
	connURL := config.Db().ConnString()
	var err error
	runner, err = migrate.New(migrationPath, connURL)
	if err != nil {
		panic(err)
	}
	runner.Steps(1)
}

func Up() error {
	if err := runner.Up(); err != nil {
		return fmt.Errorf("Error while migration up: %v", err)
	}

	fmt.Println("Migration successful")
	return nil
}

func Down() error {
	if err := runner.Down(); err != nil {
		return fmt.Errorf("Error while migration down: %v", err)
	}

	fmt.Println("Migration successful")
	return nil
}
