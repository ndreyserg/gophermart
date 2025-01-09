package order

import (
	"regexp"
	"strconv"

	"github.com/ndreyserg/gophermart/internal/model"
)

func (s *service) CheckNumber(num string) error {
	match, _ := regexp.MatchString(`^[0-9]+$`, num)

	if !match {
		return model.ErrorUncorrectOrederNumber
	}

	sum := 0
	nDigits := len(num)
	parity := nDigits % 2

	for i := 0; i < len(num); i++ {
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
		return model.ErrorUncorrectOrederNumber
	}

	return nil
}
