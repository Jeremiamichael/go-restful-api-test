package helper

import (
	"github.com/aronipurwanto/go-restful-api/model/domain"
	"github.com/aronipurwanto/go-restful-api/model/web"
)

// Convert single Category domain model to CategoryResponse
func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

// Convert multiple Category domain models to a slice of CategoryResponse
func ToCategoryResponses(categories []domain.Category) []web.CategoryResponse {
	var categoryResponses []web.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToCategoryResponse(category))
	}
	return categoryResponses
}

// Convert single Customer domain model to CustomerResponse
func ToCustomerResponse(customer domain.Customer) web.CustomerResponse {
	return web.CustomerResponse{
		CustomerID: customer.CustomerID,
		Name:       customer.Name,
		Email:      customer.Email,
		Phone:      customer.Phone,
		Address:    customer.Address,
		LoyaltyPts: customer.LoyaltyPts,
	}
}

// Convert multiple Customer domain models to a slice of CustomerResponse
func ToCustomerResponses(customers []domain.Customer) []web.CustomerResponse {
	var responses []web.CustomerResponse
	for _, customer := range customers {
		responses = append(responses, ToCustomerResponse(customer))
	}
	return responses
}

// Convert single Discount domain model to DiscountResponse
func ToDiscountResponse(discount domain.Discount) web.DiscountResponse {
	return web.DiscountResponse{
		DiscountID:  discount.DiscountID,
		Description: discount.Description,
		DiscountPct: discount.DiscountPct,
		ValidFrom:   discount.ValidFrom,
		ValidUntil:  discount.ValidUntil,
	}
}

// Convert multiple Discount domain models to a slice of DiscountResponse
func ToDiscountResponses(discounts []domain.Discount) []web.DiscountResponse {
	var responses []web.DiscountResponse
	for _, discount := range discounts {
		responses = append(responses, ToDiscountResponse(discount))
	}
	return responses
}

// Convert single Employee domain model to EmployeeResponse
func ToEmployeeResponse(employee domain.Employee) web.EmployeeResponse {
	return web.EmployeeResponse{
		EmployeeID: employee.EmployeeID,
		Name:       employee.Name,
		Role:       employee.Role,
		Email:      employee.Email,
		Phone:      employee.Phone,
		DateHired:  employee.DateHired,
	}
}

// Convert multiple Employee domain models to a slice of EmployeeResponse
func ToEmployeeResponses(employees []domain.Employee) []web.EmployeeResponse {
	var employeeResponses []web.EmployeeResponse
	for _, employee := range employees {
		employeeResponses = append(employeeResponses, ToEmployeeResponse(employee))
	}
	return employeeResponses
}

func ToProductResponse(product domain.Product) web.ProductResponse {
	return web.ProductResponse{
		Id:          product.ProductID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		StockQty:    product.StockQty,
		CategoryID:  int(product.CategoryId), // Converting uint64 to int for consistency with web models
		SKU:         product.SKU,
		TaxRate:     product.TaxRate,
	}
}

// Convert a slice of domain.Product to a slice of web.ProductResponse
func ToProductResponses(products []domain.Product) []web.ProductResponse {
	var responses []web.ProductResponse
	for _, product := range products {
		responses = append(responses, ToProductResponse(product))
	}
	return responses
}
