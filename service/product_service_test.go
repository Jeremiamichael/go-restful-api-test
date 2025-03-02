package service

import (
	"context"
	"errors"
	"github.com/aronipurwanto/go-restful-api/model/domain"
	"github.com/aronipurwanto/go-restful-api/model/web"
	"github.com/aronipurwanto/go-restful-api/repository/mocks"
	"github.com/go-playground/validator/v10"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockProductRepository(ctrl)
	mockValidator := validator.New()
	productService := NewProductService(mockRepo, mockValidator)

	tests := []struct {
		name      string
		input     web.ProductCreateRequest
		mock      func()
		expect    web.ProductResponse
		expectErr bool
	}{
		{
			name: "success",
			input: web.ProductCreateRequest{
				Name:        "Laptop",
				Description: "A high-end gaming laptop",
				Price:       1500.99,
				StockQty:    10,
				CategoryID:  1,
				SKU:         "LAPTOP123",
				TaxRate:     10.5,
			},
			mock: func() {
				mockRepo.EXPECT().Save(gomock.Any(), gomock.Any()).Return(domain.Product{
					ProductID:   1,
					Name:        "Laptop",
					Description: "A high-end gaming laptop",
					Price:       1500.99,
					StockQty:    10,
					CategoryId:  1,
					SKU:         "LAPTOP123",
					TaxRate:     10.5,
				}, nil)
			},
			expect: web.ProductResponse{
				Id:          1,
				Name:        "Laptop",
				Description: "A high-end gaming laptop",
				Price:       1500.99,
				StockQty:    10,
				CategoryID:  1,
				SKU:         "LAPTOP123",
				TaxRate:     10.5,
			},
			expectErr: false,
		},
		{
			name: "validation error",
			input: web.ProductCreateRequest{
				Name:        "",
				Description: "",
				Price:       -1,
				StockQty:    -1,
				CategoryID:  0,
				SKU:         "",
				TaxRate:     -1,
			},
			mock:      func() {},
			expect:    web.ProductResponse{},
			expectErr: true,
		},
		{
			name: "repository error",
			input: web.ProductCreateRequest{
				Name:        "Smartphone",
				Description: "A flagship smartphone",
				Price:       999.99,
				StockQty:    5,
				CategoryID:  2,
				SKU:         "PHONE123",
				TaxRate:     12.0,
			},
			mock: func() {
				mockRepo.EXPECT().Save(gomock.Any(), gomock.Any()).Return(domain.Product{}, errors.New("database error"))
			},
			expect:    web.ProductResponse{},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			resp, err := productService.Create(context.Background(), tt.input)
			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expect, resp)
			}
		})
	}
}
