package user

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/ndreyserg/gophermart/internal/repository"
)

func NewService(userRepository repository.UserRepository) *service {
	return &service{
		userRepository: userRepository,
	}
}

type service struct {
	userRepository repository.UserRepository
}

func getPassHash(pass string) string {
	h := sha256.New()
	h.Write([]byte(pass))
	return hex.EncodeToString(h.Sum(nil))
}
