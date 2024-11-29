package controllers

import (
	"project/internal/dto"
	"project/internal/repos"
	"project/internal/router"

	"github.com/gin-gonic/gin"
)

func init() {
	router.ControllerRegister(&router.Controller{
		Path:    "/posts",
		Method:  router.GET,
		Handler: getAllPost,
	})

	router.ControllerRegister(&router.Controller{
		Path:    "/post",
		Method:  router.POST,
		Handler: createNewPost,
	})
}

func getAllPost(ctx *gin.Context) (interface{}, error) {
	data, err := repos.Post().GetAllPosts()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func createNewPost(ctx *gin.Context) (interface{}, error) {
	var newPost dto.CreatePost
	if e := ctx.ShouldBindJSON(&newPost); e != nil {
		return nil, e
	}

	data, err := repos.Post().CreateNewPost(&newPost)
	if err != nil {
		return nil, err
	}
	return data, nil
}
