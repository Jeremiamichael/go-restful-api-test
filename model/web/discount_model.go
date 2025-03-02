package web

type DiscountCreateRequest struct {
	Description string  `json:"description" validate:"required,max=255"`
	DiscountPct float64 `json:"discount_pct" validate:"required,gte=0,lte=100"` // Ensure percentage is between 0-100
	ValidFrom   string  `json:"valid_from" validate:"required,datetime=2006-01-02"`
	ValidUntil  string  `json:"valid_until" validate:"required,datetime=2006-01-02,gtfield=ValidFrom"` // Must be after ValidFrom
}

type DiscountUpdateRequest struct {
	DiscountID  string  `json:"discount_id" validate:"required"`
	Description string  `json:"description" validate:"required,max=255"`
	DiscountPct float64 `json:"discount_pct" validate:"required,gte=0,lte=100"`
	ValidFrom   string  `json:"valid_from" validate:"required,datetime=2006-01-02"`
	ValidUntil  string  `json:"valid_until" validate:"required,datetime=2006-01-02,gtfield=ValidFrom"`
}

type DiscountResponse struct {
	DiscountID  string  `json:"discount_id"`
	Description string  `json:"description"`
	DiscountPct float64 `json:"discount_pct"`
	ValidFrom   string  `json:"valid_from"`
	ValidUntil  string  `json:"valid_until"`
}
