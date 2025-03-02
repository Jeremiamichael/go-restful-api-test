package repository

import (
	"context"
	"errors"
	"github.com/aronipurwanto/go-restful-api/model/domain"
	"gorm.io/gorm"
)

type PaymentRepositoryImpl struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &PaymentRepositoryImpl{db: db}
}

// Save payment
func (repository *PaymentRepositoryImpl) Save(ctx context.Context, payment domain.Payment) (domain.Payment, error) {
	if err := repository.db.WithContext(ctx).Create(&payment).Error; err != nil {
		return domain.Payment{}, err
	}
	return payment, nil
}

// Update payment
func (repository *PaymentRepositoryImpl) Update(ctx context.Context, payment domain.Payment) (domain.Payment, error) {
	if err := repository.db.WithContext(ctx).Save(&payment).Error; err != nil {
		return domain.Payment{}, err
	}
	return payment, nil
}

// Delete payment
func (repository *PaymentRepositoryImpl) Delete(ctx context.Context, payment domain.Payment) error {
	if err := repository.db.WithContext(ctx).Delete(&payment).Error; err != nil {
		return err
	}
	return nil
}

// FindById - Get payment by Payment ID
func (repository *PaymentRepositoryImpl) FindById(ctx context.Context, paymentId string) (domain.Payment, error) {
	var payment domain.Payment
	err := repository.db.WithContext(ctx).First(&payment, "payment_id = ?", paymentId).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return payment, errors.New("payment is not found")
	}
	return payment, err
}

// FindAll - Get all payments
func (repository *PaymentRepositoryImpl) FindAll(ctx context.Context) ([]domain.Payment, error) {
	var payments []domain.Payment
	err := repository.db.WithContext(ctx).Find(&payments).Error
	return payments, err
}

// FindByOrderId - Get payments by Order ID
func (repository *PaymentRepositoryImpl) FindByOrderId(ctx context.Context, orderId string) ([]domain.Payment, error) {
	var payments []domain.Payment
	err := repository.db.WithContext(ctx).Where("order_id = ?", orderId).Find(&payments).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return payments, errors.New("no payments found for this order")
	}
	return payments, err
}
