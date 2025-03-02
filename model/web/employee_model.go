package web

type EmployeeCreateRequest struct {
	Name      string `json:"name" validate:"required,max=64,min=3"`
	Role      string `json:"role" validate:"required,max=32"` // e.g., Cashier, Manager
	Email     string `json:"email" validate:"required,email"`
	Phone     string `json:"phone" validate:"required,max=15"`
	DateHired string `json:"date_hired" validate:"required,datetime=2006-01-02"`
}

type EmployeeUpdateRequest struct {
	EmployeeID uint64 `json:"employee_id" validate:"required,gte=0"`
	Name       string `json:"name" validate:"required,max=64,min=3"`
	Role       string `json:"role" validate:"required,max=32"`
	Email      string `json:"email" validate:"required,email"`
	Phone      string `json:"phone" validate:"required,max=15"`
	DateHired  string `json:"date_hired" validate:"required,datetime=2006-01-02"`
}

type EmployeeResponse struct {
	EmployeeID uint64 `json:"employee_id"`
	Name       string `json:"name"`
	Role       string `json:"role"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	DateHired  string `json:"date_hired"`
}
