package app

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"practice/go-todo-list/todo"
)

func PrintInstructions() {
	fmt.Println("To create a new task (eg.) : \ncreate\n<task_name>")
	fmt.Println("To read an existing task (eg.) : \nread\n<task_id>")
	fmt.Println("To show all tasks (eg.) : \nshow_all")
	fmt.Println("To update an existing task (eg.) : \nupdate\n<task_id>\n<new_task_name>")
	fmt.Println("To delete an existing task (eg.) : \ndelete\n<task_id>")
}

func StartApp(dbIns *sql.DB) {
	bio := bufio.NewReader(os.Stdin)
	var choice string
	for {
		fmt.Scanln(&choice)
		switch choice {
		case "create":
			var task string
			task_byte, _, _ := bio.ReadLine()
			task = string(task_byte)
			task_id, err := todo.Create(dbIns, task)
			if err != nil {
				fmt.Println("Couldn't create this task. Check and try again.")
			} else {
				fmt.Println("Task with task id:", task_id, "created.")
			}

		case "read":
			var task_id int
			fmt.Scanln(&task_id)
			task, err := todo.Read(dbIns, task_id)
			if err != nil {
				fmt.Println("Couldn't read this task. Check and try again.")
			} else {
				fmt.Println(task_id, " ", task)
			}

		case "show_all":
			fmt.Println("In show all")
			err := todo.ShowAll(dbIns)
			if err != nil {
				fmt.Println("Couldn't show tasks. Check and try again.")
			}

		case "update":
			var task_id int
			var task string
			fmt.Scanln(&task_id)
			task_byte, _, _ := bio.ReadLine()
			task = string(task_byte)
			err := todo.Update(dbIns, task_id, task)
			if err != nil {
				fmt.Println("Couldn't perform this update. Check and try again.")
			} else {
				fmt.Println("Task with task id:", task_id, "updated.")
			}

		case "delete":
			var task_id int
			fmt.Scanln(&task_id)
			err := todo.Delete(dbIns, task_id)
			if err != nil {
				fmt.Println("Couldn't perform this delete. Check and try again.")
			} else {
				fmt.Println("Task with task id:", task_id, "deleted.")
			}

		default:
			fmt.Println("Invalid option. Refer to the instructions given above.")
		}
	}
}
