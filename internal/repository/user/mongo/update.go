package user

import (
	"context"
	"login-api/internal/models"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *repository) UpdateUser(ctx context.Context, user_detail models.User) (err error) {
	_, err = r.conn.Collection("user").UpdateOne(ctx, bson.M{"username": user_detail.Username}, bson.M{"$set": user_detail})

	return
}
