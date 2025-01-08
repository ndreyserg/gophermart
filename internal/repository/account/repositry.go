package account

import (
	"database/sql"

	def "github.com/ndreyserg/gophermart/internal/repository"
)

var _ def.AccountReposity = (*repository)(nil)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *repository {
	return &repository{
		db: db,
	}
}
