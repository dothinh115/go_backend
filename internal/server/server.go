package server

import (
	"log"
	"net/http"
	"project/internal/router"

	"github.com/gin-gonic/gin"
)

func init() {
	// gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	controllers := router.GetControllers()
	for _, routeData := range controllers {
		switch routeData.Method {
		case "GET":
			{
				r.GET(routeData.Path, router.ExecuteHandler(routeData.Handler))
			}
		case "POST":
			{
				r.POST(routeData.Path, router.ExecuteHandler(routeData.Handler))
			}
		case "PATCH":
			{
				r.PATCH(routeData.Path, router.ExecuteHandler(routeData.Handler))
			}
		case "DELETE":
			{
				r.DELETE(routeData.Path, router.ExecuteHandler(routeData.Handler))
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
