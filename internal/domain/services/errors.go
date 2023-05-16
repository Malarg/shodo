package services

import (
	"fmt"
)

type NotAllowedError struct{}

func (e *NotAllowedError) Error() string {
	return fmt.Errorf("Operation not allowed.").Error()
}

type ListNotFoundError struct {
	ListId string
}

func (e *ListNotFoundError) Error() string {
	return fmt.Errorf("List with id %s not found.", e.ListId).Error()
}

type TaskNotFoundError struct {
	TaskId string
}

func (e *TaskNotFoundError) Error() string {
	return fmt.Errorf("Task with id %s not found.", e.TaskId).Error()
}

type InternalError struct {
	Err error
}

func (e *InternalError) Error() string {
	return fmt.Errorf("Internal error: %s", e.Err.Error()).Error()
}
