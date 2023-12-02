package datasource

import (
	"database/sql"
)

type TxHandler interface {
	Transaction(fc func(tx TxHandler) error, opts ...*sql.TxOptions) error
}
