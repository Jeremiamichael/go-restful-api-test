package repository

import (
	"context"
	"github.com/aronipurwanto/go-restful-api/model/domain"
)

type TaxRepository interface {
	Save(ctx context.Context, tax domain.Tax) (domain.Tax, error)
	Update(ctx context.Context, tax domain.Tax) (domain.Tax, error)
	Delete(ctx context.Context, tax domain.Tax) error
	FindById(ctx context.Context, taxId string) (domain.Tax, error)
	FindAll(ctx context.Context) ([]domain.Tax, error)
	FindByType(ctx context.Context, taxType string) ([]domain.Tax, error)
}
