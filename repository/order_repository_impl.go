package repository

import (
	"context"
	"errors"
	"github.com/aronipurwanto/go-restful-api/model/domain"
	"gorm.io/gorm"
)

type OrderRepositoryImpl struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &OrderRepositoryImpl{db: db}
}

// Save order
func (repository *OrderRepositoryImpl) Save(ctx context.Context, order domain.Order) (domain.Order, error) {
	if err := repository.db.WithContext(ctx).Create(&order).Error; err != nil {
		return domain.Order{}, err
	}
	return order, nil
}

// Update order
func (repository *OrderRepositoryImpl) Update(ctx context.Context, order domain.Order) (domain.Order, error) {
	if err := repository.db.WithContext(ctx).Save(&order).Error; err != nil {
		return domain.Order{}, err
	}
	return order, nil
}

// Delete order
func (repository *OrderRepositoryImpl) Delete(ctx context.Context, order domain.Order) error {
	if err := repository.db.WithContext(ctx).Delete(&order).Error; err != nil {
		return err
	}
	return nil
}

// FindById - Get order by ID
func (repository *OrderRepositoryImpl) FindById(ctx context.Context, orderId string) (domain.Order, error) {
	var order domain.Order
	err := repository.db.WithContext(ctx).Preload("OrderItems").First(&order, "order_id = ?", orderId).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return order, errors.New("order is not found")
	}
	return order, err
}

// FindAll - Get all orders
func (repository *OrderRepositoryImpl) FindAll(ctx context.Context) ([]domain.Order, error) {
	var orders []domain.Order
	err := repository.db.WithContext(ctx).Preload("OrderItems").Find(&orders).Error
	return orders, err
}
