package v1

import (
	"context"
	"time"
)

func (u *useCase) UserDelete(ctx context.Context, username string) error {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*5000)
	defer cancel()

	return u.UserRepo.DeleteUser(ctx, username)
}
