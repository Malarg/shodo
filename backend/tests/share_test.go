package tests

import (
	"net/http"
	"shodo/models"
	"testing"

	"github.com/stretchr/testify/require"
)

// + 1. Share list with another user
// 2. Stop sharing list with user
// 3. Try to add task to list which someone shared
// 4. Try to add task to list which has not access
// 5. Try to remove task remove list which has not access
// 6. Try to get all tasks from list which has not access

func (s *APITestSuite) TestShareList() {
	s.T().Run("Share list with another user", func(t *testing.T) {
		registerRequest := s.testData.registerModels.johnDoe
		johnTokens, err := s.registerUser(registerRequest)
		require.NoError(t, err)

		registerRequest = s.testData.registerModels.mikeMiles
		mikeTokens, err := s.registerUser(registerRequest)
		require.NoError(t, err)

		johnLists, err := s.getLists(johnTokens.Access)
		require.NoError(t, err)

		users, err := s.getAllUsers(johnTokens.Access)
		require.NoError(t, err)

		listId := johnLists[0].ID

		mikeUserId := ""
		for _, user := range users {
			if user.Email == s.testData.registerModels.mikeMiles.Email {
				mikeUserId = user.ID
			}
		}

		shareListRequest := models.ShareListRequest{
			ListId: listId,
			UserId: mikeUserId,
		}

		resp, err := s.sendStartShareListRequest(shareListRequest, johnTokens.Access)
		require.NoError(t, err)
		defer resp.Body.Close()

		require.Equal(t, http.StatusOK, resp.StatusCode, "Expected status code 200")

		mikeLists, err := s.getLists(mikeTokens.Access)
		require.NoError(t, err)

		require.Equal(t, 2, len(mikeLists), "Expected 2 lists")
	})
}
