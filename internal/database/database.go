package database

import (
	"fmt"
	"log"
	"project/internal/util"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DbHost     string = util.GetEnv("DB_HOST")
	DbUsername string = util.GetEnv("DB_USERNAME")
	DbPassword string = util.GetEnv("DB_PASSWORD")
	DbName     string = util.GetEnv("DB_NAME")
)

var DB *gorm.DB
var err error

func ConnectDB() {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=5432", DbUsername, DbPassword, DbName, DbHost)
	DB, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Fatal(err)
	}
	util.SuccessLog("Connect db thành công!")
}