package services

import (
	"shodo/internal/domain/helpers"
	"shodo/internal/repository"
	"shodo/models"
)

type UsersService struct {
	UsersRepository *repository.UsersRepository
}

func (this *UsersService) GetAllUsers(userToken string) ([]models.UserShort, error) {
	userId, err := helpers.GetUserIdFromToken(userToken)
	if err != nil {
		return nil, err
	}

	return this.UsersRepository.GetAllUsers(userId)
}
