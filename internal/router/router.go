package router

import (
	"project/internal/interceptors"
	"project/internal/observable"

	"github.com/gin-gonic/gin"
)

var controllers []Controller = make([]Controller, 0)

type Handler func(*gin.Context) (interface{}, error)

type Controller struct {
	Path    string
	Method  Method
	Handler Handler
}

func ControllerRegister(controller *Controller) {
	controllers = append(controllers, *controller)
}

func GetControllers() []Controller {
	return controllers
}

func ExecuteHandler(handler Handler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		data, err := handler(ctx)
		observable := observable.NewObservable()
		observable.Subscribe(func(data interface{}) {
			if err != nil {
				response := interceptors.HandleError(err)
				ctx.JSON(response.StatusCode, response)
			} else {
				ctx.JSON(200, data)

			}
		})

		interceptors := interceptors.GetInterceptors()

		for _, intercept := range interceptors {
			observable.Map(intercept)
		}

		if err != nil {
			observable.Next(err)
		} else {
			observable.Next(data)
		}

		observable.Complete()
	}
}
