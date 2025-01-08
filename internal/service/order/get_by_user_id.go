package order

import (
	"context"

	"github.com/ndreyserg/gophermart/internal/model"
)

func (s *service) GetByUserID(ctx context.Context, userID int) ([]*model.Order, error) {
	return s.orderRepository.GetByUserID(ctx, userID)
}
