package model

type Account struct {
	ID      int
	UserID  int
	Balance float64
}

type AccountBalance struct {
	Current   float64 `json:"current"`
	Withdrawn float64 `json:"withdrawn"`
}

type AccountWithdrawals struct {
	ID          int     `json:"-"`
	Sum         float64 `json:"sum"`
	AccountID   int     `json:"-"`
	OrderNumber string  `json:"order"`
	ProcessedAt string	`json:"processed_at"`
}
