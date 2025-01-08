package account

import (
	"github.com/ndreyserg/gophermart/internal/repository"
	def "github.com/ndreyserg/gophermart/internal/service"
)

var _ def.AccountService = (*service)(nil)

func NewService(repository repository.AccountReposity) *service {
	return &service{
		accountReposity: repository,
	}
}

type service struct {
	accountReposity repository.AccountReposity
}
