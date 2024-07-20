package usecase

import (
	"ap_sell_products/common/errors"
	"ap_sell_products/common/utils"
	"ap_sell_products/core/domain"
	"ap_sell_products/core/entities"
	"context"
	"net/http"
)

type ProductUseCase struct {
	product domain.RepositoryProducts
}

func NewProductUseCase(product domain.RepositoryProducts) *ProductUseCase {
	return &ProductUseCase{
		product: product,
	}
}

func (u *ProductUseCase) AddProduct(ctx context.Context, req *entities.Product) errors.Error {
	product := domain.Product{
		ID:              utils.GenerateUniqueKey(),
		CreatorID:       req.CreatorID,
		Name:            req.Name,
		Price:           req.Price,
		Quantity:        req.Quantity,
		Description:     req.Description,
		DiscountPercent: req.DiscountPercent,
		StatusSell:      true,
		IsActive:        true,
		CreateTime:      utils.GenTimeStemp(),
		UpdateTime:      utils.GenTimeStemp(),
	}
	err := u.product.CreateProduct(ctx, &product)
	if err != nil {
		return errors.NewCustomHttpError(http.StatusInternalServerError, errors.SYSTEM_ERROR_CODE, errors.SYSTEM_ERROR_MESS)
	}
	return nil
}
