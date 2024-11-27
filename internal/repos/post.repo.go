package repos

import (
	"project/internal/database"
	"project/internal/models"
	"project/internal/orm"

	"gorm.io/gorm"
)

type postRepository interface {
	GetAllPosts() (interface{}, error)
	CreateNewPost(newPost *models.CreatePost) (interface{}, error)
}

type postRepo struct {
	DB *gorm.DB
}

func (pr *postRepo) GetAllPosts() (interface{}, error) {
	var posts []models.Post
	if err := pr.DB.Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (pr *postRepo) CreateNewPost(newPost *models.CreatePost) (interface{}, error) {

	var post models.Post
	orm.Service().Create(newPost, &post)
	if er := pr.DB.Create(&post).Error; er != nil {
		return nil, er

	}
	return post, nil
}

func Post() postRepository {
	return &postRepo{
		DB: database.GetDb(),
	}
}
