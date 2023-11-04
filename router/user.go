package router

import (
	"github.com/SDmrly/go_fiber_crud/api"
	"github.com/SDmrly/go_fiber_crud/pkg/repository"
	"github.com/SDmrly/go_fiber_crud/pkg/service"
	"github.com/SDmrly/go_fiber_crud/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserAPI struct {
	DB  *gorm.DB
	App *fiber.App
}

func InitUserAPI(db *gorm.DB) api.Handler {
	validate := validator.New()
	userRepo := repository.UserRepository(db)
	err := userRepo.Migration()
	utils.ErrorPanics(err)

	userService := service.UserService(userRepo, validate)
	userHandler := api.UserHandler(userService)

	return userHandler
}

func (api *UserAPI) Router() {
	handler := InitUserAPI(api.DB)

	routers := api.App.Group("/api")

	routers.Get("/healt", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
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
}
