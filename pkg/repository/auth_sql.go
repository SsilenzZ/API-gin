package repository

import (
	"API"
	"gorm.io/gorm"
)

type AuthSQL struct {
	db *gorm.DB
}

func NewAuthSQL(db *gorm.DB) *AuthSQL {
	return &AuthSQL{db: db}
}

func (r *AuthSQL) CreateUser(user API_gin.Users) (int, error) {
	field := API_gin.Users{Email: user.Email, Password: user.Password}
	err := r.db.Create(&field).Error
	r.db.Select(field, "email", "password").Last(&field)

	return field.Id, err
}

func (r *AuthSQL) GetUser(email, password string) (API_gin.Users, error) {
	var user API_gin.Users

	err := r.db.Table("users").Take(&user).Where(email, password).Error

	return user, err
}
