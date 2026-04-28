package routes

import (
	controller "go-resm/controllers"

	"github.com/gin-gonic/gin"
)

func UserRouters(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUsers())
	incomingRoutes.POST("/users/signup", controller.SignUp())
	incomingRoutes.POST("/users/login", controller.login())

}
