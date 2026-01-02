package main

import (
	"fmt"
	"net/http"
	"taskREST/taskModule"
)

func main() {
	taskModule.RegisterTaskModule()

	fmt.Println("Task Module registered")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}

	fmt.Println("Task Module stopped")
}
