package router

import (
	"github.com/labstack/echo"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/database"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/order/controller"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/order/repository"
)

var (
	orderRepository repository.OrderRepository
	orderController controller.OrderController
)

func SetOrderGroup(g *echo.Group) {
	orderRepository = repository.OrderRepository{DB: database.DB}
	orderController = controller.OrderController{&orderRepository}

	// g.GET("", inventoryController.GetProductStocks)
	// g.GET("/:related_id", inventoryController.GetProductStock)
	// g.GET("/generate-import-inventory", inventoryController.GenerateExportInventoryTemplate)
	g.POST("", orderController.PostOrder)
}
