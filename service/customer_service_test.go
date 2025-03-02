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

func TestCreateCustomer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockCustomerRepository(ctrl)
	mockValidator := validator.New()
	customerService := NewCustomerService(mockRepo, mockValidator)

	tests := []struct {
		name      string
		input     web.CustomerCreateRequest
		mock      func()
		expect    web.CustomerResponse
		expectErr bool
	}{
		{
			name:  "success",
			input: web.CustomerCreateRequest{Name: "John Doe", Email: "john.doe@example.com", Phone: "123456789", Address: "123 Street", LoyaltyPts: 100},
			mock: func() {
				mockRepo.EXPECT().Save(gomock.Any(), gomock.Any()).Return(domain.Customer{CustomerID: 1, Name: "John Doe", Email: "john.doe@example.com", Phone: "123456789", Address: "123 Street", LoyaltyPts: 100}, nil)
			},
			expect:    web.CustomerResponse{CustomerID: 1, Name: "John Doe", Email: "john.doe@example.com", Phone: "123456789", Address: "123 Street", LoyaltyPts: 100},
			expectErr: false,
		},
		{
			name:      "validation error",
			input:     web.CustomerCreateRequest{Name: "", Email: "", Phone: "", Address: "", LoyaltyPts: 0},
			mock:      func() {},
			expect:    web.CustomerResponse{},
			expectErr: true,
		},
		{
			name:  "repository error",
			input: web.CustomerCreateRequest{Name: "Jane Doe", Email: "jane.doe@example.com", Phone: "987654321", Address: "456 Avenue", LoyaltyPts: 50},
			mock: func() {
				mockRepo.EXPECT().Save(gomock.Any(), gomock.Any()).Return(domain.Customer{}, errors.New("database error"))
			},
			expect:    web.CustomerResponse{},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			resp, err := customerService.Create(context.Background(), tt.input)
			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expect, resp)
			}
		})
	}
}
