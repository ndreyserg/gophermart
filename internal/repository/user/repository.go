package user

import (
	"database/sql"

	def "github.com/ndreyserg/gophermart/internal/repository"
)

var _ def.UserRepository = (*repository)(nil)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *repository {
	return &repository{
		db: db,
	}
}
