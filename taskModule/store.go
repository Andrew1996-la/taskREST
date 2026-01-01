package taskModule

import "github.com/google/uuid"

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
