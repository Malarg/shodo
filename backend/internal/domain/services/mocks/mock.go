// Code generated by MockGen. DO NOT EDIT.
// Source: services.go

// Package mock_services is a generated GoMock package.
package mock_services

import (
	reflect "reflect"
	models "shodo/internal/models"

	gomock "github.com/golang/mock/gomock"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

// MockAuthentication is a mock of Authentication interface.
type MockAuthentication struct {
	ctrl     *gomock.Controller
	recorder *MockAuthenticationMockRecorder
}

// MockAuthenticationMockRecorder is the mock recorder for MockAuthentication.
type MockAuthenticationMockRecorder struct {
	mock *MockAuthentication
}

// NewMockAuthentication creates a new mock instance.
func NewMockAuthentication(ctrl *gomock.Controller) *MockAuthentication {
	mock := &MockAuthentication{ctrl: ctrl}
	mock.recorder = &MockAuthenticationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthentication) EXPECT() *MockAuthenticationMockRecorder {
	return m.recorder
}

// IsAuthorized mocks base method.
func (m *MockAuthentication) IsAuthorized(token string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsAuthorized", token)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsAuthorized indicates an expected call of IsAuthorized.
func (mr *MockAuthenticationMockRecorder) IsAuthorized(token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsAuthorized", reflect.TypeOf((*MockAuthentication)(nil).IsAuthorized), token)
}

// LogIn mocks base method.
func (m *MockAuthentication) LogIn(request models.LoginUserRequest) (*models.AuthTokens, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LogIn", request)
	ret0, _ := ret[0].(*models.AuthTokens)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LogIn indicates an expected call of LogIn.
func (mr *MockAuthenticationMockRecorder) LogIn(request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogIn", reflect.TypeOf((*MockAuthentication)(nil).LogIn), request)
}

// MockRegistration is a mock of Registration interface.
type MockRegistration struct {
	ctrl     *gomock.Controller
	recorder *MockRegistrationMockRecorder
}

// MockRegistrationMockRecorder is the mock recorder for MockRegistration.
type MockRegistrationMockRecorder struct {
	mock *MockRegistration
}

// NewMockRegistration creates a new mock instance.
func NewMockRegistration(ctrl *gomock.Controller) *MockRegistration {
	mock := &MockRegistration{ctrl: ctrl}
	mock.recorder = &MockRegistrationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRegistration) EXPECT() *MockRegistrationMockRecorder {
	return m.recorder
}

// Register mocks base method.
func (m *MockRegistration) Register(request models.RegisterUserRequest) (*models.AuthTokens, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", request)
	ret0, _ := ret[0].(*models.AuthTokens)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Register indicates an expected call of Register.
func (mr *MockRegistrationMockRecorder) Register(request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockRegistration)(nil).Register), request)
}

// MockTaskList is a mock of TaskList interface.
type MockTaskList struct {
	ctrl     *gomock.Controller
	recorder *MockTaskListMockRecorder
}

// MockTaskListMockRecorder is the mock recorder for MockTaskList.
type MockTaskListMockRecorder struct {
	mock *MockTaskList
}

// NewMockTaskList creates a new mock instance.
func NewMockTaskList(ctrl *gomock.Controller) *MockTaskList {
	mock := &MockTaskList{ctrl: ctrl}
	mock.recorder = &MockTaskListMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTaskList) EXPECT() *MockTaskListMockRecorder {
	return m.recorder
}

// AddTaskToList mocks base method.
func (m *MockTaskList) AddTaskToList(listId *primitive.ObjectID, task *models.Task, userToken string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTaskToList", listId, task, userToken)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddTaskToList indicates an expected call of AddTaskToList.
func (mr *MockTaskListMockRecorder) AddTaskToList(listId, task, userToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTaskToList", reflect.TypeOf((*MockTaskList)(nil).AddTaskToList), listId, task, userToken)
}

// CreateDefaultTaskList mocks base method.
func (m *MockTaskList) CreateDefaultTaskList(ownerId *primitive.ObjectID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CreateDefaultTaskList", ownerId)
}

// CreateDefaultTaskList indicates an expected call of CreateDefaultTaskList.
func (mr *MockTaskListMockRecorder) CreateDefaultTaskList(ownerId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDefaultTaskList", reflect.TypeOf((*MockTaskList)(nil).CreateDefaultTaskList), ownerId)
}

// CreateTaskList mocks base method.
func (m *MockTaskList) CreateTaskList(list *models.TaskList) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CreateTaskList", list)
}

// CreateTaskList indicates an expected call of CreateTaskList.
func (mr *MockTaskListMockRecorder) CreateTaskList(list interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTaskList", reflect.TypeOf((*MockTaskList)(nil).CreateTaskList), list)
}

// IsEditListAllowed mocks base method.
func (m *MockTaskList) IsEditListAllowed(listId *primitive.ObjectID, userToken string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsEditListAllowed", listId, userToken)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsEditListAllowed indicates an expected call of IsEditListAllowed.
func (mr *MockTaskListMockRecorder) IsEditListAllowed(listId, userToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsEditListAllowed", reflect.TypeOf((*MockTaskList)(nil).IsEditListAllowed), listId, userToken)
}

// RemoveTaskFromList mocks base method.
func (m *MockTaskList) RemoveTaskFromList(listId *primitive.ObjectID, task *models.Task, userToken string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveTaskFromList", listId, task, userToken)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveTaskFromList indicates an expected call of RemoveTaskFromList.
func (mr *MockTaskListMockRecorder) RemoveTaskFromList(listId, task, userToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveTaskFromList", reflect.TypeOf((*MockTaskList)(nil).RemoveTaskFromList), listId, task, userToken)
}

// MockTokens is a mock of Tokens interface.
type MockTokens struct {
	ctrl     *gomock.Controller
	recorder *MockTokensMockRecorder
}

// MockTokensMockRecorder is the mock recorder for MockTokens.
type MockTokensMockRecorder struct {
	mock *MockTokens
}

// NewMockTokens creates a new mock instance.
func NewMockTokens(ctrl *gomock.Controller) *MockTokens {
	mock := &MockTokens{ctrl: ctrl}
	mock.recorder = &MockTokensMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTokens) EXPECT() *MockTokensMockRecorder {
	return m.recorder
}

// GenerateAndSaveTokens mocks base method.
func (m *MockTokens) GenerateAndSaveTokens(userId *primitive.ObjectID) (*models.AuthTokens, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateAndSaveTokens", userId)
	ret0, _ := ret[0].(*models.AuthTokens)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateAndSaveTokens indicates an expected call of GenerateAndSaveTokens.
func (mr *MockTokensMockRecorder) GenerateAndSaveTokens(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateAndSaveTokens", reflect.TypeOf((*MockTokens)(nil).GenerateAndSaveTokens), userId)
}

// GetTokens mocks base method.
func (m *MockTokens) GetTokens(userId string) (*models.AuthTokens, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTokens", userId)
	ret0, _ := ret[0].(*models.AuthTokens)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTokens indicates an expected call of GetTokens.
func (mr *MockTokensMockRecorder) GetTokens(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTokens", reflect.TypeOf((*MockTokens)(nil).GetTokens), userId)
}