package web

type OrderCreateRequest struct {
	CustomerID  string             `json:"customer_id" validate:"required"`
	OrderDate   string             `json:"order_date" validate:"required,datetime=2006-01-02"`
	TotalAmount float64            `json:"total_amount" validate:"required,gte=0"`
	OrderItems  []OrderItemRequest `json:"order_items" validate:"required,dive"`
}

type OrderUpdateRequest struct {
	OrderID     string             `json:"order_id" validate:"required"`
	CustomerID  string             `json:"customer_id" validate:"required"`
	OrderDate   string             `json:"order_date" validate:"required,datetime=2006-01-02"`
	TotalAmount float64            `json:"total_amount" validate:"required,gte=0"`
	OrderItems  []OrderItemRequest `json:"order_items" validate:"required,dive"`
}

type OrderResponse struct {
	OrderID     string              `json:"order_id"`
	CustomerID  string              `json:"customer_id"`
	OrderDate   string              `json:"order_date"`
	TotalAmount float64             `json:"total_amount"`
	OrderItems  []OrderItemResponse `json:"order_items"`
}

type OrderItemRequest struct {
	ProductID  string  `json:"product_id" validate:"required"`
	Quantity   int     `json:"quantity" validate:"required,gte=1"`
	UnitPrice  float64 `json:"unit_price" validate:"required,gte=0"`
	TotalPrice float64 `json:"total_price" validate:"required,gte=0"`
}

type OrderItemResponse struct {
	ProductID  string  `json:"product_id"`
	Quantity   int     `json:"quantity"`
	UnitPrice  float64 `json:"unit_price"`
	TotalPrice float64 `json:"total_price"`
}
