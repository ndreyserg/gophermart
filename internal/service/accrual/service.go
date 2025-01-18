package accrual

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/ndreyserg/gophermart/internal/logger"
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

var repeatStatus = map[string]bool{
	string(model.OrderStatusRegistered): true,
	string(model.OrderStatusProcessing): true,
}

type repeatError struct {
	Err   error
	After time.Duration
}

func (e *repeatError) Error() string {
	return fmt.Sprintf("repeat after %v: %v", e.After, e.Err)
}

func newRepeatError(after time.Duration, err error) error {
	return &repeatError{
		After: after,
		Err:   err,
	}
}

type service struct {
	accrualReposity repository.AccrualReposity
	orderRepository repository.OrderRepository
	accountService  def.AccountService
}

func (s *service) AccrueAsync(order *model.Order) {
	go func(order *model.Order) {
		err := s.Accure(order)
		if err != nil {
			var re *repeatError
			logger.Log.Error(err)
			if errors.As(err, &re) {
				time.Sleep(re.After)
				s.AccrueAsync(order)
			}
		}
	}(order)
}

func (s *service) Accure(order *model.Order) error {
	acc, err := s.accrualReposity.Get(order.Number)
	if err != nil {
		if errors.Is(err, model.ErrAccrualSystemTooManyRequests) {
			return newRepeatError(time.Second*60, err)
		}
		return fmt.Errorf("get order %s status error: %w", order.Number, err)
	}

	if acc.Status != model.OrderStatusRegistered && order.Status != acc.Status {
		order.Status = acc.Status
		order.Accrual = acc.Accrual
		err := s.updateOrder(context.Background(), order)
		if err != nil {
			return fmt.Errorf("update order status error %s: %w", order.Number, err)
		}
	}

	if _, ok := repeatStatus[string(acc.Status)]; ok {
		return newRepeatError(time.Second*1, fmt.Errorf("order number %s status %s need repeat", order.Number, order.Status))
	}
	return nil
}

func (s *service) updateOrder(ctx context.Context, order *model.Order) error {
	if order.Status == model.OrderStatusProcessed {
		acc, err := s.accountService.GetOrCreate(ctx, order.UserID)
		if err != nil {
			return fmt.Errorf("get user acc error %d: %w", order.UserID, err)
		}
		err = s.orderRepository.UpdateAndAccrue(ctx, order, acc.ID)
		if err != nil {
			return fmt.Errorf("update and accrue order error %s: %w", order.Number, err)
		}
		return nil
	}

	err := s.orderRepository.Update(ctx, order)

	if err != nil {
		return fmt.Errorf("update order error %s: %w", order.Number, err)
	}
	return nil
}
