package tests

import (
	"net/http"
	"shodo/models"
	"testing"

	"github.com/stretchr/testify/require"
)

// + 1. Share list with another user
// + 2. Stop sharing list with user
// 3. Try to add task to list which someone shared
// 4. Try to add task to list which has not access
// 5. Try to remove task remove list which has not access
// 6. Try to get all tasks from list which has not access

type shareTestUserInput struct {
	registerRequest models.RegisterUserRequest
	tasks           []models.Task
	shareList       []string
}

type shareTestRequestInput struct {
	registerRequest models.RegisterUserRequest
	tokens          models.AuthTokens
	defautListId    string
}

func (s *APITestSuite) TestShareList2() {
	tests := []struct {
		name         string
		users        []shareTestUserInput
		request      func(t *testing.T, requestInputs []shareTestRequestInput) (resp *http.Response, err error)
		responseCode int //generic response type?
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

				for _, task := range user.tasks {
					resp, err := s.sendAddTaskRequest(
						models.AddTaskRequest{
							Task:   task,
							ListId: lists[0].ID,
						},
						tokens.Access,
					)

					require.NoError(t, err)

					require.Equal(t, http.StatusOK, resp.StatusCode, "Expected status code 200")
				}

				requestInputs = append(requestInputs, shareTestRequestInput{
					registerRequest: user.registerRequest,
					tokens:          *tokens,
					defautListId:    lists[0].ID,
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
