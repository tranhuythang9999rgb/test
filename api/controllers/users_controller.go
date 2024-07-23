package controllers

import (
	"ap_sell_products/common/utils"
	"ap_sell_products/core/entities"
	"ap_sell_products/core/usecase"
	"fmt"
	"path/filepath"

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
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	file, err := ctx.FormFile("avatar")
	if err == nil {
		// Nếu có file được upload
		ext := filepath.Ext(file.Filename)
		newFileName := fmt.Sprintln(utils.GenerateUniqueKey()) + ext
		uploadPath := filepath.Join("public", newFileName)

		if err := ctx.SaveFile(file, uploadPath); err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save file"})
		}

		req.Avatar = fmt.Sprintf("http://localhost:8080/download/%s", newFileName)
	}

	err = u.user.AddUser(ctx.Context(), &req)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "User added successfully", "avatar": req.Avatar})
}
