package web

type TaxCreateRequest struct {
	TaxRate     float64 `json:"tax_rate" validate:"required,gte=0,lte=100"` // Percentage between 0-100
	TaxType     string  `json:"tax_type" validate:"required,max=32"`        // e.g., Sales Tax, VAT
	Description string  `json:"description" validate:"required,max=255"`
}

type TaxUpdateRequest struct {
	TaxID       string  `json:"tax_id" validate:"required"`
	TaxRate     float64 `json:"tax_rate" validate:"required,gte=0,lte=100"`
	TaxType     string  `json:"tax_type" validate:"required,max=32"`
	Description string  `json:"description" validate:"required,max=255"`
}

type TaxResponse struct {
	TaxID       string  `json:"tax_id"`
	TaxRate     float64 `json:"tax_rate"`
	TaxType     string  `json:"tax_type"`
	Description string  `json:"description"`
}
