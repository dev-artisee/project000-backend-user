package user

import (
	userRepo "project000-backend-user/internal/repository/user"
	"project000-backend-user/pkg/datasource"
)

type (
	UseCase interface {
	}

	service struct {
		repo userRepo.ReadWriter
	}
)

var _ UseCase = (*service)(nil)

func NewService(repo userRepo.ReadWriter) UseCase {
	return &service{repo}
}

func (s *service) Sample() error {
	err := s.repo.Transaction(func(txHandler datasource.TxHandler) error {
		userRepo, ok := txHandler.(userRepo.ReadWriter)
		if !ok {
			return cerror.ErrInternalServer
		}

		// TODO: do something with txRepo

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
