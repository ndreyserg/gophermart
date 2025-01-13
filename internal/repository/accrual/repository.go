package accrual

import (
	def "github.com/ndreyserg/gophermart/internal/repository"
)

var _ def.AccrualReposity = (*repository)(nil)

type repository struct {
	uri string
}

func NewRepository(uri string) *repository {
	return &repository{
		uri: uri,
	}
}
