package repos

import (
	"project/internal/database"
	"project/internal/dto"
	"project/internal/errors"
	"project/internal/models"
	"project/internal/orm"

	"gorm.io/gorm"
)

type postRepository interface {
	GetAll() (interface{}, error)
	Create(newPost *dto.CreatePost) (interface{}, error)
	Update(id string, updatedPost *dto.UpdatePost) (interface{}, error)
}

type postRepo struct {
	DB *gorm.DB
}

func (pr *postRepo) GetAll() (interface{}, error) {
	var posts []models.Post
	if err := pr.DB.Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (pr *postRepo) Create(newPost *dto.CreatePost) (interface{}, error) {
	var post models.Post
	err := orm.Service().Create(newPost, &post)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (pr *postRepo) Update(id string, updatedPost *dto.UpdatePost) (interface{}, error) {
	var post models.Post
	if err := pr.DB.Where("id = ?", id).First(&post).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NewBadRequestException("Không có bài viết này!")
		}
		return nil, err
	}
	err := orm.Service().Update(updatedPost, &post)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func Post() postRepository {
	return &postRepo{
		DB: database.GetDb(),
	}
}
