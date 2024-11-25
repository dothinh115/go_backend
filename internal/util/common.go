package util

import (
	"os"

	"github.com/fatih/color"
)

func GetEnv(key string) string {
	value := os.Getenv(key)
	return value
}

func SuccessLog(value string) {
	successColor := color.New(color.BgGreen).Add(color.Bold)
	successColor.Println(value)
}
