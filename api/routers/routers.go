package routers

import (
	"ap_sell_products/api/controllers"
	"ap_sell_products/api/middleware"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type ApiRouter struct {
	Router *fiber.App
}

func NewApiRouter(
	user *controllers.UserController,
	jwt *controllers.JwtController,
	middleware *middleware.MiddleWare,
) *ApiRouter {

	r := fiber.New()
	r.Use(logger.New())
	r.Use(cors.New())
	r.Use(recover.New())

	r.Get("/ping", func(c *fiber.Ctx) error {
		c.JSON("ping")
		return nil
	})
	userGroup := r.Group("/user")
	{
		userGroup.Post("/register", user.AddUser) //thangth1 password 1234
		userGroup.Post("/login", jwt.Login)
		userMiddleware := userGroup.Group("/", middleware.Authenticate())
		{
			userMiddleware.Get("/list", jwt.ListSession)
			userMiddleware.Post("/logout", jwt.Logout)
		}

	}
	//http://localhost:8080/dowload/cell.png
	r.Get("/dowload/:name", func(c *fiber.Ctx) error {
		name := c.Params("name")
		err := filesystem.SendFile(c, http.Dir("public"), name)
		if err != nil {
			// Handle the error, e.g., return a 404 Not Found response
			return c.Status(fiber.StatusNotFound).SendString("File not found")
		}
		return nil
	})
	return &ApiRouter{
		Router: r,
	}
}
