package web

type PaymentCreateRequest struct {
	OrderID     string  `json:"order_id" validate:"required"`
	Amount      float64 `json:"amount" validate:"required,gte=0"`
	PaymentType string  `json:"payment_type" validate:"required,max=32"` // e.g., Cash, Card, Online
	PaymentDate string  `json:"payment_date" validate:"required,datetime=2006-01-02"`
	Status      string  `json:"status" validate:"required,oneof=Completed Pending Failed"`
}

type PaymentUpdateRequest struct {
	PaymentID   string  `json:"payment_id" validate:"required"`
	OrderID     string  `json:"order_id" validate:"required"`
	Amount      float64 `json:"amount" validate:"required,gte=0"`
	PaymentType string  `json:"payment_type" validate:"required,max=32"`
	PaymentDate string  `json:"payment_date" validate:"required,datetime=2006-01-02"`
	Status      string  `json:"status" validate:"required,oneof=Completed Pending Failed"`
}

type PaymentResponse struct {
	PaymentID   string  `json:"payment_id"`
	OrderID     string  `json:"order_id"`
	Amount      float64 `json:"amount"`
	PaymentType string  `json:"payment_type"`
	PaymentDate string  `json:"payment_date"`
	Status      string  `json:"status"`
}
