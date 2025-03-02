package service

import (
	"context"
	"github.com/aronipurwanto/go-restful-api/model/web"
)

type DiscountService interface {
	Create(ctx context.Context, request web.DiscountCreateRequest) (web.DiscountResponse, error)
	Update(ctx context.Context, request web.DiscountUpdateRequest) (web.DiscountResponse, error)
	Delete(ctx context.Context, discountId string) error
	FindById(ctx context.Context, discountId string) (web.DiscountResponse, error)
	FindAll(ctx context.Context) ([]web.DiscountResponse, error)
}
