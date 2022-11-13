package http

import uUser "login-api/internal/usecase/user"

type Delivery struct {
	userUC uUser.UseCase
}

func NewTestApplicationHTTP(userUC uUser.UseCase) *Delivery {
	return &Delivery{
		userUC: userUC,
	}
}
