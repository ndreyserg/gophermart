package user

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
	repo *mocks.MockUserRepository
}

type tCase struct {
	name      string
	login     string
	pass      string
	prepare   func(d *deps)
	expectErr string
}

func initDeps(ctrl *gomock.Controller) *deps {
	return &deps{
		repo: mocks.NewMockUserRepository(ctrl),
	}
}

//nolint:dupl //описание посдовательности вызовов
func TestRegister(t *testing.T) {
	tests := []tCase{
		{
			name:  "user allready exist",
			login: "andrey",
			pass:  "secret",
			prepare: func(d *deps) {
				d.repo.EXPECT().Create(
					gomock.Any(),
					gomock.Eq("andrey"),
					gomock.Eq(getPassHash("secret")),
				).Return(nil, errors.New("repo error"))
			},
			expectErr: "register user service error: repo error",
		},
		{
			name:  "success",
			login: "andrey",
			pass:  "secret",
			prepare: func(d *deps) {
				d.repo.EXPECT().Create(
					gomock.Any(),
					gomock.Eq("andrey"),
					gomock.Eq(getPassHash("secret")),
				).Return(&model.User{}, nil)
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
			s := NewService(d.repo)

			_, err := s.Register(context.Background(), test.login, test.pass)

			if test.expectErr != "" {
				assert.ErrorContains(t, err, test.expectErr)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

//nolint:dupl //описание посдовательности вызовов
func TestLogin(t *testing.T) {
	tests := []tCase{
		{
			name:  "repo error",
			login: "andrey",
			pass:  "secret",
			prepare: func(d *deps) {
				d.repo.EXPECT().FindByLoginAndPassHash(
					gomock.Any(),
					gomock.Eq("andrey"),
					gomock.Eq(getPassHash("secret")),
				).Return(nil, errors.New("repo error"))
			},
			expectErr: "login serivice error: repo error",
		},
		{
			name:  "success",
			login: "andrey",
			pass:  "secret",
			prepare: func(d *deps) {
				d.repo.EXPECT().FindByLoginAndPassHash(
					gomock.Any(),
					gomock.Eq("andrey"),
					gomock.Eq(getPassHash("secret")),
				).Return(&model.User{}, nil)
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
			s := NewService(d.repo)

			_, err := s.Login(context.Background(), test.login, test.pass)

			if test.expectErr != "" {
				assert.ErrorContains(t, err, test.expectErr)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
