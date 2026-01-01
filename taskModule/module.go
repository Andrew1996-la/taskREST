package taskModule

import "net/http"

/*
RegisterTaskModule
Регистириуем стор и хендлер
*/
func RegisterTaskModule() {
	store := NewTaskStore()
	handler := NewTaskHandler(store)

	http.HandleFunc("/task", handler.Create)
}
