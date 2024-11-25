package server

import (
	"log"
	"net/http"
	"project/internal/router"

	"github.com/gin-gonic/gin"
)

func ServerInit() {
	// gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	for _, routeData := range router.Controllers {
		switch routeData.Method {
		case "GET":
			{
				r.GET(routeData.Path, router.ExecuteHandler(routeData.Handler))
			}
		default:
			{
				r.Any(routeData.Path, func(ctx *gin.Context) {
					ctx.JSON(http.StatusMethodNotAllowed, gin.H{"message": "Method not allowed", "statusCode": 405})
				})
			}
		}

	}

	if err := r.Run(":3000"); err != nil {
		log.Fatalln(err)
	}
}
