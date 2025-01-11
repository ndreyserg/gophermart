package order

import (
	"context"
	"errors"
	"fmt"

	"github.com/ndreyserg/gophermart/internal/model"
)

func (s *service) Create(ctx context.Context, number string, userID int) (*model.Order, error) {
	err := s.CheckNumber(number)

	if err != nil {
		return nil, fmt.Errorf("create order check number err: %w", err)
	}

	order, err := s.orderRepository.Create(ctx, number, userID)

	if err == nil {
		return order, nil
	}

	if !errors.Is(err, model.ErrOrderAllreadyExist) {
		return nil, fmt.Errorf("create order error: %w", err)
	}

	order, err = s.orderRepository.FindByNumber(ctx, number)

	if err != nil {
		return nil, fmt.Errorf("create order error: %w", err)
	}

	if order.UserID == userID {
		return nil, model.ErrOrderAllreadyExistOnUser
	}
	return nil, model.ErrOrderAllreadyExist
}
