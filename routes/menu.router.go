package routes

import(
	"github.com/gin-gonic/gin"
	controller "resturant-management/controllers"
)

func MenuRoutes(incomingRoutes *gin.Engine){

	incomingRoutes.GET("/menus", controller.GetMenus())
	incomingRoutes.GET("/menu/:id", controller.GetMenu())
	incomingRoutes.POST("/menus", controller.CreateMenu())
	incomingRoutes.PATCH("/menus/:id", controller.UpdateMenu())
}