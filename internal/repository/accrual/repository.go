package accrual

import (
	def "github.com/ndreyserg/gophermart/internal/repository"
)

var _ def.AccrualReposity = (*repository)(nil)

type repository struct {
	addr string
}

func NewRepository(addr string) *repository {
	return &repository{
		addr: addr,
	}
}
