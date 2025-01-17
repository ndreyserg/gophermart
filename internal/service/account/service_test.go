package account

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
	accRepo *mocks.MockAccountReposity
	checker *mocks.MockCheckerSevice
}

type tCase struct {
	name      string
	userID    int
	prepare   func(d *deps)
	expectErr string
}

func initDeps(ctrl *gomock.Controller) *deps {
	return &deps{
		accRepo: mocks.NewMockAccountReposity(ctrl),
		checker: mocks.NewMockCheckerSevice(ctrl),
	}
}

func TestGetOrCreate(t *testing.T) {
	tests := []tCase{
		{
			name:   "repo error",
			userID: 1,
			prepare: func(d *deps) {
				d.accRepo.EXPECT().FindByUserID(gomock.Any(), gomock.Eq(1)).Return(
					nil, errors.New("repo err"))
			},
			expectErr: "get or create acc error: repo err",
		},
		{
			name:   "created",
			userID: 1,
			prepare: func(d *deps) {
				d.accRepo.EXPECT().FindByUserID(gomock.Any(), gomock.Eq(1)).Return(
					nil, model.ErrAccountNotFound)
				d.accRepo.EXPECT().Create(gomock.Any(), gomock.Eq(1)).Return(
					&model.Account{ID: 1},
					nil,
				)
			},
			expectErr: "",
		},
		{
			name:   "success find",
			userID: 1,
			prepare: func(d *deps) {
				d.accRepo.EXPECT().FindByUserID(gomock.Any(), gomock.Eq(1)).Return(&model.Account{ID: 1}, nil)
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
			s := NewService(d.accRepo, d.checker)

			_, err := s.GetOrCreate(context.Background(), test.userID)

			if test.expectErr != "" {
				assert.ErrorContains(t, err, test.expectErr)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestGetBalance(t *testing.T) {
	type tCaseBalance struct {
		tCase
		resultBalance  float64
		resultWithdawn float64
	}
	tests := []tCaseBalance{
		{
			tCase: tCase{
				name:   "succes get balance",
				userID: 1,
				prepare: func(d *deps) {
					d.accRepo.EXPECT().FindByUserID(gomock.Any(),
						gomock.Eq(1)).Return(&model.Account{ID: 2, Balance: 100}, nil)
					d.accRepo.EXPECT().GetWithdrawn(gomock.Any(),
						gomock.Eq(2)).Return(float64(200), nil)
				},
				expectErr: "",
			},
			resultBalance:  100,
			resultWithdawn: 200,
		},
		{
			tCase: tCase{
				name:   "repo error",
				userID: 1,
				prepare: func(d *deps) {
					d.accRepo.EXPECT().FindByUserID(gomock.Any(),
						gomock.Eq(1)).Return(&model.Account{ID: 2, Balance: 100}, nil)
					d.accRepo.EXPECT().GetWithdrawn(gomock.Any(),
						gomock.Eq(2)).Return(float64(0), errors.New("repo error"))
				},
				expectErr: "get balance get withdrawn error: repo error",
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			d := initDeps(ctrl)
			test.prepare(d)
			s := NewService(d.accRepo, d.checker)

			balance, err := s.GetBalance(context.Background(), test.userID)

			if test.expectErr != "" {
				assert.ErrorContains(t, err, test.expectErr)
			} else {
				assert.Nil(t, err)
				assert.Equal(
					t,
					test.resultBalance,
					balance.Current,
					"expected current \"%d\" got  \"%d\"",
					test.resultBalance,
					balance.Current,
				)
				assert.Equal(
					t,
					test.resultWithdawn,
					balance.Withdrawn,
					"expected withdrawn \"%d\" got  \"%d\"",
					test.resultWithdawn,
					balance.Withdrawn,
				)
			}
		})
	}
}
