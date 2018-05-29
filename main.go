package main

import (
	"fmt"
	"os"
	"practice/go-todo-list/app"
	"practice/go-todo-list/config"
	"practice/go-todo-list/migration"
	"practice/go-todo-list/todo"

	"github.com/urfave/cli"
)

func main() {
	config.Load()
	app.InitApp()

	app := cli.NewApp()
	app.Name = "todo-app"
	app.Version = "0.0.1"
	fmt.Println("####DEBUG####")

	app.Commands = []cli.Command{
		{
			Name:        "migrate",
			Description: "Run Database migration",
			Action: func(c *cli.Context) error {
				migration.Init()
				return migration.Up()
			},
		},
		{
			Name:        "rollback",
			Description: "Rollback latest database migration",
			Action: func(c *cli.Context) error {
				migration.Init()
				return migration.Down()
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println("###PANIC YAHA##")
		panic(err)
	}

	startApp()
}

func startApp() {
	var choice string

	for {
		fmt.Scanln(&choice)
		switch choice {
		case "create":
			var task string
			fmt.Scanln(&task)
			task_id, _ := todo.Create(task)
			fmt.Println(task_id)

		case "read":
			var task_id int
			fmt.Scanln(&task_id)
			task, _ := todo.Read(task_id)
			fmt.Println(task)

		case "show_all":
			fmt.Println("In show all")
			err := todo.ShowAll()
			if err != nil {
				fmt.Println(err)
			}

		case "update":
			var task_id int
			var task string
			fmt.Scanln(&task_id)
			fmt.Scanln(&task)
			err := todo.Update(task_id, task)
			if err != nil {
				fmt.Println(err)
			}

		case "delete":
			var task_id int
			fmt.Scanln(&task_id)
			err := todo.Delete(task_id)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
