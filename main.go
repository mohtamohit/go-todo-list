package main

import (
	"os"

	"github.com/mohtamohit/go-todo-list/app"
	"github.com/mohtamohit/go-todo-list/config"
	"github.com/mohtamohit/go-todo-list/migration"
	"github.com/urfave/cli"
)

func main() {
	config.Load()
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
			Name:        "web",
			Description: "Start the web app",
			Action: func(c *cli.Context) {
				app.StartServer()
			},
		},
		{
			Name:        "cli",
			Description: "Launch the cli app",
			Action: func(c *cli.Context) {
				app.StartCLI()
			},
		},
	}

	if err := cliApp.Run(os.Args); err != nil {
		panic(err)
	}

	app.PrintInstructions()
}
