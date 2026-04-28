package routes

import (
	"github.com/gin-gonic/gin"

	controller "go-resm/controllers"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orderItems", controller.GetOrderItems())
	incomingRoutes.GET("/orderItemss/:order_id", controller.GetOrderItem())
	incomingRoutes.GET("/orderItems-order/:order_id", controller.GetOrderItemByOrder())    
	incomingRoutes.POST("/orderItems", controller.CreatOrderItem())
	incomingRoutes.PATCH("/orderItems/:order_id", controller.UpdateOrderItem())


}
