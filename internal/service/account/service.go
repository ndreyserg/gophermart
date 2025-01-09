package account

import (
	"github.com/ndreyserg/gophermart/internal/repository"
	def "github.com/ndreyserg/gophermart/internal/service"
)

var _ def.AccountService = (*service)(nil)

type orderService interface {
	CheckNumber(num string) error
}

func NewService(repository repository.AccountReposity, orders orderService) *service {
	return &service{
		accountReposity: repository,
		orders:          orders,
	}
}

type service struct {
	accountReposity repository.AccountReposity
	orders          orderService
}
