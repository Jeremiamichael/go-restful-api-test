package repository

import (
	"context"
	"github.com/aronipurwanto/go-restful-api/model/domain"
)

type InventoryRepository interface {
	Save(ctx context.Context, inventory domain.Inventory) (domain.Inventory, error)
	Update(ctx context.Context, inventory domain.Inventory) (domain.Inventory, error)
	Delete(ctx context.Context, inventory domain.Inventory) error
	FindById(ctx context.Context, productId string) (domain.Inventory, error)
	FindAll(ctx context.Context) ([]domain.Inventory, error)
}
