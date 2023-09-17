package models

type User struct {
	Id        uint   `gorm:"type:int;primary_key"`
	UserName  string `gorm:"not null" json:"user_name,omitempty"`
	Password  string `gorm:"not null" json:"password,omitempty"`
	FirstName string `gorm:"not null" json:"first_name,omitempty"`
	LastName  string `gorm:"not null" json:"last_name,omitempty"`
	Email     string `gorm:"not null" json:"email,omitempty"`
}

type CreateUser struct {
	UserName  string `validate:"required,min=6,max=16" json:"user_name"`
	Password  string `validate:"required,min=8,max=16" json:"password"`
	FirstName string `validate:"required,min=2" json:"first_name"`
	LastName  string `validate:"required,min=2" json:"last_name"`
	Email     string `validate:"required,email"`
}

type UpdateUser struct {
	Id        int    `validate:"required"`
	FirstName string `validate:"required,min=2" json:"first_name"`
	LastName  string `validate:"required,min=2" json:"last_name"`
	Email     string `validate:"required,email"`
}

type UpdatePassword struct {
	Id       int    `validate:"required"`
	Password string `validate:"required,min=8,max=16" json:"password"`
}

type UserResponse struct {
	Id        uint   `json:"id"`
	UserName  string `json:"user_name"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}
