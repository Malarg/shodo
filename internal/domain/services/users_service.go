package services

import (
	"shodo/internal/domain/helpers"
	"shodo/internal/repository"
	"shodo/models"

	"golang.org/x/net/context"
)

type UsersService struct {
	UsersRepository *repository.UsersRepository
}

func (this *UsersService) GetAllUsers(ctx context.Context, userToken string) ([]models.UserShort, error) {
	userId, err := helpers.GetUserIdFromToken(userToken)
	if err != nil {
		return nil, err
	}

	return this.UsersRepository.GetAllUsers(ctx, userId)
}
