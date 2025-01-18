package order

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/ndreyserg/gophermart/internal/mocks"
	"github.com/ndreyserg/gophermart/internal/model"
	"github.com/stretchr/testify/assert"
)

type deps struct {
	repo           *mocks.MockOrderRepository
	accrualService *mocks.MockAccrualService
	accountService *mocks.MockAccountService
	checker        *mocks.MockCheckerSevice
}

type tCase struct {
	name      string
	userID    int
	prepare   func(d *deps)
	expectErr string
}

func initDeps(ctrl *gomock.Controller) *deps {
	return &deps{
		repo:           mocks.NewMockOrderRepository(ctrl),
		accrualService: mocks.NewMockAccrualService(ctrl),
		accountService: mocks.NewMockAccountService(ctrl),
		checker:        mocks.NewMockCheckerSevice(ctrl),
	}
}

func TestCreateOrder(t *testing.T) {
	type tCaseCreate struct {
		tCase
		orderNumber string
	}
	tests := []tCaseCreate{
		{
			tCase: tCase{
				name:   "uncorrect number",
				userID: 1,
				prepare: func(d *deps) {
					d.checker.EXPECT().Check(gomock.Eq("1234")).Return(errors.New("unvalid"))
				},
				expectErr: "create order check number err: unvalid",
			},
			orderNumber: "1234",
		},
		//nolint:dupl //описание посдовательности вызовов
		{
			tCase: tCase{
				name:   "order allready exist",
				userID: 1,
				prepare: func(d *deps) {
					d.checker.EXPECT().Check(gomock.Eq("1234")).Return(nil)
					d.repo.EXPECT().Create(gomock.Any(), gomock.Eq("1234"),
						gomock.Eq(1)).Return(nil, model.ErrOrderAllreadyExist)
					d.repo.EXPECT().FindByNumber(gomock.Any(),
						gomock.Eq("1234")).Return(&model.Order{UserID: 3}, nil)
				},
				expectErr: model.ErrOrderAllreadyExist.Error(),
			},
			orderNumber: "1234",
		},
		//nolint:dupl //описание посдовательности вызовов
		{
			tCase: tCase{
				name:   "order allready exist on user",
				userID: 1,
				prepare: func(d *deps) {
					d.checker.EXPECT().Check(gomock.Eq("1234")).Return(nil)
					d.repo.EXPECT().Create(gomock.Any(), gomock.Eq("1234"),
						gomock.Eq(1)).Return(nil, model.ErrOrderAllreadyExist)
					d.repo.EXPECT().FindByNumber(gomock.Any(),
						gomock.Eq("1234")).Return(&model.Order{UserID: 1}, nil)
				},
				expectErr: model.ErrOrderAllreadyExistOnUser.Error(),
			},
			orderNumber: "1234",
		},
		{
			tCase: tCase{
				name:   "order created",
				userID: 1,
				prepare: func(d *deps) {
					d.checker.EXPECT().Check(gomock.Eq("1234")).Return(nil)
					d.repo.EXPECT().Create(gomock.Any(), gomock.Eq("1234"),
						gomock.Eq(1)).Return(&model.Order{}, nil)
					d.accrualService.EXPECT().AccrueAsync(gomock.Any())
				},
				expectErr: "",
			},
			orderNumber: "1234",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			d := initDeps(ctrl)
			test.prepare(d)
			s := NewService(d.repo, d.accrualService, d.accountService, d.checker)

			_, err := s.Create(context.Background(), test.orderNumber, test.userID)

			if test.expectErr != "" {
				assert.ErrorContains(t, err, test.expectErr)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestGetByUserID(t *testing.T) {
	tests := []tCase{
		{
			name:   "repo err",
			userID: 1,
			prepare: func(d *deps) {
				d.repo.EXPECT().GetByUserID(gomock.Any(), gomock.Eq(1)).Return(nil, errors.New("repo error"))
			},
			expectErr: "get user error: repo error",
		},
		{
			name:   "success",
			userID: 1,
			prepare: func(d *deps) {
				d.repo.EXPECT().GetByUserID(gomock.Any(),
					gomock.Eq(1)).Return([]*model.Order{{ID: 1}}, nil)
			},
			expectErr: "",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			d := initDeps(ctrl)
			test.prepare(d)
			s := NewService(d.repo, d.accrualService, d.accountService, d.checker)

			_, err := s.GetByUserID(context.Background(), test.userID)

			if test.expectErr != "" {
				assert.ErrorContains(t, err, test.expectErr)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
