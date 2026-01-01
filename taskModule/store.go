package taskModule

import (
	"fmt"

	"github.com/google/uuid"
)

/*
	TaskStore

ничего не знает про HTTP
ничего не знает про JSON
просто хранит данные
*/
type TaskStore struct {
	tasks map[uuid.UUID]Task
}

func NewTaskStore() *TaskStore {
	return &TaskStore{
		tasks: make(map[uuid.UUID]Task),
	}
}

// Create Метод добавления новой задачи в мапу
func (s *TaskStore) Create(task Task) {
	s.tasks[task.Id] = task
}

// GetAll Метод получения всех задач
func (s TaskStore) GetAll() []Task {
	tasks := make([]Task, 0, len(s.tasks))
	fmt.Println(tasks)
	for _, task := range s.tasks {
		tasks = append(tasks, task)
	}

	return tasks
}
