package user

import (
	"log"

	"go.mongodb.org/mongo-driver/mongo"

	rUser "login-api/internal/repository/user"
)

type repository struct {
	conn *mongo.Database
}

func NewUserRepo(db *mongo.Database) rUser.Repository {
	if db == nil {
		log.Panic("missing database connection")
	}

	repo := &repository{
		conn: db,
	}

	return repo
}
