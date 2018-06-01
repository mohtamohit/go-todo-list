package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/mohtamohit/go-todo-list/db"
	"github.com/mohtamohit/go-todo-list/todo"
)

var tmpl = template.Must(template.ParseGlob("templates/*"))

func Index(w http.ResponseWriter, r *http.Request) {
	dbIns := db.InitDB()
	defer dbIns.Close()

	res, err := todo.ShowAll(dbIns)
	if err != nil {
		fmt.Println("Error while showing all.")
	}
	tmpl.ExecuteTemplate(w, "Index", res)
}

func Read(w http.ResponseWriter, r *http.Request) {
	dbIns := db.InitDB()
	defer dbIns.Close()

	task_id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	res, err := todo.Read(dbIns, task_id)
	if err != nil {
		fmt.Println("Error while showing all.")
	}
	tmpl.ExecuteTemplate(w, "Read", res)
}

func Create(w http.ResponseWriter, r *http.Request) {
	dbIns := db.InitDB()
	defer dbIns.Close()

	if r.Method == "POST" {
		task := r.FormValue("Task")
		task_id, err := todo.Create(dbIns, task)
		if err != nil {
			fmt.Println("Error while showing all.")
		} else {
			fmt.Println("Task with task id:", task_id, "created.")
		}
	}
	http.Redirect(w, r, "/", 301)
}

func Update(w http.ResponseWriter, r *http.Request) {
	dbIns := db.InitDB()
	defer dbIns.Close()

	if r.Method == "POST" {
		task_id, _ := strconv.Atoi(r.FormValue("Task_id"))
		task := r.FormValue("Task")
		err := todo.Update(dbIns, task_id, task)
		if err != nil {
			fmt.Println("Error while showing all.")
		}
	}
	http.Redirect(w, r, "/", 301)
}

func MarkDone(w http.ResponseWriter, r *http.Request) {
	dbIns := db.InitDB()
	defer dbIns.Close()

	task_id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	err := todo.MarkDone(dbIns, task_id)
	if err != nil {
		fmt.Println("Error while showing all.")
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	dbIns := db.InitDB()
	defer dbIns.Close()

	task_id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	err := todo.Delete(dbIns, task_id)
	if err != nil {
		fmt.Println("Error while showing all.")
	}
	http.Redirect(w, r, "/", 301)
}
