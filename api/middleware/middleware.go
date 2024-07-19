package middleware

import (
	"ap_sell_products/core/usecase"
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type MiddleWare struct {
	jwtUseCase *usecase.JwtUseCase
}

func NewMiddleware(
	jwtUseCase *usecase.JwtUseCase,

) *MiddleWare {
	return &MiddleWare{
		jwtUseCase: jwtUseCase,
	}
}
func (m *MiddleWare) Authenticate() fiber.Handler {

	return func(c *fiber.Ctx) error {
		tokenString := c.Get("Authorization")
		if tokenString == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "request does not contain an access token",
			})
		}

		_, err := m.jwtUseCase.VerifyToken(c.Context(), tokenString)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid or expired token",
			})
		}
		listSession, err := m.jwtUseCase.ListSession(context.Background(), tokenString)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": fmt.Sprintln(err),
			})
		}
		var isCheck bool
		for _, v := range listSession {
			if v.Token == tokenString {
				isCheck = true
				break
			}
		}
		if !isCheck {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid or expired token",
			})
		}

		return c.Next()
	}
}
