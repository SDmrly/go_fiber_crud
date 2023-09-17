package routers

import (
	"github.com/SDmrly/go_fiber_crud/handlers"
	"github.com/gofiber/fiber/v2"
)

func UserRouters(handler handlers.Handler) *fiber.App {
	routers := fiber.New()

	routers.Get("/healt", func(ctx *fiber.Ctx) error {
		return ctx.Status(200).JSON(fiber.Map{
			"status":  "success",
			"message": "Welcome to Golang fiber and gorm",
		})
	})

	routers.Route("/users", func(router fiber.Router) {
		router.Post("", handler.Create)
		router.Get("", handler.FindAll)
	})

	routers.Route("/users/:userId", func(router fiber.Router) {
		router.Patch("", handler.Update)
		router.Get("", handler.FindByID)
		router.Delete("", handler.Delete)
	})

	routers.Route("/users/change_password/:userId", func(router fiber.Router) {
		router.Patch("", handler.ChangePassword)
	})

	return routers
}
