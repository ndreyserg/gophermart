package app

import (
	"database/sql"

	"github.com/ndreyserg/gophermart/internal/repository"
	userRepository "github.com/ndreyserg/gophermart/internal/repository/user"
	"github.com/ndreyserg/gophermart/internal/service"
	"github.com/ndreyserg/gophermart/internal/service/user"
)

type serviceProvider struct {
	userService    service.UserService
	userRepository repository.UserRepository
	db             *sql.DB
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
