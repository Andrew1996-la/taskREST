package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type HTTPServer struct {
	httpHandlers *httpHandlers
}

func NewHTTPServer(httpHandlers *httpHandlers) *HTTPServer {
	return &HTTPServer{
		httpHandlers: httpHandlers,
	}
}

/*
Регистириуем стор и хендлер
*/
func (h HTTPServer) StartServer() error {
	router := mux.NewRouter()

	router.Path("/tasks").Methods(http.MethodPost).HandlerFunc(h.httpHandlers.Create)
	router.Path("/tasks").Methods(http.MethodGet).Queries("isDone", "false").HandlerFunc(h.httpHandlers.GetUncompleted)
	router.Path("/tasks").Methods(http.MethodGet).HandlerFunc(h.httpHandlers.GetAll)
	router.Path("/tasks/{id}").Methods(http.MethodGet).HandlerFunc(h.httpHandlers.GetById)
	router.Path("/tasks/{id}").Methods(http.MethodPatch).HandlerFunc(h.httpHandlers.Complete)
	router.Path("/tasks/{id}").Methods(http.MethodDelete).HandlerFunc(h.httpHandlers.DeleteById)

	fmt.Println("Task Module registered")
	return http.ListenAndServe(":8080", router)
}
