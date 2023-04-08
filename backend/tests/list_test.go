package tests

import (
	"encoding/json"
	"net/http"
	"shodo/models"
	"testing"

	"github.com/stretchr/testify/require"
)

// 1. Get all lists for a user
// 2. Try to get all tasks for a list
// 3. Add task to a user
// 4. Delete task
// 5. Share list with another user
// 6. Stop sharing list with user
// 7. Try to add task to list which someone shared
// 8. Try to add task to list which has not access
// 9. Try to remove task remove list which has not access
// 10. Try to get all tasks from list which has not access

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
		require.Equal(t, s.Config.DefaultTaskListTitle, lists.Lists[0].Title)
	})
}
