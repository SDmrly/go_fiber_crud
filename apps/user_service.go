package apps

import (
	"github.com/SDmrly/go_fiber_crud/models"
	"github.com/SDmrly/go_fiber_crud/utils"
	"github.com/go-playground/validator/v10"
)

type Service interface {
	Create(user models.CreateUser)
	FindAll() []models.UserResponse
	Update(user models.UpdateUser)
	ChangePassword(user models.UpdatePassword)
	FindByID(userId int) models.UserResponse
	Delete(userId int)
}

type service struct {
	repo     Repository
	validate *validator.Validate
}

func UserService(repo Repository, validate *validator.Validate) Service {
	return &service{
		repo:     repo,
		validate: validate,
	}
}

func (s *service) Create(user models.CreateUser) {
	err := s.validate.Struct(user)
	utils.ErrorPanics(err)

	newUser := models.User{
		UserName:  user.UserName,
		Password:  user.Password,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}

	s.repo.Save(newUser)
}

func (s *service) FindAll() []models.UserResponse {
	result := s.repo.FindAll()

	var users []models.UserResponse

	for _, value := range result {
		user := models.UserResponse{
			Id:        value.Id,
			UserName:  value.UserName,
			Password:  value.Password,
			FirstName: value.FirstName,
			LastName:  value.LastName,
			Email:     value.Email,
		}
		users = append(users, user)
	}

	return users
}

func (s *service) Update(user models.UpdateUser) {
	userData, err := s.repo.FindByID(user.Id)
	utils.ErrorPanics(err)

	userData.FirstName = user.FirstName
	userData.LastName = user.LastName
	userData.Email = user.Email

	s.repo.Update(userData)
}

func (s *service) ChangePassword(user models.UpdatePassword) {
	userPassword, err := s.repo.FindByID(user.Id)
	utils.ErrorPanics(err)

	userPassword.Password = user.Password

	s.repo.ChangePassword(userPassword)
}

func (s *service) FindByID(userId int) models.UserResponse {
	user, err := s.repo.FindByID(userId)
	utils.ErrorPanics(err)

	userResponse := models.UserResponse{
		Id:        user.Id,
		UserName:  user.UserName,
		Password:  user.Password,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}
	return userResponse
}

func (s *service) Delete(userId int) {
	s.repo.Delete(userId)
}
