package v1

import (
	rUser "login-api/internal/repository/user"
	uUser "login-api/internal/usecase/user"
)

type useCase struct {
	UserRepo rUser.Repository
}

func NewUserUseCase(userRepository rUser.Repository) uUser.UseCase {
	return &useCase{
		UserRepo: userRepository,
	}
}
