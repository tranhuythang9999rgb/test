package controllers

import (
	"ap_sell_products/api/resources"
	"ap_sell_products/core/entities"
	"ap_sell_products/core/usecase"
	"context"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type JwtController struct {
	jwt *usecase.JwtUseCase
}

func NewJwtController(jwt *usecase.JwtUseCase) *JwtController {
	return &JwtController{
		jwt: jwt,
	}
}
func (u *JwtController) Login(ctx *fiber.Ctx) error {
	var req entities.LoginRequest
	if err := ctx.BodyParser(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return err
	}
	resp, err := u.jwt.Login(context.Background(), &req)
	if err != nil {
		ctx.JSON(err)
		return nil
	}
	resources.ResponseSuccess(ctx, resp)
	return nil
}
func (u *JwtController) Logout(ctx *fiber.Ctx) error {
	tokenString := ctx.Get("Authorization")

	err := u.jwt.Logout(context.Background(), tokenString)
	if err != nil {
		return ctx.JSON(err)
	}

	return resources.ResponseSuccess(ctx)
}
func (u *JwtController) ListSession(ctx *fiber.Ctx) error {
	tokenString := ctx.Get("Authorization")
	resp, err := u.jwt.ListSession(ctx.Context(), tokenString)
	if err != nil {
		ctx.JSON(err)
		return nil
	}
	resources.ResponseSuccess(ctx, resp)
	return nil
}
