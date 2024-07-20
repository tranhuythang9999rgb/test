package controllers

import (
	"ap_sell_products/api/resources"
	"ap_sell_products/core/entities"
	"ap_sell_products/core/usecase"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type OrderController struct {
	order *usecase.OrderUseCase
}

func NewOrderController(order *usecase.OrderUseCase) *OrderController {
	return &OrderController{
		order: order,
	}
}
func (u *OrderController) RegisterOrder(ctx *fiber.Ctx) error {
	tokenString := ctx.Get("Authorization")
	var req entities.Order
	if err := ctx.BodyParser(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return err
	}
	err := u.order.RegisterOrder(ctx.Context(), tokenString, &req)
	if err != nil {
		return ctx.JSON(err)
	}
	resources.ResponseSuccess(ctx)
	return nil
}
