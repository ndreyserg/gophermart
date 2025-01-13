package model

type Accrual struct {
	Status  OrderStatus `json:"status"`
	Order   string      `json:"order"`
	Accrual float64     `json:"accrual"`
}
