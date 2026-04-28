package routes

import (
	"github.com/gin-gonic/gin"

	controller "go-resm/controllers"
)

func TableRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/tables", controller.GetTables())
	incomingRoutes.GET("/tables/:table_id", controller.GetTable())
	incomingRoutes.POST("/tables", controller.CreatTable())
	incomingRoutes.PATCH("/tables/:table_id", controller.UpdateTable())

}

