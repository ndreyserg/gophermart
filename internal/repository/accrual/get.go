package accrual

import "github.com/ndreyserg/gophermart/internal/model"

func (r repository) Get(orderNumber string) (*model.Accrual, error) {
	res := model.Accrual{
		Order:   orderNumber,
		Status:  model.OrderStatusProcessed,
		Accrual: 200.5,
	}
	return &res, nil
}
