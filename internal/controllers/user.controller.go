package controllers

import (
	"project/internal/repos"
	"project/internal/router"

	"github.com/gin-gonic/gin"
)

func init() {
	router.ControllerRegister(router.Controller{
		Path:    "/users",
		Method:  router.GET,
		Handler: getAllUsers,
	})
}

func getAllUsers(ctx *gin.Context) (interface{}, error) {
	users, err := repos.User.GetAllUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}
