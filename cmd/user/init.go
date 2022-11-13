package user

import ("go.mongodb.org/mongo-driver/mongo"

httpUser "login-api/internal/delivery/user/http"
	rUser "login-api/internal/repository/user"
	msUserRepository "login-api/internal/repository/user/mongo"
	uUser "login-api/internal/usecase/user"
	uUserV1 "login-api/internal/usecase/user/v1"
)

var (
	msUserRepo   rUser.Repository
	userUseCase  uUser.UseCase
	HTTPDelivery *httpUser.Delivery
)


func Initialize(userDB *mongo.Database) {
	msUserRepo = msUserRepository.NewUserRepo(userDB)
	userUseCase = uUserV1.NewUserUseCase(msUserRepo)
	HTTPDelivery = httpUser.NewTestApplicationHTTP(userUseCase)
}