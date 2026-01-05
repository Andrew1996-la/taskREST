package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"taskREST/taskModule"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

/*
принимает HTTP-запрос
проверяет метод
читает JSON
вызывает taskStore
*/

type httpHandlers struct {
	taskStore *taskModule.TaskStore
}

func NewTaskHandler(taskStore *taskModule.TaskStore) *httpHandlers {
	return &httpHandlers{
		taskStore: taskStore,
	}
}

func (h httpHandlers) Create(w http.ResponseWriter, r *http.Request) {
	var taskDto TaskDto

	// Досатаем из JSON данные и записываем в переменную input
	if err := json.NewDecoder(r.Body).Decode(&taskDto); err != nil {
		errDto := ErrorDto{
			Message: err.Error(),
		}

		http.Error(w, errDto.Message, http.StatusBadRequest)
		return
	}

	// валидация на входящих данных
	if err := taskDto.ValidateForCreate(); err != nil {
		errDto := ErrorDto{
			Message: err.Error(),
		}
		http.Error(w, errDto.Message, http.StatusBadRequest)
		return
	}

	// Формируем задачу
	task := taskModule.Task{
		Id:          uuid.New(),
		Title:       taskDto.Title,
		Description: taskDto.Description,
		IsDone:      false,
	}

	// Записываю задачу в стор. Для этого и я ссылался
	h.taskStore.Create(task)

	// Устанавливаю тип ответа, заголовок и отвчаю
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(task); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h httpHandlers) GetAll(w http.ResponseWriter, r *http.Request) {
	tasks := h.taskStore.GetAll()

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(tasks); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h httpHandlers) GetUncompleted(w http.ResponseWriter, r *http.Request) {
	uncompletedTask := h.taskStore.GetUncompleted()

	w.Header().Set("Content-Type", "application/json")

	fmt.Println(uncompletedTask)
	if err := json.NewEncoder(w).Encode(uncompletedTask); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h httpHandlers) GetById(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := uuid.Parse(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	task, err := h.taskStore.GetById(id)
	if err != nil {
		errDto := ErrorDto{
			Message: err.Error(),
		}
		http.Error(w, errDto.Message, http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(task); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h httpHandlers) Complete(w http.ResponseWriter, r *http.Request) {
	var completeDto CompleteDto
	if err := json.NewDecoder(r.Body).Decode(&completeDto); err != nil {
		errDto := ErrorDto{
			Message: err.Error(),
		}
		http.Error(w, errDto.Message, http.StatusBadRequest)
		return
	}
	idStr := mux.Vars(r)["id"]
	id, err := uuid.Parse(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if completeDto.IsDone {
		if err := h.taskStore.Complete(id); err != nil {
			errDto := ErrorDto{
				Message: err.Error(),
			}
			http.Error(w, errDto.Message, http.StatusBadRequest)
			return
		}
	} else {
		if err := h.taskStore.Uncomplete(id); err != nil {
			errDto := ErrorDto{
				Message: err.Error(),
			}
			http.Error(w, errDto.Message, http.StatusBadRequest)
			return
		}
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h httpHandlers) DeleteById(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := uuid.Parse(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := h.taskStore.DeleteTaskById(id); err != nil {
		errDto := ErrorDto{
			Message: err.Error(),
		}
		http.Error(w, errDto.Message, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
