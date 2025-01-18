package checker

import (
	"testing"

	"github.com/ndreyserg/gophermart/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestChecker(t *testing.T) {
	tests := []struct {
		name    string
		value   string
		isValid bool
	}{
		{
			name:    "Valid number 1",
			value:   "5062821234567892",
			isValid: true,
		},
		{
			name:    "Valid number 2",
			value:   "1234561239",
			isValid: true,
		},
		{
			name:    "unvalidnumber",
			value:   "324423",
			isValid: false,
		},
	}

	checker := NewService()
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := checker.Check(test.value)
			if test.isValid {
				assert.Nil(t, err)
			} else {
				assert.ErrorIs(t, err, model.ErrUncorrectOrederNumber)
			}
		})
	}
}
