package v1

import (
	"context"
	"login-api/internal/models"
	"time"
)

func (u *useCase) UserUpdate(ctx context.Context, user_detail models.User) error {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*5000)
	defer cancel()

	return u.UserRepo.UpdateUser(ctx, user_detail)
}
