package accrual

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ndreyserg/gophermart/internal/model"
)

func (r repository) Get(orderNumber string) (*model.Accrual, error) {
	uri := r.addr + "/api/orders/" + orderNumber
	resp, err := http.Get(uri)

	if err != nil {
		return nil, fmt.Errorf("error accrual system request %s: %w", uri, err)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode == http.StatusOK {
		accrual := model.Accrual{}
		decoder := json.NewDecoder(resp.Body)
		err := decoder.Decode(&accrual)
		if err != nil {
			return nil, fmt.Errorf("accrual parse error %s: %w", uri, err)
		}
		return &accrual, nil
	}

	if resp.StatusCode == http.StatusNoContent {
		return nil, fmt.Errorf("unknown order number %s", orderNumber)
	}

	if resp.StatusCode == http.StatusTooManyRequests {
		return nil, model.ErrAccrualSystemTooManyRequests
	}

	return nil, fmt.Errorf("accrual system error request %s status %d", uri, resp.StatusCode)
}
