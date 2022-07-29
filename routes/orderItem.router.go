package routes

import(
	"github.com/gin-gonic/gin"
	controller "resturant-management/controllers"
)

func OrderRoutes(incomingRoutes *gin.Engine){

	incomingRoutes.GET('/orderItems', controller.GetOrderItems())
	incomingRoutes.GET('/orderItem/:id', controller.GetOrderItem())
	incomingRoutes.POST('/orderItems', controller.CreateOrderItem())
	incomingRoutes.PATCH('/orderItems/:id', controller.UpdateOrderItem())
}