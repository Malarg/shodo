package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"shodo/models"
)

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

func (s *APITestSuite) sendLoginRequest(request models.LoginUserRequest) (*http.Response, error) {
	requestBytes, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	requestReader := bytes.NewReader(requestBytes)

	req, err := http.NewRequest("POST", baseUrl+"/api/v1/auth/login", requestReader)
	if err != nil {
		return nil, err
	}
	req.Header.Set(ContentType, ApplicationJSON)

	client := &http.Client{}
	return client.Do(req)
}
