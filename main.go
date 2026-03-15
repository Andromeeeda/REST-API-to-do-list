package main

import (
	"RestApi/http"
	"RestApi/todo"
	"fmt"
)

func main() {
	todoList := todo.NewList()
	HttpHandlers := http.NewHttpHandler(todoList)
	HttpServer := http.NewHttpServer(HttpHandlers)

	if err := HttpServer.StartServer(); err != nil {
		fmt.Println("failed to start http server", err)
	}
}
