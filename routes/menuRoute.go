package routes

import (
	"github.com/gin-gonic/gin"

	controller "go-resm/controllers"
)

func MenuRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/menus", controller.GetMenus())
	incomingRoutes.GET("/menus/:menu_id", controller.GetMenu())
	incomingRoutes.POST("/menus", controller.CreatMenu())
	incomingRoutes.PATCH("/menus/:menu_id", controller.UpdateMenu())

}
