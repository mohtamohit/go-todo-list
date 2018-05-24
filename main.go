package main

import (
	"practice/go-todo-list/app"
	"practice/go-todo-list/config"
)

func main() {
	config.Load()

	app.Init()
	defer app.Close()

	startApp()
}

func startApp() {

}
