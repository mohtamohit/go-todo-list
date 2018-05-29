package migration

import (
	"fmt"
	"os"
	"practice/go-todo-list/config"

	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/golang-migrate/migrate/source/github"
	_ "github.com/lib/pq"
)

var migrationPath = "file://" + os.Getenv("GOPATH") + "/src/practice/go-todo-list/migration/queries/"

var runner *migrate.Migrate

func Init() {
	connURL := config.Db().ConnString()
	var err error
	fmt.Println(migrationPath, connURL)
	runner, err = migrate.New(migrationPath, connURL)
	if err != nil {
		panic(err)
	}
	runner.Steps(1)
}

func Up() error {
	// runner. up makes it as up as possible
	if err := runner.Up(); err != nil {
		return fmt.Errorf("Error while migration up: %v", err)
	}

	fmt.Println("Migration successful")
	return nil
}

func Down() error {
	// runner.steps -1 brings it down just by one step
	// if err := runner.Steps(-1); err != nil {
	// 	fmt.Println("Down fata")
	// 	return err
	// }
	// fmt.Println("Migration Successfull")
	// return nil

	// runner.down brings it as down as possible, in this case will drop the table
	if err := runner.Down(); err != nil {
		return fmt.Errorf("Error while migration down: %v", err)
	}

	fmt.Println("Migration successful")
	return nil
}
