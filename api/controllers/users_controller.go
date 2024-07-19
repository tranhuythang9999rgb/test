package controllers

import (
	"ap_sell_products/api/resources"
	"ap_sell_products/core/entities"
	"ap_sell_products/core/usecase"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	user *usecase.UserUseCase
}

func NewUserController(user *usecase.UserUseCase) *UserController {
	return &UserController{
		user: user,
	}
}
func (u *UserController) AddUser(ctx *fiber.Ctx) error {
	var req entities.User
	if err := ctx.BodyParser(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return err
	}
	err := u.user.AddUser(ctx.Context(), &req)
	if err != nil {
		ctx.JSON(err)
		return nil
	}
	resources.ResponseSuccess(ctx)
	return nil
}
