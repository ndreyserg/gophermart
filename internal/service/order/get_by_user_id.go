package order

import (
	"context"
	"fmt"

	"github.com/ndreyserg/gophermart/internal/model"
)

func (s *service) GetByUserID(ctx context.Context, userID int) ([]*model.Order, error) {
	res, err := s.orderRepository.GetByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("get user error: %w", err)
	}

	return res, nil
}
