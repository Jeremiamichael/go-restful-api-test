package repository

import (
	"context"
	"errors"
	"github.com/aronipurwanto/go-restful-api/model/domain"
	"gorm.io/gorm"
)

type InventoryRepositoryImpl struct {
	db *gorm.DB
}

func NewInventoryRepository(db *gorm.DB) InventoryRepository {
	return &InventoryRepositoryImpl{db: db}
}

// Save inventory
func (repository *InventoryRepositoryImpl) Save(ctx context.Context, inventory domain.Inventory) (domain.Inventory, error) {
	if err := repository.db.WithContext(ctx).Create(&inventory).Error; err != nil {
		return domain.Inventory{}, err
	}
	return inventory, nil
}

// Update inventory
func (repository *InventoryRepositoryImpl) Update(ctx context.Context, inventory domain.Inventory) (domain.Inventory, error) {
	if err := repository.db.WithContext(ctx).Save(&inventory).Error; err != nil {
		return domain.Inventory{}, err
	}
	return inventory, nil
}

// Delete inventory
func (repository *InventoryRepositoryImpl) Delete(ctx context.Context, inventory domain.Inventory) error {
	if err := repository.db.WithContext(ctx).Delete(&inventory).Error; err != nil {
		return err
	}
	return nil
}

// FindById - Get inventory by Product ID
func (repository *InventoryRepositoryImpl) FindById(ctx context.Context, productId string) (domain.Inventory, error) {
	var inventory domain.Inventory
	err := repository.db.WithContext(ctx).First(&inventory, "product_id = ?", productId).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return inventory, errors.New("inventory record is not found")
	}
	return inventory, err
}

// FindAll - Get all inventory records
func (repository *InventoryRepositoryImpl) FindAll(ctx context.Context) ([]domain.Inventory, error) {
	var inventories []domain.Inventory
	err := repository.db.WithContext(ctx).Find(&inventories).Error
	return inventories, err
}
