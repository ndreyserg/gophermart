package api

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/ndreyserg/gophermart/internal/mocks"
	"github.com/ndreyserg/gophermart/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type deps struct {
	userService  *mocks.MockUserService
	orderService *mocks.MockOrderService
	accService   *mocks.MockAccountService
	sessService  *mocks.MockSessionService
}

type tCase struct {
	name            string
	request         string
	body            string
	method          string
	wantStatusCode  int
	hasResponseBody bool
	responseBody    string
}

func initDeps(t *testing.T, ctrl *gomock.Controller) *deps {
	t.Helper()
	return &deps{
		userService:  mocks.NewMockUserService(ctrl),
		orderService: mocks.NewMockOrderService(ctrl),
		accService:   mocks.NewMockAccountService(ctrl),
		sessService:  mocks.NewMockSessionService(ctrl),
	}
}

func runTest(t *testing.T, tests []tCase, ts *httptest.Server) {
	t.Helper()
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req, err := http.NewRequest(
				test.method,
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
			if test.hasResponseBody {
				respBody, err := io.ReadAll(resp.Body)
				require.NoError(t, err)
				assert.Equal(
					t,
					test.responseBody,
					strings.Trim(string(respBody), "\n"),
					"expected body \"%s\" got  \"%s\"",
					test.responseBody,
					respBody,
				)
			}
		})
	}
}

func TestApiRegister(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	d := initDeps(t, ctrl)
	d.userService.EXPECT().Register(
		gomock.Any(), gomock.Eq("andrey"), gomock.Eq("secret")).Return(&model.User{ID: 1}, nil)
	d.sessService.EXPECT().Open(gomock.Eq(1), gomock.Any(), gomock.Any()).Return(nil)

	d.userService.EXPECT().Register(gomock.Any(), gomock.Eq("andrey"),
		gomock.Eq("secret")).Return(nil, model.ErrUserAllreadyExist)
	router := NewRouter(d.userService, d.orderService, d.accService, d.sessService)
	ts := httptest.NewServer(router)

	tests := []tCase{
		{
			name:           "success register",
			method:         http.MethodPost,
			request:        "/api/user/register",
			body:           `{"login": "andrey", "password": "secret"}`,
			wantStatusCode: http.StatusOK,
		},
		{
			name:           "login conflict",
			method:         http.MethodPost,
			request:        "/api/user/register",
			body:           `{"login": "andrey", "password": "secret"}`,
			wantStatusCode: http.StatusConflict,
		},
		{
			name:           "bad request",
			method:         http.MethodPost,
			request:        "/api/user/register",
			body:           `{"password": "secret"}`,
			wantStatusCode: http.StatusBadRequest,
		},
	}
	runTest(t, tests, ts)
}
func TestApiLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	d := initDeps(t, ctrl)

	d.userService.EXPECT().Login(gomock.Any(), gomock.Eq("andrey"),
		gomock.Eq("secret")).Return(&model.User{ID: 1}, nil)
	d.sessService.EXPECT().Open(gomock.Eq(1), gomock.Any(), gomock.Any()).Return(nil)
	d.userService.EXPECT().Login(gomock.Any(), gomock.Eq("andrey"),
		gomock.Eq("bad_secret")).Return(nil, model.ErrUserNotFound)

	router := NewRouter(d.userService, d.orderService, d.accService, d.sessService)
	ts := httptest.NewServer(router)

	tests := []tCase{
		{
			name:           "success login",
			request:        "/api/user/login",
			body:           `{"login": "andrey", "password": "secret"}`,
			method:         http.MethodPost,
			wantStatusCode: http.StatusOK,
		},
		{
			name:           "bad pass",
			request:        "/api/user/login",
			body:           `{"login": "andrey", "password": "bad_secret"}`,
			method:         http.MethodPost,
			wantStatusCode: http.StatusUnauthorized,
		},
	}
	runTest(t, tests, ts)
}

func TestApiCreateOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	d := initDeps(t, ctrl)

	d.sessService.EXPECT().GetUserID(gomock.Any()).Return(0, errors.New(""))

	d.sessService.EXPECT().GetUserID(gomock.Any()).Return(1, nil)
	d.orderService.EXPECT().Create(gomock.Any(), gomock.Eq("5062821234567892"),
		gomock.Eq(1)).Return(&model.Order{}, nil)

	d.sessService.EXPECT().GetUserID(gomock.Any()).Return(1, nil)
	d.orderService.EXPECT().Create(gomock.Any(), gomock.Eq("5062821234567892"),
		gomock.Eq(1)).Return(nil, model.ErrOrderAllreadyExistOnUser)

	d.sessService.EXPECT().GetUserID(gomock.Any()).Return(1, nil)
	d.orderService.EXPECT().Create(gomock.Any(), gomock.Eq("5062821234567892"),
		gomock.Eq(1)).Return(nil, model.ErrOrderAllreadyExist)

	d.sessService.EXPECT().GetUserID(gomock.Any()).Return(1, nil)
	d.orderService.EXPECT().Create(gomock.Any(), gomock.Eq("111"),
		gomock.Eq(1)).Return(nil, model.ErrUncorrectOrederNumber)

	router := NewRouter(d.userService, d.orderService, d.accService, d.sessService)
	ts := httptest.NewServer(router)

	tests := []tCase{
		{
			name:           "not authorized",
			request:        "/api/user/orders",
			body:           "5062821234567892",
			method:         http.MethodPost,
			wantStatusCode: http.StatusUnauthorized,
		},
		{
			name:           "success create",
			request:        "/api/user/orders",
			body:           "5062821234567892",
			method:         http.MethodPost,
			wantStatusCode: http.StatusAccepted,
		},
		{
			name:           "order already exist",
			request:        "/api/user/orders",
			body:           "5062821234567892",
			method:         http.MethodPost,
			wantStatusCode: http.StatusOK,
		},
		{
			name:           "order number conflict",
			request:        "/api/user/orders",
			body:           "5062821234567892",
			method:         http.MethodPost,
			wantStatusCode: http.StatusConflict,
		},
		{
			name:           "order number unvalid",
			request:        "/api/user/orders",
			body:           "111",
			method:         http.MethodPost,
			wantStatusCode: http.StatusUnprocessableEntity,
		},
	}
	runTest(t, tests, ts)
}

func TestApiGetUserOrders(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	d := initDeps(t, ctrl)

	router := NewRouter(d.userService, d.orderService, d.accService, d.sessService)
	ts := httptest.NewServer(router)

	d.sessService.EXPECT().GetUserID(gomock.Any()).Return(0, errors.New(""))

	d.sessService.EXPECT().GetUserID(gomock.Any()).Return(1, nil)
	d.orderService.EXPECT().GetByUserID(gomock.Any(), gomock.Eq(1)).Return(
		[]*model.Order{}, nil)

	d.sessService.EXPECT().GetUserID(gomock.Any()).Return(1, nil)
	d.orderService.EXPECT().GetByUserID(gomock.Any(), gomock.Eq(1)).Return(
		[]*model.Order{{Status: "NEW", Number: "123", UploadedAt: "2025-01-15T02:40:03+00:00", Accrual: 0}}, nil)

	tests := []tCase{
		{
			name:           "not authorized",
			request:        "/api/user/orders",
			body:           "",
			method:         http.MethodGet,
			wantStatusCode: http.StatusUnauthorized,
		},
		{
			name:           "no orders",
			request:        "/api/user/orders",
			body:           "",
			method:         http.MethodGet,
			wantStatusCode: http.StatusNoContent,
		},

		{
			name:            "success",
			request:         "/api/user/orders",
			body:            "",
			method:          http.MethodGet,
			wantStatusCode:  http.StatusOK,
			hasResponseBody: true,
			responseBody:    `[{"status":"NEW","number":"123","uploaded_at":"2025-01-15T02:40:03+00:00"}]`,
		},
	}
	runTest(t, tests, ts)
}

