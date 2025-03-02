package repository

import (
	"context"
	"errors"
	"github.com/aronipurwanto/go-restful-api/model/domain"
	"gorm.io/gorm"
)

type TaxRepositoryImpl struct {
	db *gorm.DB
}

func NewTaxRepository(db *gorm.DB) TaxRepository {
	return &TaxRepositoryImpl{db: db}
}

// Save tax
func (repository *TaxRepositoryImpl) Save(ctx context.Context, tax domain.Tax) (domain.Tax, error) {
	if err := repository.db.WithContext(ctx).Create(&tax).Error; err != nil {
		return domain.Tax{}, err
	}
	return tax, nil
}

// Update tax
func (repository *TaxRepositoryImpl) Update(ctx context.Context, tax domain.Tax) (domain.Tax, error) {
	if err := repository.db.WithContext(ctx).Save(&tax).Error; err != nil {
		return domain.Tax{}, err
	}
	return tax, nil
}

// Delete tax
func (repository *TaxRepositoryImpl) Delete(ctx context.Context, tax domain.Tax) error {
	if err := repository.db.WithContext(ctx).Delete(&tax).Error; err != nil {
		return err
	}
	return nil
}

// FindById - Get tax by ID
func (repository *TaxRepositoryImpl) FindById(ctx context.Context, taxId string) (domain.Tax, error) {
	var tax domain.Tax
	err := repository.db.WithContext(ctx).First(&tax, "tax_id = ?", taxId).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return tax, errors.New("tax is not found")
	}
	return tax, err
}

// FindAll - Get all taxes
func (repository *TaxRepositoryImpl) FindAll(ctx context.Context) ([]domain.Tax, error) {
	var taxes []domain.Tax
	err := repository.db.WithContext(ctx).Find(&taxes).Error
	return taxes, err
}

// FindByType - Get taxes by tax type
func (repository *TaxRepositoryImpl) FindByType(ctx context.Context, taxType string) ([]domain.Tax, error) {
	var taxes []domain.Tax
	err := repository.db.WithContext(ctx).Where("tax_type = ?", taxType).Find(&taxes).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return taxes, errors.New("no taxes found for this type")
	}
	return taxes, err
}
