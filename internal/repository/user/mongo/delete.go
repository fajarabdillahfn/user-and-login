package user

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *repository) DeleteUser(ctx context.Context, username string) (err error) {
	_, err = r.conn.Collection("user").DeleteOne(ctx, bson.M{"username": username})

	return
}
