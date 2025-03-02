package repository

import (
	"context"
	"github.com/aronipurwanto/go-restful-api/model/domain"
)

type PaymentRepository interface {
	Save(ctx context.Context, payment domain.Payment) (domain.Payment, error)
	Update(ctx context.Context, payment domain.Payment) (domain.Payment, error)
	Delete(ctx context.Context, payment domain.Payment) error
	FindById(ctx context.Context, paymentId string) (domain.Payment, error)
	FindAll(ctx context.Context) ([]domain.Payment, error)
	FindByOrderId(ctx context.Context, orderId string) ([]domain.Payment, error)
}
