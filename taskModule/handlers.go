package taskModule

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

/*
принимает HTTP-запрос
проверяет метод
читает JSON
вызывает store
*/

var input struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type TaskHandler struct {
	store *TaskStore
}

func NewTaskHandler(store *TaskStore) *TaskHandler {
	return &TaskHandler{
		store: store,
	}
}

func (h TaskHandler) Create(w http.ResponseWriter, r *http.Request) {
	// Проверяю на метод запроса
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Досатаем из JSON данные и записываем в переменную input
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Формируем задачу
	task := Task{
		Id:          uuid.New(),
		Title:       input.Title,
		Description: input.Description,
		IsDone:      false,
	}

	// Записываю задачу в стор. Для этого и я ссылался
	h.store.Create(task)

	// Устанавливаю тип ответа, заголовок и отвчаю
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(task); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h TaskHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	tasks := h.store.GetAll()

	if err := json.NewEncoder(w).Encode(tasks); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
