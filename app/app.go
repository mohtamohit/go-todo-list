package app

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/mohtamohit/go-todo-list/db"
	"github.com/mohtamohit/go-todo-list/handler"
	"github.com/mohtamohit/go-todo-list/todo"
)

func PrintInstructions() {
	fmt.Println("To create a new task (eg.) : \ncreate\n<task_name>")
	fmt.Println("To read an existing task (eg.) : \nread\n<task_id>")
	fmt.Println("To show all tasks (eg.) : \nshow_all")
	fmt.Println("To update an existing task (eg.) : \nupdate\n<task_id>\n<new_task_name>")
	fmt.Println("To delete an existing task (eg.) : \ndelete\n<task_id>")
	fmt.Println("To mark a task as done (eg.) : \nmark_done\n<task_id>")
}

func StartServer() {
	log.Println("Server started on: http://localhost:80")
	http.HandleFunc("/", handler.Index)
	http.HandleFunc("/read", handler.Read)
	http.HandleFunc("/create", handler.Create)
	http.HandleFunc("/update", handler.Update)
	http.HandleFunc("/delete", handler.Delete)
	http.HandleFunc("/markdone", handler.MarkDone)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}

func StartCLI() {
	PrintInstructions()
	dbIns := db.InitDB()
	defer dbIns.Close()

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
			ts, err := todo.Read(dbIns, task_id)

			if err != nil {
				fmt.Println("Couldn't read this task. Check and try again.")
			} else {
				fmt.Println(ts.Task_id, " ", ts.Task, " ", ts.Created_at, " ", ts.Status)
			}

		case "show_all":
			fmt.Println("In show all")
			ts, err := todo.ShowAll(dbIns)
			if err != nil {
				fmt.Println("Couldn't show tasks. Check and try again.")
			} else {
				for _, t := range ts {
					fmt.Println(t.Task_id, t.Task, t.Created_at, t.Status)
				}
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

		case "mark_done":
			var task_id int
			fmt.Scanln(&task_id)
			err := todo.MarkDone(dbIns, task_id)
			if err != nil {
				fmt.Println("Couldn't mark the task as done. Check and try again.")
			} else {
				fmt.Println("Task with task id:", task_id, "marked as done.")
			}

		default:
			fmt.Println("Invalid option. Refer to the instructions given above.")
		}
	}
}
