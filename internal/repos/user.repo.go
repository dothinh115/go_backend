package repos

import (
	"project/internal/database"
	"project/internal/models"

	"gorm.io/gorm"
)

type userRepository interface {
	GetAllUsers() (interface{}, error)
}

type userRepo struct {
	DB *gorm.DB
}

func (ur *userRepo) GetAllUsers() (interface{}, error) {
	var users []models.User
	if err := database.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

var User userRepository = &userRepo{
	DB: database.DB,
}
