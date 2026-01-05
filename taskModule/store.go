package taskModule

import (
	"errors"
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

// GetById Метод получения одной задачи
func (s TaskStore) GetById(id uuid.UUID) (Task, error) {
	if _, ok := s.tasks[id]; !ok {
		return Task{}, errors.New("task not found")
	}

	return s.tasks[id], nil
}

func (s *TaskStore) Complete(id uuid.UUID) error {
	task, ok := s.tasks[id]
	if !ok {
		return errors.New("task not found")
	}

	task.IsDone = true
	s.tasks[id] = task
	return nil
}

func (s *TaskStore) Uncomplete(id uuid.UUID) error {
	task, ok := s.tasks[id]
	if !ok {
		return errors.New("task not found")
	}

	task.IsDone = false
	s.tasks[id] = task
	return nil
}

func (s TaskStore) GetNotIsDone() []Task {
	tasks := make([]Task, 0, len(s.tasks))
	for _, task := range s.tasks {
		if !task.IsDone {
			tasks = append(tasks, task)
		}
	}
	return tasks
}

func (s *TaskStore) DeleteTaskById(id uuid.UUID) error {
	if _, ok := s.tasks[id]; !ok {
		return errors.New("task not found")
	}

	delete(s.tasks, id)
	return nil
}
