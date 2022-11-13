package user

import (
	"context"
	"login-api/internal/models"
)

func (r *repository) CreateUser(ctx context.Context, user_detail models.User) (err error) {
	_, err = r.conn.Collection("user").InsertOne(ctx, user_detail)

	return
}
