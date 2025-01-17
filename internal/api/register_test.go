package api

import (
	"strings"
	"testing"

	"net/http"
	"net/http/httptest"

	"github.com/golang/mock/gomock"
	"github.com/ndreyserg/gophermart/internal/mocks"
	"github.com/ndreyserg/gophermart/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestApiRegister(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	userService := mocks.NewMockUserService(ctrl)
	orderService := mocks.NewMockOrderService(ctrl)
	accService := mocks.NewMockAccountService(ctrl)
	sessSevice := mocks.NewMockSessionService(ctrl)

	userService.EXPECT().Register(
		gomock.Any(), gomock.Eq("andrey"), gomock.Eq("secret")).Return(&model.User{ID: 1}, nil)
	sessSevice.EXPECT().Open(gomock.Eq(1), gomock.Any(), gomock.Any()).Return(nil)

	userService.EXPECT().Register(gomock.Any(), gomock.Eq("andrey"),
		gomock.Eq("secret")).Return(nil, model.ErrUserAllreadyExist)

	router := NewRouter(userService, orderService, accService, sessSevice)
	ts := httptest.NewServer(router)

	tests := []struct {
		name           string
		request        string
		body           string
		method         string
		wantStatusCode int
	}{
		{
			name:           "success register",
			request:        "/api/user/register",
			body:           `{"login": "andrey", "password": "secret"}`,
			wantStatusCode: http.StatusOK,
		},
		{
			name:           "login conflict",
			request:        "/api/user/register",
			body:           `{"login": "andrey", "password": "secret"}`,
			wantStatusCode: http.StatusConflict,
		},
		{
			name:           "bad request",
			request:        "/api/user/register",
			body:           `{"password": "secret"}`,
			wantStatusCode: http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req, err := http.NewRequest(
				http.MethodPost,
				ts.URL+test.request,
				strings.NewReader(test.body),
			)
			require.NoError(t, err)
			resp, err := ts.Client().Do(req)
			require.NoError(t, err)
			defer func() {
				_ = resp.Body.Close()
			}()
			assert.Equal(
				t,
				test.wantStatusCode,
				resp.StatusCode,
				"expected status code %d got %d",
				test.wantStatusCode,
				resp.StatusCode,
			)
		})
	}
}
