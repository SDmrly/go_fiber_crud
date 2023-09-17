package apps

import (
	"errors"
	"github.com/SDmrly/go_fiber_crud/models"
	"github.com/SDmrly/go_fiber_crud/utils"
	"gorm.io/gorm"
)

type Repository interface {
	Save(user models.User)
	FindAll() []models.User
	Update(user models.User)
	FindByID(userId int) (models.User, error)
	ChangePassword(user models.User)
	Delete(userId int)
	Migration() error
}

type database struct {
	Db *gorm.DB
}

func UserRepository(Db *gorm.DB) Repository {
	return &database{Db: Db}
}

func (db *database) Save(user models.User) {
	result := db.Db.Create(&user)
	utils.ErrorPanics(result.Error)
}

func (db *database) FindAll() []models.User {
	var users []models.User
	result := db.Db.Find(&users)
	utils.ErrorPanics(result.Error)

	return users
}

func (db *database) Update(user models.User) {
	result := db.Db.Model(&user).Updates(user)
	utils.ErrorPanics(result.Error)
}

func (db *database) FindByID(userId int) (models.User, error) {
	var user models.User

	if result := db.Db.Find(&user, userId); result != nil {
		return user, nil
	} else {
		return user, errors.New("not found user")
	}
}

func (db *database) ChangePassword(user models.User) {
	db.Update(user)
}

func (db *database) Delete(userId int) {
	var user models.User
	result := db.Db.Find(&user, userId).Delete(&user)
	utils.ErrorPanics(result.Error)
}

func (db *database) Migration() error {
	return db.Db.AutoMigrate(&models.User{})
}
