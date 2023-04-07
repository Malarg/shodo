package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"shodo/models"

	"github.com/stretchr/testify/require"
)

const (
	baseUrl         = "http://localhost:8080"
	ContentType     = "Content-Type"
	ApplicationJSON = "application/json"
)

func (s *APITestSuite) TestRegisterNewUser() {
	s.T().Run("Register new user", func(t *testing.T) {
		request := models.RegisterUserRequest{
			Email:    "test.user@gmail.com",
			Password: "test_password123",
			Username: "John Doe",
		}

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
		request := models.RegisterUserRequest{
			Email:    "test.user@gmail.com",
			Password: "test_password123",
			Username: "John Doe",
		}

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

func (s *APITestSuite) sendRegisterRequest(request models.RegisterUserRequest) (*http.Response, error) {
	requestBytes, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	requestReader := bytes.NewReader(requestBytes)

	req, err := http.NewRequest("POST", baseUrl+"/api/v1/auth/register", requestReader)
	if err != nil {
		return nil, err
	}
	req.Header.Set(ContentType, ApplicationJSON)

	client := &http.Client{}
	return client.Do(req)
}

func (s *APITestSuite) registerUser(request models.RegisterUserRequest) (*models.AuthTokens, error) {
	resp, err := s.sendRegisterRequest(request)
	defer resp.Body.Close()

	var response models.AuthTokens
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
