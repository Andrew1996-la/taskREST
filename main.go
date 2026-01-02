package main

import (
	"fmt"
	"net/http"
	"taskREST/taskModule"
)

/*
TODO 1.СОЗДАТЬ ЗАДАЧУ done
TODO 2.ПОЛУЧИТЬ ИНФОРМАЦИЮ ПО ОДНОЙ ЗАДАЧЕ done
TODO 3.ПОЛУЧИТЬ ИНФОРМАЦИЮ ПО ВСЕМ ЗАДАЧАМ done
TODO 4.ОТМЕТИТЬ ОПРЕДЕЛЕННУЮ ЗАДАЧУ КАК ВЫПОЛНЕННУЮ done
TODO 5.ПОЛУЧИТЬ СПИСОК НЕВЫПОЛНЕННЫХ ЗАДАЧ done
TODO 6.УДАЛИТЬ КОНКРЕТНУЮ ЗАДАЧУ
*/
func main() {
	taskModule.RegisterTaskModule()

	fmt.Println("Task Module registered")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}

	fmt.Println("Task Module stopped")
}
