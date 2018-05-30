package main

import (
	"os"
	"practice/go-todo-list/app"
	"practice/go-todo-list/config"
	"practice/go-todo-list/db"
	"practice/go-todo-list/migration"

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
	}

	if err := cliApp.Run(os.Args); err != nil {
		panic(err)
	}

	app.PrintInstructions()
	app.StartApp(dbIns)
}
