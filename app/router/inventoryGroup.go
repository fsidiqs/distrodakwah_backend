package router

import (
	"github.com/labstack/echo"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/database"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/inventory/controller"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/inventory/repository"
)

var (
	inventoryRepository repository.InventoryRepository
	inventoryController controller.InventoryController
)

func SetInventoryGroup(g *echo.Group) {
	inventoryRepository = repository.InventoryRepository{database.DB}
	inventoryController = controller.InventoryController{&inventoryRepository}

	g.GET("", inventoryController.GetProductStocks)
	g.GET("/:item_inventory_id", inventoryController.GetProductStock)
	g.GET("/generate-import-inventory", inventoryController.ExportStocks)
	g.POST("/update-stocks", inventoryController.ImportStocks)
}
