package order

import (
	"github.com/ndreyserg/gophermart/internal/repository"
	def "github.com/ndreyserg/gophermart/internal/service"
)

var _ def.OrderService = (*service)(nil)

func NewService(repository repository.OrderRepository) *service {
	return &service{
		orderRepository: repository,
	}
}

type service struct {
	orderRepository repository.OrderRepository
}
