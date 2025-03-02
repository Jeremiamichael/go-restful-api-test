package web

type InventoryCreateRequest struct {
	ProductID    string `json:"product_id" validate:"required"`
	StockQty     int    `json:"stock_qty" validate:"required,gte=0"`
	RestockLevel int    `json:"restock_level" validate:"required,gte=0"`
	LastRestock  string `json:"last_restock" validate:"omitempty,datetime=2006-01-02"` // Optional, but must follow date format
}

type InventoryUpdateRequest struct {
	ProductID    string `json:"product_id" validate:"required"`
	StockQty     int    `json:"stock_qty" validate:"required,gte=0"`
	RestockLevel int    `json:"restock_level" validate:"required,gte=0"`
	LastRestock  string `json:"last_restock" validate:"omitempty,datetime=2006-01-02"`
}

type InventoryResponse struct {
	ProductID    string `json:"product_id"`
	StockQty     int    `json:"stock_qty"`
	RestockLevel int    `json:"restock_level"`
	LastRestock  string `json:"last_restock"`
}
