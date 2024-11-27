package repos

import (
	"project/internal/database"
	"project/internal/models"

	"gorm.io/gorm"
)

type userRepository interface {
	GetAllUsers() (interface{}, error)
	GetUserById(id string) (interface{}, error)
}

type userRepo struct {
	DB *gorm.DB
}

func (ur *userRepo) GetAllUsers() (interface{}, error) {
	var users []models.User
	if err := ur.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (ur *userRepo) GetUserById(id string) (interface{}, error) {
	var user models.User
	if err := ur.DB.Where("id = ?", id).Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func User() userRepository {
	return &userRepo{
		DB: database.GetDb(),
	}
}
