package user

import (
	"context"
	"login-api/internal/models"

	"go.mongodb.org/mongo-driver/bson"
)

func(r *repository) GetUserByUserName(ctx context.Context, username string) (user models.User, err error) {
	err = r.conn.Collection("user").FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		return models.User{}, err
	}

	return 
}

func(r *repository) GetUserAll(ctx context.Context) (users []models.User, err error) {
	res, err := r.conn.Collection("user").Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	for res.Next(ctx) {
		var user models.User
		err := res.Decode(&user)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return 
}