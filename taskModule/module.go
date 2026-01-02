package taskModule

import "net/http"

/*
RegisterTaskModule
Регистириуем стор и хендлер
*/
func RegisterTaskModule() {
	store := NewTaskStore()
	handler := NewTaskHandler(store)

	http.HandleFunc("/task", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			handler.Create(w, r)
		case "GET":
			handler.GetAll(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/task/id", handler.GetById)

}
