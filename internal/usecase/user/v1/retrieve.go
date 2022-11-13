package v1

import (
	"context"
	"time"

	"login-api/internal/models"
)

func (u *useCase) UserGetByUsername(ctx context.Context, username string) (models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*5000)
	defer cancel()

	return u.UserRepo.GetUserByUserName(ctx, username)
}

func (u *useCase) GetUserAll(ctx context.Context) (users []models.User, err error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*5000)
	defer cancel()

	return u.UserRepo.GetUserAll(ctx)
}
