package internal

import (
	_ "project/internal/controllers"
	"project/internal/database"
	_ "project/internal/env"
	"project/internal/server"
)

func init() {
	database.ConnectDB()
	server.ServerInit()
}