func TestApiWithdraw(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	d := initDeps(t, ctrl)

	router := NewRouter(d.userService, d.orderService, d.accService, d.sessService)
	ts := httptest.NewServer(router)

	d.sessService.EXPECT().GetUserID(gomock.Any()).Return(0, errors.New(""))
	d.sessService.EXPECT().GetUserID(gomock.Any()).Return(1, nil)
	d.sessService.EXPECT().GetUserID(gomock.Any()).Return(1, nil)
	d.sessService.EXPECT().GetUserID(gomock.Any()).Return(1, nil)
	d.accService.EXPECT().Withdraw(
		gomock.Any(), gomock.Eq(1),
		gomock.Eq("123"), gomock.Eq(float64(1))).Return(model.ErrUncorrectOrederNumber)

	d.sessService.EXPECT().GetUserID(gomock.Any()).Return(1, nil)
	d.accService.EXPECT().Withdraw(
		gomock.Any(), gomock.Eq(1),
		gomock.Eq("5062821234567892"), gomock.Eq(float64(1))).Return(model.ErrAccountNegativeBalance)

	d.sessService.EXPECT().GetUserID(gomock.Any()).Return(1, nil)
	d.accService.EXPECT().Withdraw(
		gomock.Any(), gomock.Eq(1),
		gomock.Eq("5062821234567892"), gomock.Eq(float64(1))).Return(nil)

	tests := []tCase{
		{
			name:           "not authorized",
			request:        "/api/user/balance/withdraw",
			body:           "",
			method:         http.MethodPost,
			wantStatusCode: http.StatusUnauthorized,
		},
		{
			name:           "empty request",
			request:        "/api/user/balance/withdraw",
			body:           "",
			method:         http.MethodPost,
			wantStatusCode: http.StatusBadRequest,
		},
		{
			name:           "zero sum",
			request:        "/api/user/balance/withdraw",
			body:           `{"order": "5062821234567892", "sum": 0}`,
			method:         http.MethodPost,
			wantStatusCode: http.StatusUnprocessableEntity,
		},
		{
			name:           "unvalid order",
			request:        "/api/user/balance/withdraw",
			body:           `{"order": "123", "sum": 1}`,
			method:         http.MethodPost,
			wantStatusCode: http.StatusUnprocessableEntity,
		},
		{
			name:           "negative balance",
			request:        "/api/user/balance/withdraw",
			body:           `{"order": "5062821234567892", "sum": 1}`,
			method:         http.MethodPost,
			wantStatusCode: http.StatusPaymentRequired,
		},
		{
			name:           "success",
			request:        "/api/user/balance/withdraw",
			body:           `{"order": "5062821234567892", "sum": 1}`,
			method:         http.MethodPost,
			wantStatusCode: http.StatusOK,
		},
	}
	runTest(t, tests, ts)
}

func TestGetBalance(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	d := initDeps(t, ctrl)

	d.sessService.EXPECT().GetUserID(gomock.Any()).Return(0, errors.New(""))
	d.sessService.EXPECT().GetUserID(gomock.Any()).Return(1, nil)
	d.accService.EXPECT().GetBalance(gomock.Any(), gomock.Eq(1)).Return(
		&model.AccountBalance{
			Current:   22,
			Withdrawn: 123,
		},
		nil,
	)

	router := NewRouter(d.userService, d.orderService, d.accService, d.sessService)
	ts := httptest.NewServer(router)
	tests := []tCase{
		{
			name:           "not authorized",
			request:        "/api/user/balance",
			body:           "",
			method:         http.MethodGet,
			wantStatusCode: http.StatusUnauthorized,
		},
		{
			name:            "success",
			request:         "/api/user/balance",
			body:            "",
			method:          http.MethodGet,
			wantStatusCode:  http.StatusOK,
			hasResponseBody: true,
			responseBody:    `{"current":22,"withdrawn":123}`,
		},
	}
	runTest(t, tests, ts)
}

func TestGetWithdrawals(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	d := initDeps(t, ctrl)

	d.sessService.EXPECT().GetUserID(gomock.Any()).Return(0, errors.New(""))

	d.sessService.EXPECT().GetUserID(gomock.Any()).Return(1, nil)
	d.accService.EXPECT().GetWithdrawals(gomock.Any(), gomock.Eq(1)).Return(
		[]*model.AccountWithdrawals{},
		nil,
	)

	d.sessService.EXPECT().GetUserID(gomock.Any()).Return(1, nil)
	d.accService.EXPECT().GetWithdrawals(gomock.Any(), gomock.Eq(1)).Return(
		[]*model.AccountWithdrawals{
			{OrderNumber: "123", AccountID: 1, Sum: 100, ID: 1, ProcessedAt: "2025-01-15T02:40:03+00:00"},
		},
		nil,
	)

	router := NewRouter(d.userService, d.orderService, d.accService, d.sessService)
	ts := httptest.NewServer(router)
	tests := []tCase{
		{
			name:           "not authorized",
			request:        "/api/user/withdrawals",
			body:           "",
			method:         http.MethodGet,
			wantStatusCode: http.StatusUnauthorized,
		},
		{
			name:           "empty withdrawals",
			request:        "/api/user/withdrawals",
			body:           "",
			method:         http.MethodGet,
			wantStatusCode: http.StatusNoContent,
		},
		{
			name:            "success",
			request:         "/api/user/withdrawals",
			body:            "",
			method:          http.MethodGet,
			wantStatusCode:  http.StatusOK,
			hasResponseBody: true,
			responseBody:    `[{"order":"123","processed_at":"2025-01-15T02:40:03+00:00","sum":100}]`,
		},
	}
	runTest(t, tests, ts)
}
