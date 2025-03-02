package service

import (
	"context"
	"errors"
	"github.com/aronipurwanto/go-restful-api/exception"
	"github.com/aronipurwanto/go-restful-api/helper"
	"github.com/aronipurwanto/go-restful-api/model/domain"
	"github.com/aronipurwanto/go-restful-api/model/web"
	"github.com/aronipurwanto/go-restful-api/repository"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DiscountServiceImpl struct {
	DiscountRepository repository.DiscountRepository
	Validate           *validator.Validate
}

func NewDiscountService(discountRepository repository.DiscountRepository, validate *validator.Validate) DiscountService {
	return &DiscountServiceImpl{
		DiscountRepository: discountRepository,
		Validate:           validate,
	}
}

// Create Discount
func (service *DiscountServiceImpl) Create(ctx context.Context, request web.DiscountCreateRequest) (web.DiscountResponse, error) {
	if err := service.Validate.Struct(request); err != nil {
		return web.DiscountResponse{}, err
	}

	discount := domain.Discount{
		DiscountID:  uuid.New().String(), // Generate UUID for discount ID
		Description: request.Description,
		DiscountPct: request.DiscountPct,
		ValidFrom:   request.ValidFrom,
		ValidUntil:  request.ValidUntil,
	}

	savedDiscount, err := service.DiscountRepository.Save(ctx, discount)
	if err != nil {
		return web.DiscountResponse{}, err
	}

	return helper.ToDiscountResponse(savedDiscount), nil
}

// Update Discount
func (service *DiscountServiceImpl) Update(ctx context.Context, request web.DiscountUpdateRequest) (web.DiscountResponse, error) {
	if err := service.Validate.Struct(request); err != nil {
		return web.DiscountResponse{}, err
	}

	discount, err := service.DiscountRepository.FindById(ctx, request.DiscountID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return web.DiscountResponse{}, exception.NewNotFoundError("Discount not found")
	} else if err != nil {
		return web.DiscountResponse{}, err
	}

	discount.Description = request.Description
	discount.DiscountPct = request.DiscountPct
	discount.ValidFrom = request.ValidFrom
	discount.ValidUntil = request.ValidUntil

	updatedDiscount, err := service.DiscountRepository.Update(ctx, discount)
	if err != nil {
		return web.DiscountResponse{}, err
	}

	return helper.ToDiscountResponse(updatedDiscount), nil
}

// Delete Discount
func (service *DiscountServiceImpl) Delete(ctx context.Context, discountId string) error {
	discount, err := service.DiscountRepository.FindById(ctx, discountId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return exception.NewNotFoundError("Discount not found")
	} else if err != nil {
		return err
	}

	return service.DiscountRepository.Delete(ctx, discount)
}

// Find Discount By ID
func (service *DiscountServiceImpl) FindById(ctx context.Context, discountId string) (web.DiscountResponse, error) {
	discount, err := service.DiscountRepository.FindById(ctx, discountId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return web.DiscountResponse{}, exception.NewNotFoundError("Discount not found")
	} else if err != nil {
		return web.DiscountResponse{}, err
	}

	return helper.ToDiscountResponse(discount), nil
}

// Find All Discounts
func (service *DiscountServiceImpl) FindAll(ctx context.Context) ([]web.DiscountResponse, error) {
	discounts, err := service.DiscountRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return helper.ToDiscountResponses(discounts), nil
}
