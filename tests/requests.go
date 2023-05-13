package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"shodo/models"
)

func (s *APITestSuite) sendRequest(method, url string, body []byte, authToken *string) (*http.Response, error) {
	requestReader := bytes.NewReader(body)

	req, err := http.NewRequest(method, baseUrl+url, requestReader)
	if err != nil {
		return nil, err
	}

	req.Header.Set(ContentType, ApplicationJSON)
	if authToken != nil {
		req.Header.Set(Authorization, *authToken)
	}

	client := &http.Client{}
	return client.Do(req)
}

func (s *APITestSuite) sendJSONRequest(method, url string, request interface{}, authToken *string) (*http.Response, error) {
	requestBytes, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	return s.sendRequest(method, url, requestBytes, authToken)
}

func (s *APITestSuite) sendRegisterRequest(request models.RegisterUserRequest) (*http.Response, error) {
	return s.sendJSONRequest("POST", "/api/v1/auth/register", request, nil)
}

func (s *APITestSuite) registerUser(request models.RegisterUserRequest) (*models.AuthTokens, error) {
	resp, err := s.sendRegisterRequest(request)
	if err != nil {
		return nil, err
	}
	if resp != nil {
		defer resp.Body.Close()
	}

	var response models.AuthTokens
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (s *APITestSuite) sendLoginRequest(request models.LoginUserRequest) (*http.Response, error) {
	return s.sendJSONRequest("POST", "/api/v1/auth/login", request, nil)
}

func (s *APITestSuite) sendGetListsRequest(token string) (*http.Response, error) {
	return s.sendRequest("GET", "/api/v1/lists", nil, &token)
}

func (s *APITestSuite) getLists(token string) ([]models.TaskListShort, error) {
	resp, err := s.sendGetListsRequest(token)
	defer resp.Body.Close()

	var response models.GetListsResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return response.Lists, nil
}

func (s *APITestSuite) sendAddTaskRequest(request models.AddTaskRequest, token string) (*http.Response, error) {
	return s.sendJSONRequest("POST", "/api/v1/tasks/add", request, &token)
}

func (s *APITestSuite) addTask(request models.AddTaskRequest, token string) (*models.IdResponse, error) {
	resp, err := s.sendAddTaskRequest(request, token)
	defer resp.Body.Close()

	var response models.IdResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (s *APITestSuite) sendGetTasksRequest(listId string, token string) (*http.Response, error) {
	return s.sendRequest("GET", "/api/v1/lists/"+listId, nil, &token)
}

func (s *APITestSuite) sendRemoveTaskRequest(request models.RemoveTaskRequest, token string) (*http.Response, error) {
	return s.sendJSONRequest("POST", "/api/v1/tasks/remove", request, &token)
}

func (s *APITestSuite) sendGetUsersRequest(token string) (*http.Response, error) {
	return s.sendRequest("GET", "/api/v1/users", nil, &token)
}

func (s *APITestSuite) getAllUsers(token string) ([]models.UserShort, error) {
	resp, err := s.sendGetUsersRequest(token)
	defer resp.Body.Close()

	var response []models.UserShort
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *APITestSuite) sendStartShareListRequest(request models.ShareListRequest, token string) (*http.Response, error) {
	return s.sendJSONRequest("POST", "/api/v1/share/start", request, &token)
}

func (s *APITestSuite) shareList(request models.ShareListRequest, token string) (*models.EmptyResponse, error) {
	resp, err := s.sendStartShareListRequest(request, token)
	defer resp.Body.Close()

	var response models.EmptyResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (s *APITestSuite) sendStopShareListRequest(request models.ShareListRequest, token string) (*http.Response, error) {
	return s.sendJSONRequest("POST", "/api/v1/share/stop", request, &token)
}

func (s *APITestSuite) stopShareList(request models.ShareListRequest, token string) (*models.EmptyResponse, error) {
	resp, err := s.sendStopShareListRequest(request, token)
	defer resp.Body.Close()

	var response models.EmptyResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
