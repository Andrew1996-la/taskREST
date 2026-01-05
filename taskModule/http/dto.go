package http

import "errors"

// DTO - data transfer object
type TaskDto struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (t TaskDto) ValidateForCreate() error {
	if t.Title == "" {
		return errors.New("title is required")
	}

	if t.Description == "" {
		return errors.New("description is required")
	}

	return nil
}

type ErrorDto struct {
	Message string `json:"message"`
}

type CompleteDto struct {
	IsDone bool `json:"isDone"`
}
