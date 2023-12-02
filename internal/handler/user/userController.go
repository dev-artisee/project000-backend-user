package user

import (
	usecase "project000-backend-user/internal/usecase/user"
)

type (
	Handler interface {
	}

	controller struct {
		service usecase.UseCase
	}
)

var _ Handler = (*controller)(nil)

func NewController(service usecase.UseCase) Handler {
	return &controller{service}
}
