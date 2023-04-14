package tests

import (
	"shodo/models"
)

const (
	johnDoeUsername = "John Doe"
	jojnDoeEmail    = "john.doe@gmail.com"
	johnDoePassword = "john_doe_password123"

	mikeMilesUsername = "Mike Miles"
	mikeMilesEmail    = "mike.miles@gmail.com"
	mikeMilesPassword = "mike_miles_password123"

	lukeSkywalkerUsername = "Luke Skywalker"
	lukeSkywalkerEmail    = "luke.skywalker@gmail.com"
	lukeSkywalkerPassword = "luke_skywalker_password123"
)

type shareTestUserInput struct {
	registerRequest models.RegisterUserRequest
	tasks           []models.Task
	shareList       []string
	// TODO: add response model and compare it
}

type shareTestRequestInput struct {
	registerRequest models.RegisterUserRequest
	tokens          models.AuthTokens
	defautListId    string
	tasks           []models.Task
}

type TestData struct {
	registerModels RegisterModels
	loginModels    LoginModels
}

func (t *TestData) Init() {
	t.registerModels.InitRegisterModels()
	t.loginModels.InitLoginModels()
}

type RegisterModels struct {
	johnDoe       models.RegisterUserRequest
	mikeMiles     models.RegisterUserRequest
	lukeSkywalker models.RegisterUserRequest
}

func (m *RegisterModels) InitRegisterModels() {
	m.johnDoe = models.RegisterUserRequest{
		Email:    jojnDoeEmail,
		Password: johnDoePassword,
		Username: johnDoeUsername,
	}

	m.mikeMiles = models.RegisterUserRequest{
		Email:    mikeMilesEmail,
		Password: mikeMilesPassword,
		Username: mikeMilesUsername,
	}

	m.lukeSkywalker = models.RegisterUserRequest{
		Email:    lukeSkywalkerEmail,
		Password: lukeSkywalkerPassword,
		Username: lukeSkywalkerUsername,
	}
}

type LoginModels struct {
	johnDoe       models.LoginUserRequest
	mikeMiles     models.LoginUserRequest
	lukeSkywalker models.LoginUserRequest
}

func (m *LoginModels) InitLoginModels() {
	m.johnDoe = models.LoginUserRequest{
		Email:    jojnDoeEmail,
		Password: johnDoePassword,
	}

	m.mikeMiles = models.LoginUserRequest{
		Email:    mikeMilesEmail,
		Password: mikeMilesPassword,
	}

	m.lukeSkywalker = models.LoginUserRequest{
		Email:    lukeSkywalkerEmail,
		Password: lukeSkywalkerPassword,
	}
}
