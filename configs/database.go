package configs

import (
	"fmt"
	"github.com/SDmrly/go_fiber_crud/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func DatabaseConnection(cnf *Config) *gorm.DB {
	dbInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cnf.DBHost, cnf.DBPort, cnf.DBUser, cnf.DBPassword, cnf.DBName)

	db, err := gorm.Open(postgres.Open(dbInfo), &gorm.Config{})
	utils.ErrorPanics(err)
	log.Println("Connected successfully to the database!")

	return db
}
