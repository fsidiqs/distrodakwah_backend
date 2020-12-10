package router

import (
	"distrodakwah_backend/app/database"
	ddMiddleware "distrodakwah_backend/app/middleware"
	"distrodakwah_backend/app/services/controller/inventorycontroller"
	"distrodakwah_backend/app/services/repository/inventoryrepository"

	"github.com/labstack/echo"
)

var (
	inventoryController inventorycontroller.InventoryController
)

func inventoryAdminRole(g *echo.Group) {
	g.GET("", inventoryController.GetProductStocks)
	g.GET("/:item_inventory_id", inventoryController.GetProductStock)
	g.GET("/generate-import-inventory", inventoryController.ExportStocks)
	g.POST("/update-stocks", inventoryController.ImportStocks)
}

func SetInventoryGroup(g *echo.Group) {
	inventoryRepository := inventoryrepository.InventoryRepository{database.DB}
	inventoryController = inventorycontroller.InventoryController{&inventoryRepository}

	adminRoleMiddleware := g.Group("", ddMiddleware.AdminRoleMiddleware)
	inventoryAdminRole(adminRoleMiddleware)

}
