package controllers

import (
	"project/internal/dto"
	"project/internal/errors"
	"project/internal/repos"
	"project/internal/router"

	"github.com/gin-gonic/gin"
)

func init() {
	router.ControllerRegister(&router.Controller{
		Path:    "/posts",
		Method:  router.GET,
		Handler: getAll,
	})

	router.ControllerRegister(&router.Controller{
		Path:    "/post",
		Method:  router.POST,
		Handler: create,
	})

	router.ControllerRegister(&router.Controller{
		Path:    "/post/:id",
		Method:  router.PATCH,
		Handler: update,
	})
}

func getAll(ctx *gin.Context) (interface{}, error) {
	data, err := repos.Post().GetAll()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func create(ctx *gin.Context) (interface{}, error) {
	if ctx.Request.ContentLength == 0 {
		return nil, errors.NewBadRequestException("Request không có body!")
	}
	var newPost dto.CreatePost
	if e := ctx.ShouldBindJSON(&newPost); e != nil {
		return nil, e
	}

	data, err := repos.Post().Create(&newPost)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func update(ctx *gin.Context) (interface{}, error) {
	if ctx.Request.ContentLength == 0 {
		return nil, errors.NewBadRequestException("Request không có body!")
	}

	id := ctx.Param("id")
	var updatedPost dto.UpdatePost
	if err := ctx.ShouldBindJSON(&updatedPost); err != nil {
		return nil, err
	}

	data, err := repos.Post().Update(id, &updatedPost)
	if err != nil {
		return nil, err
	}
	return data, nil
}
