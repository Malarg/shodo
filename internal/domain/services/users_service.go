package services

import (
	"shodo/internal/domain/tokens"
	"shodo/internal/repository"
	"shodo/models"

	"golang.org/x/net/context"
)

type UsersService struct {
	UsersRepository *repository.UsersRepository
}

func (s *UsersService) GetAllUsers(ctx context.Context, userToken string) ([]models.UserShort, error) {
	userId, err := tokens.GetUserIdFromToken(userToken)
	if err != nil {
		return nil, err
	}

	return s.UsersRepository.GetAllUsers(ctx, userId)
}
