package account

import (
	"github.com/ndreyserg/gophermart/internal/repository"
	def "github.com/ndreyserg/gophermart/internal/service"
)

var _ def.AccountService = (*service)(nil)

type numChecker interface {
	Check(num string) error
}

func NewService(repo repository.AccountReposity, checker numChecker) *service {
	return &service{
		accountReposity: repo,
		checker:         checker,
	}
}

type service struct {
	accountReposity repository.AccountReposity
	checker         numChecker
}
