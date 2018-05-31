package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mohtamohit/go-todo-list/handler"
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
	log.Println("Server started on: http://localhost:8080")
	http.HandleFunc("/", handler.Index)
	http.HandleFunc("/read", handler.Read)
	http.HandleFunc("/create", handler.Create)
	http.HandleFunc("/update", handler.Update)
	http.HandleFunc("/delete", handler.Delete)
	http.HandleFunc("/markdone", handler.MarkDone)
	http.ListenAndServe(":8080", nil)
}
