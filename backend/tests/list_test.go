package tests

import (
	"encoding/json"
	"net/http"
	"shodo/models"
	"testing"

	"github.com/stretchr/testify/require"
)

// Get all lists for a user
// Get all lists for user with wrong token

// Add task to a list
// Add task to a list with wrong token
// Add task to a list with wrong list id

// Remove task from a list
// Remove task from a list with wrong token
// Remove task from a list with wrong list id
// Remove task from a list with wrong task id

// Get all tasks for a list
// Get all tasks for a list with wrong token
// Get all tasks for a list with wrong list id

func (s *APITestSuite) TestLists() {
	tests := []struct {
		name            string
		users           []testUserInput
		request         func(t *testing.T, requestInputs []testRequestInput) (resp *http.Response, err error)
		responseCode    int
		responseChecker func(t *testing.T, resp *http.Response, requestInputs []testRequestInput)
	}{
		{
			name: "Get all lists for a user",
			users: []testUserInput{
				{
					registerRequest: s.testData.registerModels.johnDoe,
				},
			},
			request: func(t *testing.T, requestInputs []testRequestInput) (resp *http.Response, err error) {
				johnRi := getRequestInputByEmail(s.testData.registerModels.johnDoe.Email, requestInputs)
				return s.sendGetListsRequest(johnRi.tokens.Access)
			},
			responseCode: http.StatusOK,
			responseChecker: func(t *testing.T, resp *http.Response, requestInputs []testRequestInput) {
				var response models.GetListsResponse
				err := json.NewDecoder(resp.Body).Decode(&response)
				require.NoError(t, err)
				require.Equal(t, 1, len(response.Lists))
			},
		},

		{
			name: "Get all lists for user with wrong token",
			users: []testUserInput{
				{
					registerRequest: s.testData.registerModels.johnDoe,
				},
			},
			request: func(t *testing.T, requestInputs []testRequestInput) (resp *http.Response, err error) {
				return s.sendGetListsRequest("wrong token")
			},
			responseCode: http.StatusUnauthorized,
		},

		{
			name: "Add task to a list",
			users: []testUserInput{
				{
					registerRequest: s.testData.registerModels.johnDoe,
				},
			},
			request: func(t *testing.T, requestInputs []testRequestInput) (resp *http.Response, err error) {
				johnRi := getRequestInputByEmail(s.testData.registerModels.johnDoe.Email, requestInputs)
				return s.sendAddTaskRequest(models.AddTaskRequest{
					Task:   s.testData.taskModels.task1,
					ListId: johnRi.defautListId,
				}, johnRi.tokens.Access)
			},
			responseCode: http.StatusOK,
			responseChecker: func(t *testing.T, resp *http.Response, requestInputs []testRequestInput) {
				var response models.IdResponse
				err := json.NewDecoder(resp.Body).Decode(&response)
				require.NoError(t, err)
				require.NotEmpty(t, response.Id)
			},
		},

		{
			name: "Add task to a list with wrong token",
			users: []testUserInput{
				{
					registerRequest: s.testData.registerModels.johnDoe,
				},
			},
			request: func(t *testing.T, requestInputs []testRequestInput) (resp *http.Response, err error) {
				johnRi := getRequestInputByEmail(s.testData.registerModels.johnDoe.Email, requestInputs)
				return s.sendAddTaskRequest(models.AddTaskRequest{
					Task:   s.testData.taskModels.task1,
					ListId: johnRi.defautListId,
				}, "wrong token")
			},
			responseCode: http.StatusUnauthorized,
		},

		{
			name: "Add task to a list with wrong list id",
			users: []testUserInput{
				{
					registerRequest: s.testData.registerModels.johnDoe,
				},
			},
			request: func(t *testing.T, requestInputs []testRequestInput) (resp *http.Response, err error) {
				johnRi := getRequestInputByEmail(s.testData.registerModels.johnDoe.Email, requestInputs)
				return s.sendAddTaskRequest(models.AddTaskRequest{
					Task:   s.testData.taskModels.task1,
					ListId: "wrong list id",
				}, johnRi.tokens.Access)
			},
			responseCode: http.StatusNotFound,
		},

		{
			name: "Remove task from a list",
			users: []testUserInput{
				{
					registerRequest: s.testData.registerModels.johnDoe,
					tasks: []models.Task{
						s.testData.taskModels.task1,
					},
				},
			},
			request: func(t *testing.T, requestInputs []testRequestInput) (resp *http.Response, err error) {
				johnRi := getRequestInputByEmail(s.testData.registerModels.johnDoe.Email, requestInputs)
				return s.sendRemoveTaskRequest(
					models.RemoveTaskRequest{
						TaskId: johnRi.tasks[0].ID,
						ListId: johnRi.defautListId,
					},
					johnRi.tokens.Access,
				)
			},
			responseCode: http.StatusOK,
			responseChecker: func(t *testing.T, resp *http.Response, requestInputs []testRequestInput) {
				//find user with john doe email
				//find list where he is an owner
				//find task with name task1
				//check that task is not in the list
			},
		},

		{
			name: "Remove task from a list with wrong token",
			users: []testUserInput{
				{
					registerRequest: s.testData.registerModels.johnDoe,
					tasks: []models.Task{
						s.testData.taskModels.task1,
					},
				},
			},
			request: func(t *testing.T, requestInputs []testRequestInput) (resp *http.Response, err error) {
				johnRi := getRequestInputByEmail(s.testData.registerModels.johnDoe.Email, requestInputs)
				return s.sendRemoveTaskRequest(
					models.RemoveTaskRequest{
						TaskId: johnRi.tasks[0].ID,
						ListId: johnRi.defautListId,
					},
					"wrong token",
				)
			},
			responseCode: http.StatusUnauthorized,
		},

		{
			name: "Remove task from a list with wrong list id",
			users: []testUserInput{
				{
					registerRequest: s.testData.registerModels.johnDoe,
					tasks: []models.Task{
						s.testData.taskModels.task1,
					},
				},
			},
			request: func(t *testing.T, requestInputs []testRequestInput) (resp *http.Response, err error) {
				johnRi := getRequestInputByEmail(s.testData.registerModels.johnDoe.Email, requestInputs)
				return s.sendRemoveTaskRequest(
					models.RemoveTaskRequest{
						TaskId: johnRi.tasks[0].ID,
						ListId: "wrong list id",
					},
					johnRi.tokens.Access,
				)
			},
			responseCode: http.StatusNotFound,
		},

		{
			name: "Remove task from a list with wrong task id",
			users: []testUserInput{
				{
					registerRequest: s.testData.registerModels.johnDoe,
					tasks: []models.Task{
						s.testData.taskModels.task1,
					},
				},
			},
			request: func(t *testing.T, requestInputs []testRequestInput) (resp *http.Response, err error) {
				johnRi := getRequestInputByEmail(s.testData.registerModels.johnDoe.Email, requestInputs)
				return s.sendRemoveTaskRequest(
					models.RemoveTaskRequest{
						TaskId: "wrong task id",
						ListId: johnRi.defautListId,
					},
					johnRi.tokens.Access,
				)
			},
			responseCode: http.StatusNotFound,
		},

		{
			name: "Get tasks from a list",
			users: []testUserInput{
				{
					registerRequest: s.testData.registerModels.johnDoe,
					tasks: []models.Task{
						s.testData.taskModels.task1,
						s.testData.taskModels.task2,
					},
				},
			},
			request: func(t *testing.T, requestInputs []testRequestInput) (resp *http.Response, err error) {
				johnRi := getRequestInputByEmail(s.testData.registerModels.johnDoe.Email, requestInputs)
				return s.sendGetTasksRequest(johnRi.defautListId, johnRi.tokens.Access)
			},
			responseCode: http.StatusOK,
			responseChecker: func(t *testing.T, resp *http.Response, requestInputs []testRequestInput) {
				johnRi := getRequestInputByEmail(s.testData.registerModels.johnDoe.Email, requestInputs)
				var tasks models.GetTaskListResponse
				err := json.NewDecoder(resp.Body).Decode(&tasks)
				require.NoError(t, err)
				require.Len(t, tasks.List.Tasks, len(johnRi.tasks))

				var taskTitles []string
				for _, task := range johnRi.tasks {
					taskTitles = append(taskTitles, task.Title)
				}

				for _, task := range tasks.List.Tasks {
					require.Contains(t, taskTitles, task.Title)
				}
			},
		},

		{
			name: "Get tasks from a list with wrong token",
			users: []testUserInput{
				{
					registerRequest: s.testData.registerModels.johnDoe,
					tasks: []models.Task{
						s.testData.taskModels.task1,
						s.testData.taskModels.task2,
					},
				},
			},
			request: func(t *testing.T, requestInputs []testRequestInput) (resp *http.Response, err error) {
				johnRi := getRequestInputByEmail(s.testData.registerModels.johnDoe.Email, requestInputs)
				return s.sendGetTasksRequest(johnRi.defautListId, "wrong token")
			},
			responseCode: http.StatusUnauthorized,
		},

		{
			name: "Get tasks from a list with wrong list id",
			users: []testUserInput{
				{
					registerRequest: s.testData.registerModels.johnDoe,
					tasks: []models.Task{
						s.testData.taskModels.task1,
						s.testData.taskModels.task2,
					},
				},
			},
			request: func(t *testing.T, requestInputs []testRequestInput) (resp *http.Response, err error) {
				johnRi := getRequestInputByEmail(s.testData.registerModels.johnDoe.Email, requestInputs)
				return s.sendGetTasksRequest("wrong list id", johnRi.tokens.Access)
			},
			responseCode: http.StatusNotFound,
		},
	}

	for _, test := range tests {
		s.Run(test.name, func() {
			s.SetupTest()
			var requestInputs []testRequestInput
			for _, user := range test.users {
				tokens, err := s.registerUser(user.registerRequest)
				require.NoError(s.T(), err)

				lists, err := s.getLists(tokens.Access)
				require.NoError(s.T(), err)

				listId := lists[0].ID

				addedTasks := make([]models.Task, 0)
				for _, task := range user.tasks {
					id, err := s.addTask(models.AddTaskRequest{Task: task, ListId: listId}, tokens.Access)
					require.NoError(s.T(), err)

					task.ID = id.Id
					addedTasks = append(addedTasks, task)
				}

				for _, email := range user.shareList {
					_, err := s.shareList(models.ShareListRequest{Email: email, ListId: listId}, tokens.Access)
					require.NoError(s.T(), err)
				}

				requestInputs = append(requestInputs, testRequestInput{
					registerRequest: user.registerRequest,
					tokens:          *tokens,
					defautListId:    lists[0].ID,
					tasks:           addedTasks,
				})
			}

			resp, err := test.request(s.T(), requestInputs)
			require.NoError(s.T(), err)
			require.Equal(s.T(), test.responseCode, resp.StatusCode)

			if test.responseChecker != nil {
				test.responseChecker(s.T(), resp, requestInputs)
			}
		})
	}
}
