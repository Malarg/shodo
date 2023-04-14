package tests

import (
	"net/http"
	"shodo/models"
	"testing"

	"github.com/stretchr/testify/require"
)

// + Share list with another user
// + Share not accessible list
// + Share list with yourself
// + Share list with user which already shared

// + Stop sharing list with user
// + Stop sharing list with user which is not shared
// + Stop sharing list with yourself

// + Try to add task to list which someone shared
// + Try to add task to list which has not access

// Try to remove task from list which someone shared
// + Try to remove task remove list which has not access

// Try to get all tasks from list which someone shared
// + Try to get all tasks from list which has not access

type shareTestUserInput struct {
	registerRequest models.RegisterUserRequest
	tasks           []models.Task
	shareList       []string
}

type shareTestRequestInput struct {
	registerRequest models.RegisterUserRequest
	tokens          models.AuthTokens
	defautListId    string
	tasks           []models.Task
}

func (s *APITestSuite) TestShareList() {
	tests := []struct {
		name         string
		users        []shareTestUserInput
		request      func(t *testing.T, requestInputs []shareTestRequestInput) (resp *http.Response, err error)
		responseCode int
	}{
		{
			name: "Share list with another user",
			users: []shareTestUserInput{
				{registerRequest: s.testData.registerModels.johnDoe},
				{registerRequest: s.testData.registerModels.mikeMiles},
			},
			request: func(t *testing.T, requestInputs []shareTestRequestInput) (resp *http.Response, err error) {
				johnRi := getRequestInputByEmail(s.testData.registerModels.johnDoe.Email, requestInputs)
				mikeRi := getRequestInputByEmail(s.testData.registerModels.mikeMiles.Email, requestInputs)

				return s.sendStartShareListRequest(
					models.ShareListRequest{
						ListId: johnRi.defautListId,
						Email:  mikeRi.registerRequest.Email,
					},
					johnRi.tokens.Access,
				)
			},
			responseCode: http.StatusOK,
		},
		{
			name: "Share not accessible list",
			users: []shareTestUserInput{
				{registerRequest: s.testData.registerModels.johnDoe},
				{
					registerRequest: s.testData.registerModels.mikeMiles,
					tasks:           []models.Task{{Title: "Task 1"}},
				},
				{registerRequest: s.testData.registerModels.lukeSkywalker},
			},
			request: func(t *testing.T, requestInputs []shareTestRequestInput) (resp *http.Response, err error) {
				johnRi := getRequestInputByEmail(s.testData.registerModels.johnDoe.Email, requestInputs)
				mikeRi := getRequestInputByEmail(s.testData.registerModels.mikeMiles.Email, requestInputs)
				lukeRi := getRequestInputByEmail(s.testData.registerModels.lukeSkywalker.Email, requestInputs)

				return s.sendStartShareListRequest(
					models.ShareListRequest{
						ListId: mikeRi.defautListId,
						Email:  lukeRi.registerRequest.Email,
					},
					johnRi.tokens.Access,
				)
			},
			responseCode: http.StatusForbidden,
		},
		{
			name: "Share list with yourself",
			users: []shareTestUserInput{
				{registerRequest: s.testData.registerModels.johnDoe},
			},
			request: func(t *testing.T, requestInputs []shareTestRequestInput) (resp *http.Response, err error) {
				johnRi := getRequestInputByEmail(s.testData.registerModels.johnDoe.Email, requestInputs)

				return s.sendStartShareListRequest(
					models.ShareListRequest{
						ListId: johnRi.defautListId,
						Email:  johnRi.registerRequest.Email,
					},
					johnRi.tokens.Access,
				)
			},
			responseCode: http.StatusBadRequest,
		},
		{
			name: "Share list with user which already shared",
			users: []shareTestUserInput{
				{
					registerRequest: s.testData.registerModels.johnDoe,
					shareList:       []string{s.testData.registerModels.mikeMiles.Email},
				},
				{
					registerRequest: s.testData.registerModels.mikeMiles,
				},
			},
			request: func(t *testing.T, requestInputs []shareTestRequestInput) (resp *http.Response, err error) {
				johnRi := getRequestInputByEmail(s.testData.registerModels.johnDoe.Email, requestInputs)
				mikeRi := getRequestInputByEmail(s.testData.registerModels.mikeMiles.Email, requestInputs)

				return s.sendStartShareListRequest(
					models.ShareListRequest{
						ListId: johnRi.defautListId,
						Email:  mikeRi.registerRequest.Email,
					},
					johnRi.tokens.Access,
				)
			},
			responseCode: http.StatusBadRequest,
		},
		{
			name: "Stop share list with another user",
			users: []shareTestUserInput{
				{
					registerRequest: s.testData.registerModels.johnDoe,
					shareList:       []string{s.testData.registerModels.mikeMiles.Email},
				},
				{
					registerRequest: s.testData.registerModels.mikeMiles,
				},
			},
			request: func(t *testing.T, requestInputs []shareTestRequestInput) (resp *http.Response, err error) {
				johnRi := getRequestInputByEmail(s.testData.registerModels.johnDoe.Email, requestInputs)
				mikeRi := getRequestInputByEmail(s.testData.registerModels.mikeMiles.Email, requestInputs)

				return s.sendStopShareListRequest(
					models.ShareListRequest{
						ListId: johnRi.defautListId,
						Email:  mikeRi.registerRequest.Email,
					},
					johnRi.tokens.Access,
				)
			},
			responseCode: http.StatusOK,
		},
		{
			name: "Stop sharing list with user which is not shared",
			users: []shareTestUserInput{
				{
					registerRequest: s.testData.registerModels.johnDoe,
				},
				{
					registerRequest: s.testData.registerModels.mikeMiles,
				},
			},
			request: func(t *testing.T, requestInputs []shareTestRequestInput) (resp *http.Response, err error) {
				johnRi := getRequestInputByEmail(s.testData.registerModels.johnDoe.Email, requestInputs)
				mikeRi := getRequestInputByEmail(s.testData.registerModels.mikeMiles.Email, requestInputs)

				return s.sendStopShareListRequest(
					models.ShareListRequest{
						ListId: johnRi.defautListId,
						Email:  mikeRi.registerRequest.Email,
					},
					johnRi.tokens.Access,
				)
			},
			responseCode: http.StatusBadRequest,
		},
		{
			name: "Stop share list with yourself",
			users: []shareTestUserInput{
				{
					registerRequest: s.testData.registerModels.johnDoe,
				},
			},
			request: func(t *testing.T, requestInputs []shareTestRequestInput) (resp *http.Response, err error) {
				johnRi := getRequestInputByEmail(s.testData.registerModels.johnDoe.Email, requestInputs)

				return s.sendStopShareListRequest(
					models.ShareListRequest{
						ListId: johnRi.defautListId,
						Email:  johnRi.registerRequest.Email,
					},
					johnRi.tokens.Access,
				)
			},
			responseCode: http.StatusBadRequest,
		},
		{
			name: "Try to add task to list which someone shared",
			users: []shareTestUserInput{
				{
					registerRequest: s.testData.registerModels.johnDoe,
					shareList:       []string{s.testData.registerModels.mikeMiles.Email},
				},
				{
					registerRequest: s.testData.registerModels.mikeMiles,
				},
			},
			request: func(t *testing.T, requestInputs []shareTestRequestInput) (resp *http.Response, err error) {
				johnRi := getRequestInputByEmail(s.testData.registerModels.johnDoe.Email, requestInputs)
				mikeRi := getRequestInputByEmail(s.testData.registerModels.mikeMiles.Email, requestInputs)

				return s.sendAddTaskRequest(
					models.AddTaskRequest{
						Task:   models.Task{Title: "Task 1"},
						ListId: johnRi.defautListId,
					},
					mikeRi.tokens.Access,
				)
			},
			responseCode: http.StatusOK,
		},
		{
			name: "Try to add task to list which has not access",
			users: []shareTestUserInput{
				{
					registerRequest: s.testData.registerModels.johnDoe,
				},
				{
					registerRequest: s.testData.registerModels.mikeMiles,
				},
			},
			request: func(t *testing.T, requestInputs []shareTestRequestInput) (resp *http.Response, err error) {
				johnRi := getRequestInputByEmail(s.testData.registerModels.johnDoe.Email, requestInputs)
				mikeRi := getRequestInputByEmail(s.testData.registerModels.mikeMiles.Email, requestInputs)

				return s.sendAddTaskRequest(
					models.AddTaskRequest{
						Task:   models.Task{Title: "Task 1"},
						ListId: johnRi.defautListId,
					},
					mikeRi.tokens.Access,
				)
			},
			responseCode: http.StatusForbidden,
		},
		{
			name: "Try to remove task from list which has not access",
			users: []shareTestUserInput{
				{
					registerRequest: s.testData.registerModels.johnDoe,
					tasks: []models.Task{
						{Title: "Task 1"},
					},
				},
				{
					registerRequest: s.testData.registerModels.mikeMiles,
				},
			},
			request: func(t *testing.T, requestInputs []shareTestRequestInput) (resp *http.Response, err error) {
				johnRi := getRequestInputByEmail(s.testData.registerModels.johnDoe.Email, requestInputs)
				mikeRi := getRequestInputByEmail(s.testData.registerModels.mikeMiles.Email, requestInputs)

				return s.sendRemoveTaskRequest(
					models.RemoveTaskRequest{
						TaskId: johnRi.tasks[0].ID,
						ListId: johnRi.defautListId,
					},
					mikeRi.tokens.Access,
				)
			},
			responseCode: http.StatusForbidden,
		},
		{
			name: "Try to get all tasks from list which has not access",
			users: []shareTestUserInput{
				{
					registerRequest: s.testData.registerModels.johnDoe,
					tasks: []models.Task{
						{Title: "Task 1"},
					},
				},
				{
					registerRequest: s.testData.registerModels.mikeMiles,
				},
			},
			request: func(t *testing.T, requestInputs []shareTestRequestInput) (resp *http.Response, err error) {
				johnRi := getRequestInputByEmail(s.testData.registerModels.johnDoe.Email, requestInputs)
				mikeRi := getRequestInputByEmail(s.testData.registerModels.mikeMiles.Email, requestInputs)

				return s.sendGetTasksRequest(
					johnRi.defautListId,
					mikeRi.tokens.Access,
				)
			},
			responseCode: http.StatusForbidden,
		},
	}

	for _, test := range tests {
		s.T().Run(test.name, func(t *testing.T) {
			s.SetupTest()
			requestInputs := []shareTestRequestInput{}

			for _, user := range test.users {
				tokens, err := s.registerUser(user.registerRequest)
				require.NoError(t, err)

				lists, err := s.getLists(tokens.Access)
				require.NoError(t, err)

				require.Equal(t, 1, len(lists), "Expected 1 list")

				addedTasks := []models.Task{}
				for _, task := range user.tasks {
					idResp, err := s.addTask(
						models.AddTaskRequest{
							Task:   task,
							ListId: lists[0].ID,
						},
						tokens.Access,
					)
					require.NoError(t, err)

					addedTasks = append(addedTasks, models.Task{ID: idResp.Id, Title: task.Title})
				}

				requestInputs = append(requestInputs, shareTestRequestInput{
					registerRequest: user.registerRequest,
					tokens:          *tokens,
					defautListId:    lists[0].ID,
					tasks:           addedTasks,
				})
			}

			for _, userRequest := range test.users {
				tokens := getToken(userRequest.registerRequest, requestInputs)

				for _, email := range userRequest.shareList {
					shareListRequest := models.ShareListRequest{
						ListId: requestInputs[0].defautListId,
						Email:  email,
					}

					resp, err := s.sendStartShareListRequest(shareListRequest, tokens.Access)
					require.NoError(t, err)
					defer resp.Body.Close()

					require.Equal(t, http.StatusOK, resp.StatusCode, "Expected status code 200")
				}

			}

			response, err := test.request(t, requestInputs)
			require.NoError(t, err)

			require.Equal(t, test.responseCode, response.StatusCode, "Expected status code %d", test.responseCode)
		})
	}
}

func getToken(registerRequest models.RegisterUserRequest, requestInputs []shareTestRequestInput) models.AuthTokens {
	for _, requestInput := range requestInputs {
		if requestInput.registerRequest.Email == registerRequest.Email {
			return requestInput.tokens
		}
	}

	panic("User not found")
}

func getRequestInputByEmail(email string, requestInputs []shareTestRequestInput) shareTestRequestInput {
	for _, requestInput := range requestInputs {
		if requestInput.registerRequest.Email == email {
			return requestInput
		}
	}

	panic("User not found")
}
