package repository

import (
	"context"
	"errors"
	"testing"

	"github.com/aronipurwanto/go-restful-api/model/domain"
	"github.com/aronipurwanto/go-restful-api/repository/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCustomerRepository(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockCustomerRepository(ctrl)
	ctx := context.Background()

	tests := []struct {
		name      string
		mock      func()
		method    func() (interface{}, error)
		expect    interface{}
		expectErr bool
	}{
		{
			name: "Save Success",
			mock: func() {
				customer := domain.Customer{
					CustomerID: 1,
					Name:       "John Doe",
					Email:      "john@example.com",
					Phone:      "1234567890",
					Address:    "123 Street",
					LoyaltyPts: 100,
				}
				repo.EXPECT().Save(ctx, customer).Return(customer, nil)
			},
			method: func() (interface{}, error) {
				return repo.Save(ctx, domain.Customer{
					CustomerID: 1,
					Name:       "John Doe",
					Email:      "john@example.com",
					Phone:      "1234567890",
					Address:    "123 Street",
					LoyaltyPts: 100,
				})
			},
			expect: domain.Customer{
				CustomerID: 1,
				Name:       "John Doe",
				Email:      "john@example.com",
				Phone:      "1234567890",
				Address:    "123 Street",
				LoyaltyPts: 100,
			},
			expectErr: false,
		},
		{
			name: "Save Failure",
			mock: func() {
				repo.EXPECT().Save(ctx, gomock.Any()).Return(domain.Customer{}, errors.New("error saving"))
			},
			method: func() (interface{}, error) {
				return repo.Save(ctx, domain.Customer{Name: "Invalid"})
			},
			expect:    domain.Customer{},
			expectErr: true,
		},
		{
			name: "Update Success",
			mock: func() {
				customer := domain.Customer{
					CustomerID: 1,
					Name:       "Jane Doe",
					Email:      "jane@example.com",
					Phone:      "0987654321",
					Address:    "456 Avenue",
					LoyaltyPts: 200,
				}
				repo.EXPECT().Update(ctx, customer).Return(customer, nil)
			},
			method: func() (interface{}, error) {
				return repo.Update(ctx, domain.Customer{
					CustomerID: 1,
					Name:       "Jane Doe",
					Email:      "jane@example.com",
					Phone:      "0987654321",
					Address:    "456 Avenue",
					LoyaltyPts: 200,
				})
			},
			expect: domain.Customer{
				CustomerID: 1,
				Name:       "Jane Doe",
				Email:      "jane@example.com",
				Phone:      "0987654321",
				Address:    "456 Avenue",
				LoyaltyPts: 200,
			},
			expectErr: false,
		},
		{
			name: "FindById Success",
			mock: func() {
				repo.EXPECT().FindById(ctx, uint64(1)).Return(domain.Customer{
					CustomerID: 1,
					Name:       "John Doe",
					Email:      "john@example.com",
					Phone:      "1234567890",
					Address:    "123 Street",
					LoyaltyPts: 100,
				}, nil)
			},
			method: func() (interface{}, error) {
				return repo.FindById(ctx, 1)
			},
			expect: domain.Customer{
				CustomerID: 1,
				Name:       "John Doe",
				Email:      "john@example.com",
				Phone:      "1234567890",
				Address:    "123 Street",
				LoyaltyPts: 100,
			},
			expectErr: false,
		},
		{
			name: "FindById Not Found",
			mock: func() {
				repo.EXPECT().FindById(ctx, uint64(999)).Return(domain.Customer{}, errors.New("not found"))
			},
			method: func() (interface{}, error) {
				return repo.FindById(ctx, 999)
			},
			expect:    domain.Customer{},
			expectErr: true,
		},
		{
			name: "FindAll Success",
			mock: func() {
				repo.EXPECT().FindAll(ctx).Return([]domain.Customer{
					{
						CustomerID: 1,
						Name:       "John Doe",
						Email:      "john@example.com",
						Phone:      "1234567890",
						Address:    "123 Street",
						LoyaltyPts: 100,
					},
				}, nil)
			},
			method: func() (interface{}, error) {
				return repo.FindAll(ctx)
			},
			expect: []domain.Customer{
				{
					CustomerID: 1,
					Name:       "John Doe",
					Email:      "john@example.com",
					Phone:      "1234567890",
					Address:    "123 Street",
					LoyaltyPts: 100,
				},
			},
			expectErr: false,
		},
		{
			name: "Delete Success",
			mock: func() {
				repo.EXPECT().Delete(ctx, domain.Customer{CustomerID: 1}).Return(nil)
			},
			method: func() (interface{}, error) {
				return nil, repo.Delete(ctx, domain.Customer{CustomerID: 1})
			},
			expect:    nil,
			expectErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			result, err := tt.method()

			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expect, result)
			}
		})
	}
}
