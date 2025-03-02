package repository

import (
	"context"
	"github.com/aronipurwanto/go-restful-api/model/domain"
)

type ReceiptRepository interface {
	Save(ctx context.Context, receipt domain.Receipt) (domain.Receipt, error)
	Update(ctx context.Context, receipt domain.Receipt) (domain.Receipt, error)
	Delete(ctx context.Context, receipt domain.Receipt) error
	FindById(ctx context.Context, receiptId string) (domain.Receipt, error)
	FindAll(ctx context.Context) ([]domain.Receipt, error)
	FindByOrderId(ctx context.Context, orderId string) ([]domain.Receipt, error)
}
