package main

import (
	"github.com/SDmrly/go_fiber_crud/apps"
	"github.com/SDmrly/go_fiber_crud/configs"
	"github.com/SDmrly/go_fiber_crud/handlers"
	"github.com/SDmrly/go_fiber_crud/routers"
	"github.com/SDmrly/go_fiber_crud/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"log"
)

var (
	db  *gorm.DB
	err error
)

func init() {
	loadCnf, err := configs.LoadConfig(".")
	utils.ErrorLogs("Could not load environment variables", err)

	db = configs.DatabaseConnection(loadCnf)
}

func main() {
	validate := validator.New()

	userRepo := apps.UserRepository(db)
	err = userRepo.Migration()
	utils.ErrorPanics(err)

	userService := apps.UserService(userRepo, validate)
	userHandler := handlers.UserHandler(userService)

	app := fiber.New()
	app.Mount("/api", routers.UserRouters(app, userHandler))
	log.Fatal(app.Listen(":8080"))

}
