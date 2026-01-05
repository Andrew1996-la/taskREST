package taskModule

import (
	"errors"
	"sync"

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
	mtx   sync.RWMutex
}

func NewTaskStore() *TaskStore {
	return &TaskStore{
		tasks: make(map[uuid.UUID]Task),
	}
}

// Create Метод добавления новой задачи в мапу
func (s *TaskStore) Create(task Task) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	s.tasks[task.Id] = task
}

// GetAll Метод получения всех задач
func (s TaskStore) GetAll() []Task {
	s.mtx.RLock()
	defer s.mtx.RUnlock()

	tasks := make([]Task, 0, len(s.tasks))

	for _, task := range s.tasks {
		tasks = append(tasks, task)
	}

	return tasks
}

// GetById Метод получения одной задачи
func (s TaskStore) GetById(id uuid.UUID) (Task, error) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()

	if _, ok := s.tasks[id]; !ok {
		return Task{}, errors.New("task not found")
	}

	return s.tasks[id], nil
}

func (s *TaskStore) Complete(id uuid.UUID) (Task, error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	task, ok := s.tasks[id]
	if !ok {
		return Task{}, errors.New("task not found")
	}

	task.IsDone = true
	s.tasks[id] = task
	return s.tasks[id], nil
}

func (s *TaskStore) Uncomplete(id uuid.UUID) (Task, error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	task, ok := s.tasks[id]
	if !ok {
		return Task{}, errors.New("task not found")
	}

	task.IsDone = false
	s.tasks[id] = task
	return s.tasks[id], nil
}

func (s TaskStore) GetUncompleted() []Task {
	s.mtx.RLock()
	defer s.mtx.RUnlock()

	tasks := make([]Task, 0, len(s.tasks))
	for _, task := range s.tasks {
		if !task.IsDone {
			tasks = append(tasks, task)
		}
	}
	return tasks
}

func (s *TaskStore) DeleteTaskById(id uuid.UUID) error {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	if _, ok := s.tasks[id]; !ok {
		return errors.New("task not found")
	}

	delete(s.tasks, id)
	return nil
}
