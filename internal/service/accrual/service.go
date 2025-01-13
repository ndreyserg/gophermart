package accrual

import (
	"context"
	"fmt"
	"time"

	"github.com/ndreyserg/gophermart/internal/model"
	"github.com/ndreyserg/gophermart/internal/repository"
	def "github.com/ndreyserg/gophermart/internal/service"
)

var _ def.AccrualService = (*service)(nil)

func NewService(
	accrualReposity repository.AccrualReposity,
	orderRepository repository.OrderRepository,
	accountService def.AccountService,
) *service {
	return &service{
		accrualReposity: accrualReposity,
		orderRepository: orderRepository,
		accountService:  accountService,
	}
}

var repeatStatus = map[model.OrderStatus]bool{
	model.OrderStatusRegistered: true,
	model.OrderStatusProcessing: true,
}

type service struct {
	accrualReposity repository.AccrualReposity
	orderRepository repository.OrderRepository
	accountService  def.AccountService
}

func (s *service) Accrue(order *model.Order) {
	go func(order *model.Order) {
		acc, err := s.accrualReposity.Get(order.Number)
		if err != nil {
			return
		}
		if acc.Status != model.OrderStatusRegistered && order.Status != acc.Status {
			order.Status = acc.Status
			order.Accrual = acc.Accrual
			s.updateOrder(context.Background(), order)
		}

		if _, ok := repeatStatus[acc.Status]; ok {
			time.Sleep(time.Second * 1)
			s.Accrue(order)
		}

	}(order)
}

func (s *service) updateOrder(ctx context.Context, order *model.Order) error {

	if order.Status == model.OrderStatusProcessed {
		acc, err := s.accountService.GetOrCreate(ctx, order.UserID)
		if err != nil {
			return err
		}
		err = s.orderRepository.UpdateAndAccrue(ctx, order, acc.ID)
		if err != nil {
			fmt.Println(err)
		}
		return nil
	}

	err := s.orderRepository.Update(ctx, order)

	if err != nil {
		fmt.Println(err)
	}
	return nil
}
