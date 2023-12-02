package user

import (
	"database/sql"

	"gorm.io/gorm"

	"project000-backend-user/pkg/datasource"
)

type (
	ReaderWriter interface {
		datasource.TxHandler
	}

	reader interface {
	}

	writer interface {
	}

	Repository struct {
		db *gorm.DB
	}
)

var _ ReaderWriter = (*Repository)(nil)

func NewRepository(db *gorm.DB) ReaderWriter {
	return &Repository{db}
}

func (repo *Repository) Transaction(fc func(txHandler datasource.TxHandler) error, opts ...*sql.TxOptions) error {
	return repo.db.Transaction(func(tx *gorm.DB) error {
		txRepo := NewRepository(tx)

		if err := fc(txRepo); err != nil {
			return err
		}

		return nil
	}, opts...)
}
