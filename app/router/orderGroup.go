package router

import (
	"distrodakwah_backend/app/database"
	"distrodakwah_backend/app/services/controller/ordercontroller"
	"distrodakwah_backend/app/services/repository/orderrepository"

	"github.com/labstack/echo"
)

var (
	orderController ordercontroller.OrderController
)

func SetOrderGroup(g *echo.Group) {
	orderRepository := orderrepository.OrderRepository{DB: database.DB}
	orderController = ordercontroller.OrderController{&orderRepository}

	// g.GET("", inventoryController.GetProductStocks)
	// g.GET("/:related_id", inventoryController.GetProductStock)
	// g.GET("/generate-import-inventory", inventoryController.GenerateExportInventoryTemplate)
	g.POST("", orderController.PostOrder)
}
