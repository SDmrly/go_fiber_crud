package main

import (
	"github.com/SDmrly/go_fiber_crud/configs"
	"github.com/SDmrly/go_fiber_crud/router"
	"github.com/SDmrly/go_fiber_crud/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"log"
)

var (
	db *gorm.DB
)

func init() {
	loadCnf, err := configs.LoadConfig(".")
	utils.ErrorLogs("Could not load environment variables", err)

	db = configs.DatabaseConnection(loadCnf)
}

func main() {
	app := fiber.New()
	userAPI := router.UserAPI{App: app, DB: db}
	userAPI.Router()
	log.Fatal(app.Listen(":8080"))

}
