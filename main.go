package main

import (
	"fmt"
	"taskREST/taskModule"
	"taskREST/taskModule/http"
)

func main() {
	taskStore := taskModule.NewTaskStore()
	handlers := http.NewTaskHandler(taskStore)
	server := http.NewHTTPServer(handlers)

	if err := server.StartServer(); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
