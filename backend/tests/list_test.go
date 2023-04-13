package tests

import (
	"encoding/json"
	"net/http"
	"shodo/models"
	"testing"

	"github.com/stretchr/testify/require"
)

func (s *APITestSuite) TestGetLists() {
	s.T().Run("Get all lists for a user", func(t *testing.T) {
		registerRequest := s.testData.registerModels.johnDoe
		tokens, err := s.registerUser(registerRequest)
		require.NoError(t, err)

		resp, err := s.sendGetListsRequest(tokens.Access)
		require.NoError(t, err)
		defer resp.Body.Close()

		require.Equal(t, http.StatusOK, resp.StatusCode, "Expected status code 200")

		var lists models.GetListsResponse
		err = json.NewDecoder(resp.Body).Decode(&lists)
		require.NoError(t, err)

		require.NotEmpty(t, lists.Lists, "Expected lists in the response")
		require.Equal(t, johnDoeUsername+" "+s.Config.DefaultTaskListTitle, lists.Lists[0].Title)
	})
}

// add task, get all tasks, remove task, get all tasks
func (s *APITestSuite) TestManageTasks() {
	s.T().Run("Manage tasks in default list", func(t *testing.T) {
		registerRequest := s.testData.registerModels.johnDoe
		tokens, err := s.registerUser(registerRequest)
		require.NoError(t, err)

		// get default list
		lists, err := s.getLists(tokens.Access)
		require.NoError(t, err)

		listId := lists[0].ID

		// add task
		task := models.Task{
			Title: "Test task",
		}

		addTaskRequest := models.AddTaskRequest{
			Task:   task,
			ListId: listId,
		}

		resp, err := s.sendAddTaskRequest(addTaskRequest, tokens.Access)
		require.NoError(t, err)
		defer resp.Body.Close()

		require.Equal(t, http.StatusOK, resp.StatusCode, "Expected status code 200")

		var response models.IdResponse
		err = json.NewDecoder(resp.Body).Decode(&response)
		require.NoError(t, err)

		require.NotEmpty(t, response.Id, "Expected id in the response")

		// get tasks
		resp, err = s.sendGetTasksRequest(listId, tokens.Access)
		require.NoError(t, err)
		defer resp.Body.Close()

		require.Equal(t, http.StatusOK, resp.StatusCode, "Expected status code 200", resp.Body)

		var getTaskListResponse models.GetTaskListResponse
		err = json.NewDecoder(resp.Body).Decode(&getTaskListResponse)
		require.NoError(t, err)

		require.NotEmpty(t, getTaskListResponse.List.Tasks, "Expected tasks in the response")
		require.Equal(t, task.Title, getTaskListResponse.List.Tasks[0].Title)

		// remove task
		taskId := getTaskListResponse.List.Tasks[0].ID

		removeTaskRequest := models.RemoveTaskRequest{
			TaskId: taskId,
			ListId: listId,
		}
		resp, err = s.sendRemoveTaskRequest(removeTaskRequest, tokens.Access)
		require.NoError(t, err)
		defer resp.Body.Close()

		require.Equal(t, http.StatusOK, resp.StatusCode, "Expected status code 200")

		// get tasks
		resp, err = s.sendGetTasksRequest(listId, tokens.Access)
		require.NoError(t, err)
		defer resp.Body.Close()

		require.Equal(t, http.StatusOK, resp.StatusCode, "Expected status code 200")

		err = json.NewDecoder(resp.Body).Decode(&getTaskListResponse)
		require.NoError(t, err)

		require.Empty(t, getTaskListResponse.List.Tasks, "Expected empty tasks in the response")
	})
}
