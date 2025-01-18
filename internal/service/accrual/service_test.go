package accrual

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/ndreyserg/gophermart/internal/mocks"
	"github.com/ndreyserg/gophermart/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestAccrualService(t *testing.T) {
	type deps struct {
		accrualRepo *mocks.MockAccrualReposity
		orderRepo   *mocks.MockOrderRepository
		accSevice   *mocks.MockAccountService
	}
	tests := []struct {
		name      string
		order     model.Order
		prepare   func(d *deps)
		expectErr string
	}{
		{
			name: "test too many requests",
			order: model.Order{
				Number: "2222",
			},
			prepare: func(d *deps) {
				d.accrualRepo.EXPECT().Get(gomock.Any()).Return(nil, model.ErrAccrualSystemTooManyRequests)
			},
			expectErr: "repeat after 1m0s",
		},
		{
			name: "accrual system error",
			order: model.Order{
				Number: "2222",
			},
			prepare: func(d *deps) {
				d.accrualRepo.EXPECT().Get(gomock.Any()).Return(nil, errors.New("internal server error"))
			},
			expectErr: "get order 2222 status error: internal server error",
		},
		{
			name: "order repeat status",
			order: model.Order{
				Number: "2222",
				Status: model.OrderStatusProcessing,
			},
			prepare: func(d *deps) {
				d.accrualRepo.EXPECT().Get(gomock.Any()).Return(
					&model.Accrual{
						Order:  "2222",
						Status: model.OrderStatusProcessing,
					},
					nil,
				)
			},
			expectErr: "repeat after 1s: order number 2222 status PROCESSING need repeat",
		},
		{
			name: "accrual is success",
			order: model.Order{
				Number: "2222",
				Status: model.OrderStatusProcessing,
				UserID: 1,
			},
			prepare: func(d *deps) {
				d.accrualRepo.EXPECT().Get(gomock.Any()).Return(
					&model.Accrual{
						Order:  "2222",
						Status: model.OrderStatusProcessed,
					},
					nil,
				)
				d.accSevice.EXPECT().GetOrCreate(gomock.Any(), gomock.Eq(1)).Return(
					&model.Account{ID: 2},
					nil,
				)
				d.orderRepo.EXPECT().UpdateAndAccrue(gomock.Any(), gomock.Any(), gomock.Eq(2)).Return(nil)
			},
			expectErr: "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			d := deps{
				accrualRepo: mocks.NewMockAccrualReposity(ctrl),
				orderRepo:   mocks.NewMockOrderRepository(ctrl),
				accSevice:   mocks.NewMockAccountService(ctrl),
			}
			test.prepare(&d)
			s := NewService(d.accrualRepo, d.orderRepo, d.accSevice)

			err := s.Accure(&test.order)

			if test.expectErr != "" {
				assert.ErrorContains(t, err, test.expectErr)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
