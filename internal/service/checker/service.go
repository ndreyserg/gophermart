package checker

import (
	"regexp"
	"strconv"

	"github.com/ndreyserg/gophermart/internal/model"
	def "github.com/ndreyserg/gophermart/internal/service"
)

var _ def.CheckerSevice = (*service)(nil)

func NewService() *service {
	return &service{}
}

type service struct{}

func (s *service) Check(num string) error {
	match, _ := regexp.MatchString(`^[0-9]+$`, num)

	if !match {
		return model.ErrUncorrectOrederNumber
	}

	sum := 0
	nDigits := len(num)
	parity := nDigits % 2

	for i := range num {
		digit, _ := strconv.Atoi(string(num[i]))

		if i%2 == parity {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
	}

	if sum%10 != 0 {
		return model.ErrUncorrectOrederNumber
	}

	return nil
}
