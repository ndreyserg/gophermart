package model

type OrderStatus string

const (
	OrderStatusNew        OrderStatus = "NEW"
	OrderStatusProcessed  OrderStatus = "PROCESSED"
	OrderStatusProcessing OrderStatus = "PROCESSING"
	OrderStatusInvalid    OrderStatus = "INVALID"
)

type Order struct {
	Status     OrderStatus `json:"status"`
	Number     string      `json:"number"`
	UploadedAt string      `json:"uploaded_at"`
	ID         int         `json:"-"`
	UserID     int         `json:"-"`
	Accrual    float64     `json:"accrual,omitempty"`
}
