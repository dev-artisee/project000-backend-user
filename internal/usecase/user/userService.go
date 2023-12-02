package user

import (
	"project000-backend-user/internal/repository/user"
)

type (
	UseCase interface {
	}

	service struct {
		repo user.ReaderWriter
	}
)

var _ UseCase = (*service)(nil)

func NewService(repo user.ReaderWriter) UseCase {
	return &service{repo}
}
