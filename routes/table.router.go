package routes


import(
	"github.com/gin-gonic/gin"
	controller "resturant-management/controllers"
)

func TableRoutes(incomingRoutes *gin.Engine){

	incomingRoutes.Get("/tables", controller.GetTables())
	incomingRoutes.Get("/table/:id", controller.GetTable())
	incomingRoutes.POST("/tables", controller.CreateTable())
	incomingRoutes.PATCH("/tables/:id", controller.UpdateTable())

}