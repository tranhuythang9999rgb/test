package usecase

import (
	"ap_sell_products/common/errors"
	"ap_sell_products/common/log"
	"ap_sell_products/common/utils"
	"ap_sell_products/core/domain"
	"ap_sell_products/core/entities"
	"context"
	"net/http"

	"gorm.io/gorm"
)

type OrderUseCase struct {
	order   domain.RepositoryOrder
	product domain.RepositoryProducts
	jwt     *JwtUseCase
	trans   domain.RepositoryTrans
}

func NewOrderUseCase(order domain.RepositoryOrder, product domain.RepositoryProducts, jwt *JwtUseCase, trans domain.RepositoryTrans) *OrderUseCase {
	return &OrderUseCase{
		order:   order,
		product: product,
		jwt:     jwt,
		trans:   trans,
	}
}

func (u *OrderUseCase) RegisterOrder(ctx context.Context, token string, req *entities.Order) errors.Error {
	user, err := u.jwt.VerifyToken(ctx, token)
	if err != nil {
		return errors.NewCustomHttpError(http.StatusUnauthorized, errors.NOT_EXIST_CODE, errors.NOT_EXIST_MESS)
	}

	err = u.trans.WithTransaction(func(tx *gorm.DB) error {
		product, err := u.product.GetProductByID(ctx, req.ProductID)
		if err != nil {
			log.Error("Error retrieving product: ", err)
			return errors.NewCustomHttpError(http.StatusInternalServerError, errors.SYSTEM_ERROR_CODE, "Failed to retrieve product: "+err.Error())
		}

		if product.Quantity < req.Quantity {
			return errors.NewCustomHttpError(http.StatusConflict, errors.NOT_EXIST_CODE, "Not enough product quantity")
		}

		err = u.product.UpdateProductQuantityById(ctx, tx, product.ID, product.Quantity-req.Quantity)
		if err != nil {
			log.Error("Error updating product quantity: ", err)
			return errors.NewCustomHttpError(http.StatusInternalServerError, errors.SYSTEM_ERROR_CODE, "Failed to update product quantity: "+err.Error())
		}

		price := product.Price * float64(req.Quantity)
		order := &domain.Order{
			ID:         utils.GenerateUniqueKey(),
			ProductID:  product.ID,
			CreatorID:  user.CreatorID,
			Price:      price,
			Quantity:   req.Quantity,
			Status:     true,
			CreateTime: utils.GenTimeStemp(),
			UpdateTime: utils.GenTimeStemp(),
		}
		if err := tx.Create(order).Error; err != nil {
			log.Error("Error creating order: ", err)
			return errors.NewCustomHttpError(http.StatusInternalServerError, errors.SYSTEM_ERROR_CODE, "Failed to create order: "+err.Error())
		}

		return nil
	})

	if err != nil {
		return errors.NewCustomHttpError(http.StatusConflict, errors.SYSTEM_ERROR_CODE, "Transaction failed: "+err.Error())
	}

	return nil
}
