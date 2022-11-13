package v1

import (
	"context"
	"login-api/internal/models"
	"time"
)

func (u *useCase) UserCreate(ctx context.Context, user_detail models.User) error {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*5000)
	defer cancel()

	return u.UserRepo.CreateUser(ctx, user_detail)
}
