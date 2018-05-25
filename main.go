package main

import (
	"html/template"
	"log"
	"net/http"
	"practice/go-todo-list/todo"
)

var tmpl = template.Must(template.ParseGlob("form/*"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	res := todo.ShowAll()
	tmpl.ExecuteTemplate(w, "Index", res)

}

func createHandler(w http.ResponseWriter, r *http.Request) {
	res := todo.ShowAll()
	tmpl.ExecuteTemplate(w, "Index", res)

}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	res := todo.ShowAll()
	tmpl.ExecuteTemplate(w, "Index", res)

}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	res := todo.ShowAll()
	tmpl.ExecuteTemplate(w, "Index", res)
}

func main() {

	log.Println("Server started on: http://localhost:8080")
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/create", createHandler)
	http.HandleFunc("/update", updateHandler)
	http.HandleFunc("/delete", deleteHandler)
	http.ListenAndServe(":8080", nil)

	// config.Load()

	// app.Init()
	// defer app.Close()

	// startApp()
}

func startApp() {

}
