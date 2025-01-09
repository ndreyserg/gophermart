package order

import (
	"context"
	"errors"

	"github.com/ndreyserg/gophermart/internal/model"
)

func (s *service) Create(ctx context.Context, number string, userID int) (*model.Order, error) {

	err := s.CheckNumber(number)

	if err != nil {
		return nil, err
	}

	order, err := s.orderRepository.Create(ctx, number, userID)

	if err == nil {
		return order, nil
	}

	if !errors.Is(err, model.ErrorOrderAllreadyExist) {
		return nil, err
	}

	order, err = s.orderRepository.FindByNumber(ctx, number)

	if err != nil {
		return nil, err
	}

	if order.UserID == userID {
		return nil, model.ErrorOrderAllreadyExistOnUser
	}
	return nil, model.ErrorOrderAllreadyExist
}
