package taskModule

import "github.com/google/uuid"

// Task model = данные, а не поведение
type Task struct {
	Id          uuid.UUID
	Title       string
	Description string
	IsDone      bool
}
