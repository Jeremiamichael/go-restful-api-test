package repository

import (
	"context"
	"errors"
	"github.com/aronipurwanto/go-restful-api/model/domain"
	"gorm.io/gorm"
)

type ReceiptRepositoryImpl struct {
	db *gorm.DB
}

func NewReceiptRepository(db *gorm.DB) ReceiptRepository {
	return &ReceiptRepositoryImpl{db: db}
}

// Save receipt
func (repository *ReceiptRepositoryImpl) Save(ctx context.Context, receipt domain.Receipt) (domain.Receipt, error) {
	if err := repository.db.WithContext(ctx).Create(&receipt).Error; err != nil {
		return domain.Receipt{}, err
	}
	return receipt, nil
}

// Update receipt
func (repository *ReceiptRepositoryImpl) Update(ctx context.Context, receipt domain.Receipt) (domain.Receipt, error) {
	if err := repository.db.WithContext(ctx).Save(&receipt).Error; err != nil {
		return domain.Receipt{}, err
	}
	return receipt, nil
}

// Delete receipt
func (repository *ReceiptRepositoryImpl) Delete(ctx context.Context, receipt domain.Receipt) error {
	if err := repository.db.WithContext(ctx).Delete(&receipt).Error; err != nil {
		return err
	}
	return nil
}

// FindById - Get receipt by ID
func (repository *ReceiptRepositoryImpl) FindById(ctx context.Context, receiptId string) (domain.Receipt, error) {
	var receipt domain.Receipt
	err := repository.db.WithContext(ctx).First(&receipt, "receipt_id = ?", receiptId).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return receipt, errors.New("receipt is not found")
	}
	return receipt, err
}

// FindAll - Get all receipts
func (repository *ReceiptRepositoryImpl) FindAll(ctx context.Context) ([]domain.Receipt, error) {
	var receipts []domain.Receipt
	err := repository.db.WithContext(ctx).Find(&receipts).Error
	return receipts, err
}

// FindByOrderId - Get receipts by Order ID
func (repository *ReceiptRepositoryImpl) FindByOrderId(ctx context.Context, orderId string) ([]domain.Receipt, error) {
	var receipts []domain.Receipt
	err := repository.db.WithContext(ctx).Where("order_id = ?", orderId).Find(&receipts).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return receipts, errors.New("no receipts found for this order")
	}
	return receipts, err
}
