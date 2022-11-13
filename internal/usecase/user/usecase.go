package user

import (
	"context"

	"login-api/internal/models"
)

type UseCase interface {
	//create
	UserCreate(ctx context.Context, user_detail models.User) (err error)

	//retrieve
	UserGetByUsername(ctx context.Context, username string) (models.User, error)
	GetUserAll(ctx context.Context) (users []models.User, err error)

	//update
	UserUpdate(ctx context.Context, user_detail models.User) (err error)

	//delete
	//delete
	UserDelete(ctx context.Context, username string) (err error)
}
