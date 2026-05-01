package users_service

import (
	"context"
	"github.com/DSTR19/ToDo/internal/core/domain"
)

type UsersService struct {
	userRepository UsersRepository
}

type UsersRepository interface {
	CreateUser(
		ctx context.Context,
		user domain.User,
	) (domain.User, error)
}

func NewUsersService(
	UsersRepository UsersRepository,
) *UsersService {
	return &UsersService{
		userRepository: UsersRepository,
	}
}
