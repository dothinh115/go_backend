package controllers

import (
	"project/internal/repos"
	"project/internal/router"

	"github.com/gin-gonic/gin"
)

func init() {
	router.ControllerRegister(&router.Controller{
		Path:    "/users",
		Method:  router.GET,
		Handler: getAllUsers,
	})

	router.ControllerRegister(&router.Controller{
		Path:    "/user/:id",
		Method:  router.GET,
		Handler: getUsersById,
	})
}

func getAllUsers(ctx *gin.Context) (interface{}, error) {
	users, err := repos.User().GetAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func getUsersById(ctx *gin.Context) (interface{}, error) {
	id := ctx.Param("id")
	user, err := repos.User().GetById(id)

	if err != nil {
		return nil, err
	}
	return user, nil
}
