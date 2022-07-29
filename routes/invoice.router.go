package routes

import(
	"github.com/gin-gonic/gin"
	controller "resturant-management/controllers"
)

func InvoiceRoutes(incomingRoutes *gin.Engine){

	incomingRoutes.GET("/invoices", controller.GetInvoices())
	incomingRoutes.GET("/invoice/:id", controller.GetInvoice())
	incomingRoutes.POST("/invoices", controller.CreateInvoice())
	incomingRoutes.PATCH("/invoice/:id", controller.UpdateInvoice())
}