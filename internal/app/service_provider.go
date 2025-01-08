package app

import (
	"database/sql"

	"github.com/ndreyserg/gophermart/internal/repository"
	accountRepository "github.com/ndreyserg/gophermart/internal/repository/account"
	orderRepository "github.com/ndreyserg/gophermart/internal/repository/order"
	userRepository "github.com/ndreyserg/gophermart/internal/repository/user"
	"github.com/ndreyserg/gophermart/internal/service"
	"github.com/ndreyserg/gophermart/internal/service/account"
	"github.com/ndreyserg/gophermart/internal/service/order"
	"github.com/ndreyserg/gophermart/internal/service/user"
)

type serviceProvider struct {
	userService       service.UserService
	userRepository    repository.UserRepository
	orderService      service.OrderService
	orderRepository   repository.OrderRepository
	accountService    service.AccountService
	accountRepository repository.AccountReposity
	db                *sql.DB
}

func newServiceProvider(db *sql.DB) *serviceProvider {
	return &serviceProvider{
		db: db,
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
		s.orderService = order.NewService(s.OrderRepository())
	}
	return s.orderService
}

func (s *serviceProvider) AccountRepository() repository.AccountReposity {
	if s.accountRepository == nil {
		s.accountRepository = accountRepository.NewRepository(s.db)
	}
	return s.accountRepository
}

func (s *serviceProvider) AccountService() service.AccountService {
	if s.accountService == nil {
		s.accountService = account.NewService(s.AccountRepository())
	}
	return s.accountService
}
