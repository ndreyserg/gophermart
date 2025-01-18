package app

import (
	"database/sql"

	"github.com/ndreyserg/gophermart/internal/repository"
	accountRepository "github.com/ndreyserg/gophermart/internal/repository/account"
	accrualRepository "github.com/ndreyserg/gophermart/internal/repository/accrual"
	orderRepository "github.com/ndreyserg/gophermart/internal/repository/order"
	userRepository "github.com/ndreyserg/gophermart/internal/repository/user"
	"github.com/ndreyserg/gophermart/internal/service"
	"github.com/ndreyserg/gophermart/internal/service/account"
	"github.com/ndreyserg/gophermart/internal/service/accrual"
	"github.com/ndreyserg/gophermart/internal/service/checker"
	"github.com/ndreyserg/gophermart/internal/service/order"
	"github.com/ndreyserg/gophermart/internal/service/session"
	"github.com/ndreyserg/gophermart/internal/service/user"
)

type serviceProvider struct {
	userService       service.UserService
	userRepository    repository.UserRepository
	orderService      service.OrderService
	orderRepository   repository.OrderRepository
	accountService    service.AccountService
	accountRepository repository.AccountReposity
	accrualRepository repository.AccrualReposity
	accrualService    service.AccrualService
	numChecker        service.CheckerSevice
	sessionService    service.SessionService
	db                *sql.DB
	accrualURI        string
	secret            string
}

func newServiceProvider(db *sql.DB, accrualURI string, secret string) *serviceProvider {
	return &serviceProvider{
		db:         db,
		accrualURI: accrualURI,
		secret:     secret,
	}
}

func (s *serviceProvider) UserService() service.UserService {
	if s.userService == nil {
		s.userService = user.NewService(s.UserRepository())
	}
	return s.userService
}

func (s *serviceProvider) UserRepository() repository.UserRepository {
	if s.userRepository == nil {
		s.userRepository = userRepository.NewRepository(s.db)
	}
	return s.userRepository
}

func (s *serviceProvider) OrderRepository() repository.OrderRepository {
	if s.orderRepository == nil {
		s.orderRepository = orderRepository.NewRepository(s.db)
	}
	return s.orderRepository
}

func (s *serviceProvider) OrderService() service.OrderService {
	if s.orderService == nil {
		s.orderService = order.NewService(
			s.OrderRepository(),
			s.AccrualService(),
			s.AccountService(),
			s.NumCheckerService(),
		)
	}
	return s.orderService
}

func (s *serviceProvider) AccountRepository() repository.AccountReposity {
	if s.accountRepository == nil {
		s.accountRepository = accountRepository.NewRepository(s.db)
	}
	return s.accountRepository
}

func (s *serviceProvider) AccrualRepository() repository.AccrualReposity {
	if s.accrualRepository == nil {
		s.accrualRepository = accrualRepository.NewRepository(s.accrualURI)
	}
	return s.accrualRepository
}

func (s *serviceProvider) AccountService() service.AccountService {
	if s.accountService == nil {
		s.accountService = account.NewService(s.AccountRepository(), s.NumCheckerService())
	}
	return s.accountService
}

func (s *serviceProvider) NumCheckerService() service.CheckerSevice {
	if s.numChecker == nil {
		s.numChecker = checker.NewService()
	}
	return s.numChecker
}

func (s *serviceProvider) AccrualService() service.AccrualService {
	if s.accrualService == nil {
		s.accrualService = accrual.NewService(
			s.AccrualRepository(),
			s.OrderRepository(),
			s.AccountService(),
		)
	}
	return s.accrualService
}

func (s *serviceProvider) SessionService() service.SessionService {
	if s.sessionService == nil {
		s.sessionService = session.NewService(s.secret)
	}
	return s.sessionService
}
