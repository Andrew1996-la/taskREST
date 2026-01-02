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
		case http.MethodPost:
			handler.Create(w, r)
		case http.MethodGet:
			handler.GetAll(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/task/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handler.GetById(w, r)
		case http.MethodPatch:
			handler.SetIsDone(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

}
