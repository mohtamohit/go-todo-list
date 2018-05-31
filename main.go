package main

import (
	"os"

	"github.com/mohtamohit/go-todo-list/app"
	"github.com/mohtamohit/go-todo-list/config"
	"github.com/mohtamohit/go-todo-list/db"
	"github.com/mohtamohit/go-todo-list/migration"
	"github.com/urfave/cli"
)

func main() {
	config.Load()

	dbIns := db.InitDB()
	defer dbIns.Close()

	cliApp := cli.NewApp()
	cliApp.Name = config.AppName()
	cliApp.Version = config.AppVersion()

	cliApp.Commands = []cli.Command{
		{
			Name:        "migrate",
			Description: "Run Database migration",
			Action: func(c *cli.Context) error {
				migration.Init()
				defer os.Exit(0)
				return migration.Up()
			},
		},
		{
			Name:        "rollback",
			Description: "Rollback latest database migration",
			Action: func(c *cli.Context) error {
				migration.Init()
				defer os.Exit(0)
				return migration.Down()
			},
		},
		{
			Name:        "start_server",
			Description: "Rollback latest database migration",
			Action: func(c *cli.Context) {
				app.StartServer()
			},
		},
	}

	if err := cliApp.Run(os.Args); err != nil {
		panic(err)
	}

	app.PrintInstructions()
}
