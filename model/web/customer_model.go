package web

type CustomerCreateRequest struct {
	Name       string `json:"name" validate:"required,max=64,min=3"`
	Email      string `json:"email" validate:"required,email"`
	Phone      string `json:"phone" validate:"required,max=15"`
	Address    string `json:"address" validate:"required,max=255"`
	BirthDate  string `json:"birth_date" validate:"omitempty,datetime=2006-01-02"`
	LoyaltyPts int    `json:"loyalty_pts"` // ✅ Added to match domain.Customer
}

type CustomerUpdateRequest struct {
	CustomerID uint64 `json:"customer_id" validate:"required,gte=0"` // ✅ Renamed from Id
	Name       string `json:"name" validate:"required,max=64,min=3"`
	Email      string `json:"email" validate:"required,email"`
	Phone      string `json:"phone" validate:"required,max=15"`
	Address    string `json:"address" validate:"required,max=255"`
	BirthDate  string `json:"birth_date" validate:"omitempty,datetime=2006-01-02"`
	LoyaltyPts int    `json:"loyalty_pts"` // ✅ Added to match domain.Customer
}

type CustomerResponse struct {
	CustomerID uint64 `json:"customer_id"` // ✅ Renamed from Id
	Name       string `json:"name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Address    string `json:"address"`
	BirthDate  string `json:"birth_date"`
	LoyaltyPts int    `json:"loyalty_pts"` // ✅ Added to match domain.Customer
}
