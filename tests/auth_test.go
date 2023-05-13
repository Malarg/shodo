package tests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"shodo/models"

	"github.com/stretchr/testify/require"
)

func (s *APITestSuite) TestRegisterNewUser() {
	s.T().Run("Register new user", func(t *testing.T) {
		request := s.testData.registerModels.johnDoe

		resp, err := s.sendRegisterRequest(request)
		require.NoError(t, err)
		defer resp.Body.Close()

		require.Equal(t, http.StatusOK, resp.StatusCode, "Expected status code 200")

		var response models.AuthTokens
		err = json.NewDecoder(resp.Body).Decode(&response)
		require.NoError(t, err)

		require.NotEmpty(t, response.Access, "Expected access in the response")
		require.NotEmpty(t, response.Refresh, "Expected refresh in the response")
	})
}

func (s *APITestSuite) TestRegisterExistingUser() {
	s.T().Run("Register existing user", func(t *testing.T) {
		request := s.testData.registerModels.johnDoe

		_, err := s.registerUser(request)

		resp, err := s.sendRegisterRequest(request)
		require.NoError(t, err)

		require.Equal(t, http.StatusBadRequest, resp.StatusCode)

		var response models.Error
		err = json.NewDecoder(resp.Body).Decode(&response)
		require.NoError(t, err)

		require.Equal(t, fmt.Sprintf("user with email %s already exists", request.Email), response.Message)
	})
}

func (s *APITestSuite) TestLoginUser() {
	s.T().Run("Login user", func(t *testing.T) {
		registerRequest := s.testData.registerModels.johnDoe
		_, err := s.registerUser(registerRequest)
		require.NoError(t, err)

		loginRequest := s.testData.loginModels.johnDoe
		resp, err := s.sendLoginRequest(loginRequest)
		require.NoError(t, err)

		require.Equal(t, http.StatusOK, resp.StatusCode)

		var response models.AuthTokens
		err = json.NewDecoder(resp.Body).Decode(&response)
		require.NoError(t, err)

		require.NotEmpty(t, response.Access, "Expected access in the response")
		require.NotEmpty(t, response.Refresh, "Expected refresh in the response")
	})
}

func (s *APITestSuite) TestLoginUserWithWrongPassword() {
	s.T().Run("Login user with wrong password", func(t *testing.T) {
		registerRequest := s.testData.registerModels.johnDoe
		_, err := s.registerUser(registerRequest)
		require.NoError(t, err)

		loginRequest := models.LoginUserRequest{
			Email:    registerRequest.Email,
			Password: "wrong password",
		}
		resp, err := s.sendLoginRequest(loginRequest)
		require.NoError(t, err)

		require.Equal(t, http.StatusBadRequest, resp.StatusCode)

		var response models.Error
		err = json.NewDecoder(resp.Body).Decode(&response)
		require.NoError(t, err)

		require.Equal(t, "invalid credentials", response.Message)
	})
}

func (s *APITestSuite) TestLoginUserWithWrongEmail() {
	s.T().Run("Login user with wrong email", func(t *testing.T) {
		registerRequest := s.testData.registerModels.johnDoe
		_, err := s.registerUser(registerRequest)
		require.NoError(t, err)

		wrongEmail := "wrong.email@example.com"
		loginRequest := models.LoginUserRequest{
			Email:    wrongEmail,
			Password: registerRequest.Password,
		}
		resp, err := s.sendLoginRequest(loginRequest)
		require.NoError(t, err)

		require.Equal(t, http.StatusBadRequest, resp.StatusCode)

		var response models.Error
		err = json.NewDecoder(resp.Body).Decode(&response)
		require.NoError(t, err)

		//TODO: решить как выносить строки ошибок в константы
		require.Equal(t, fmt.Sprintf("user with email %s not found", wrongEmail), response.Message)
	})
}
