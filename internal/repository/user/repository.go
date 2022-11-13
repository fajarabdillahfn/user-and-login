package user

import (
	"context"

	"login-api/internal/models"
)

type Repository interface {
	//create
	CreateUser(ctx context.Context, user_detail models.User) (err error)

	//retrieve
	GetUserByUserName(ctx context.Context, username string) (user models.User, err error)
	GetUserAll(ctx context.Context) (users []models.User, err error)

	//update
	UpdateUser(ctx context.Context, user_detail models.User) (err error)

	//delete
	DeleteUser(ctx context.Context, username string) (err error)
}
