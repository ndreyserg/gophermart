package order

import (
	"github.com/ndreyserg/gophermart/internal/repository"
	def "github.com/ndreyserg/gophermart/internal/service"
)

var _ def.OrderService = (*service)(nil)

type numChecker interface {
	Check(num string) error
}

func NewService(orderRepo repository.OrderRepository,
	accrualService def.AccrualService,
	accountService def.AccountService,
	checker numChecker,
) *service {
	return &service{
		orderRepository: orderRepo,
		accountService:  accountService,
		checker:         checker,
		accrualService:  accrualService,
	}
}

type service struct {
	orderRepository repository.OrderRepository
	accountService  def.AccountService
	checker         numChecker
	accrualService  def.AccrualService
}
