package routes

import(
	"github.com/gin-gonic/gin"
	controller "resturant-management/controllers"

)

func FoodRoutes(incomingRoutes *gin.Engine){

	incomingRoutes.GET("/foods", controller.getFoods())
	incomingRoutes.GET('food/:id', controller.geFood())
	incomingRoutes.POST("/foods", controller.CreateFood())
	incomingRoutes.PATCH("foods/:id", controller.UpdateFood())
}