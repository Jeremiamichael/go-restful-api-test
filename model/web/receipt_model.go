package web

type ReceiptCreateRequest struct {
	OrderID     string  `json:"order_id" validate:"required"`
	PaymentID   string  `json:"payment_id" validate:"required"`
	ReceiptDate string  `json:"receipt_date" validate:"required,datetime=2006-01-02"`
	TotalAmount float64 `json:"total_amount" validate:"required,gte=0"`
	Taxes       float64 `json:"taxes" validate:"required,gte=0"`
	Discount    float64 `json:"discount" validate:"required,gte=0"`
	FinalAmount float64 `json:"final_amount" validate:"required,gte=0"`
}

type ReceiptUpdateRequest struct {
	ReceiptID   string  `json:"receipt_id" validate:"required"`
	OrderID     string  `json:"order_id" validate:"required"`
	PaymentID   string  `json:"payment_id" validate:"required"`
	ReceiptDate string  `json:"receipt_date" validate:"required,datetime=2006-01-02"`
	TotalAmount float64 `json:"total_amount" validate:"required,gte=0"`
	Taxes       float64 `json:"taxes" validate:"required,gte=0"`
	Discount    float64 `json:"discount" validate:"required,gte=0"`
	FinalAmount float64 `json:"final_amount" validate:"required,gte=0"`
}

type ReceiptResponse struct {
	ReceiptID   string  `json:"receipt_id"`
	OrderID     string  `json:"order_id"`
	PaymentID   string  `json:"payment_id"`
	ReceiptDate string  `json:"receipt_date"`
	TotalAmount float64 `json:"total_amount"`
	Taxes       float64 `json:"taxes"`
	Discount    float64 `json:"discount"`
	FinalAmount float64 `json:"final_amount"`
}
