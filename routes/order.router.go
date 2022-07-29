package routes

import(
	"github.com/gin-gonic/gin"
	controller "resturant-management/controllers"
)

func OrderRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/orders", controllers.getOrders())
	incomingRoutes.GET("/order/:id", controllers.getOrder())
	incomingRoutes.POST("/orders", controllers.CreateOrder())
	incomingRoutes.PATCH("/orders/:id", controllers.EditOrder())
}