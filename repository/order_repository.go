package repository

import (
	"context"
	"github.com/aronipurwanto/go-restful-api/model/domain"
)

type OrderRepository interface {
	Save(ctx context.Context, order domain.Order) (domain.Order, error)
	Update(ctx context.Context, order domain.Order) (domain.Order, error)
	Delete(ctx context.Context, order domain.Order) error
	FindById(ctx context.Context, orderId string) (domain.Order, error)
	FindAll(ctx context.Context) ([]domain.Order, error)
}
