package main

import (
	"fmt"
	"practice/go-todo-list/config"
	"practice/go-todo-list/todo"
)

func main() {
	config.Load()

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
