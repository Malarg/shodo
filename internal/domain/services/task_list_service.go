package services

import (
	"context"
	"fmt"
	"shodo/internal/domain/tokens"
	"shodo/internal/repository"
	"shodo/models"

	"golang.org/x/exp/slices"
)

const (
	defaultTaskListTitle = "Shoppings"
	notAllowedMessage    = "operation permitted only for list owner or shared users"
)

type TaskListService struct {
	TaskListRepository    repository.TaskList
	UsersRepository       repository.Users
	AuthenticationService *AuthenticationService
}

func (s *TaskListService) CreateDefaultTaskList(ctx context.Context, username, ownerId string) {
	s.TaskListRepository.CreateTaskList(ctx, &models.TaskList{Title: username + " " + defaultTaskListTitle, Owner: ownerId})
}

func (s *TaskListService) CreateTaskList(ctx context.Context, list *models.TaskList) {
	s.TaskListRepository.CreateTaskList(ctx, list)
}

func (s *TaskListService) GetTaskLists(ctx context.Context, userToken string) ([]models.TaskListShort, error) {
	userId, err := tokens.GetUserIdFromToken(userToken)
	if err != nil {
		return nil, err
	}

	return s.TaskListRepository.GetTaskLists(ctx, userId)
}

func (s *TaskListService) GetTaskList(ctx context.Context, listId *string, userToken string) (models.TaskList, error) {
	isListExists, err := s.TaskListRepository.CheckTaskListExists(ctx, *listId)
	if !isListExists {
		return models.TaskList{}, &ListNotFoundError{ListId: *listId}
	}

	isEditListAllowed, err := s.IsEditListAllowed(ctx, listId, userToken)
	if err != nil {
		return models.TaskList{}, &InternalError{Err: err}
	}

	if !isEditListAllowed {
		return models.TaskList{}, &NotAllowedError{}
	}

	resp, err := s.TaskListRepository.GetTaskList(ctx, listId)
	if err != nil {
		return models.TaskList{}, &InternalError{Err: err}
	}

	return resp, nil
}

func (s *TaskListService) AddTaskToList(ctx context.Context, listId *string, task *models.Task, userToken string) (*string, error) {
	isListExists, err := s.TaskListRepository.CheckTaskListExists(ctx, *listId)
	if !isListExists {
		return nil, &ListNotFoundError{ListId: *listId}
	}

	isEditListAllowed, err := s.IsEditListAllowed(ctx, listId, userToken)
	if err != nil {
		return nil, &InternalError{Err: err}
	}

	if !isEditListAllowed {
		return nil, &NotAllowedError{}
	}

	taskId, err := s.TaskListRepository.AddTaskToList(ctx, listId, task)
	if err != nil {
		return nil, &InternalError{Err: err}
	}

	return taskId, nil
}

func (s *TaskListService) RemoveTaskFromList(ctx context.Context, listId *string, taskId *string, userToken string) error {
	isListExists, err := s.TaskListRepository.CheckTaskListExists(ctx, *listId)
	if !isListExists {
		return &ListNotFoundError{ListId: *listId}
	}

	isTaskExists, err := s.TaskListRepository.CheckTaskExists(ctx, *listId, *taskId)
	if !isTaskExists {
		return &TaskNotFoundError{TaskId: *taskId}
	}

	isEditListAllowed, err := s.IsEditListAllowed(ctx, listId, userToken)
	if err != nil {
		return &InternalError{Err: err}
	}

	if !isEditListAllowed {
		return &NotAllowedError{}
	}

	err = s.TaskListRepository.RemoveTaskFromList(ctx, listId, taskId)
	if err != nil {
		return &InternalError{Err: err}
	}

	return nil
}

func (s *TaskListService) IsEditListAllowed(ctx context.Context, listId *string, userToken string) (bool, error) {
	list, err := s.TaskListRepository.GetTaskList(ctx, listId)
	if err != nil {
		return false, err
	}

	userId, err := tokens.GetUserIdFromToken(userToken)
	if err != nil {
		return false, err
	}

	if userId == list.Owner || slices.Contains(list.SharedWith, userId) {
		return true, nil
	}

	return false, nil
}

func (s *TaskListService) StartShareWithUser(ctx context.Context, listId *string, email *string, userToken string) error {
	selfId, err := tokens.GetUserIdFromToken(userToken)
	if err != nil {
		return err
	}

	user, err := s.UsersRepository.GetUserByEmail(ctx, *email)
	if selfId == *&user.ID {
		return fmt.Errorf("can't share with yourself")
	}

	list, err := s.TaskListRepository.GetTaskList(ctx, listId)
	if err != nil {
		return err
	}

	if list.Owner != selfId {
		return &NotAllowedError{}
	}

	if slices.Contains(list.SharedWith, *&user.ID) {
		return fmt.Errorf("user already shared")
	}

	err = s.TaskListRepository.AddUserToList(ctx, *listId, *email)

	if err != nil {
		return err
	}

	return nil
}

func (s *TaskListService) StopShareWithUser(ctx context.Context, listId *string, email *string, userToken string) error {
	selfId, err := tokens.GetUserIdFromToken(userToken)
	if err != nil {
		return err
	}

	user, err := s.UsersRepository.GetUserByEmail(ctx, *email)
	if selfId == *&user.ID {
		return fmt.Errorf("can't stop share with yourself")
	}

	list, err := s.TaskListRepository.GetTaskList(ctx, listId)
	if err != nil {
		return err
	}

	if list.Owner != selfId {
		return &NotAllowedError{}
	}

	if !slices.Contains(list.SharedWith, user.ID) {
		return fmt.Errorf("user not shared")
	}

	err = s.TaskListRepository.RemoveUserFromList(ctx, *listId, *email)

	if err != nil {
		return err
	}

	return nil
}
