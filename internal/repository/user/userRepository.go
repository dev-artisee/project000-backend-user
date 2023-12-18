package user

import (
	"database/sql"

	"gorm.io/gorm"

	"project000-backend-user/pkg/datasource"
)

type (
	ReadWriter interface {
		datasource.TxHandler
		reader
		writer
	}

	reader interface {
	}

	writer interface {
	}

	repository struct {
		db *gorm.DB
	}
)

var (
	_ datasource.TxHandler = (*repository)(nil)
	_ ReadWriter           = (*repository)(nil)
)

func NewRepository(db *gorm.DB) ReadWriter {
	return &repository{db}
}

func (repo *repository) Transaction(fc func(txHandler datasource.TxHandler) error, opts ...*sql.TxOptions) error {
	return repo.db.Transaction(func(tx *gorm.DB) error {
		txRepo := NewRepository(tx)

		if err := fc(txRepo); err != nil {
			return err
		}

		return nil
	}, opts...)
}
