package router

import (
	"github.com/labstack/echo"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/database"
	ddMiddleware "github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/middleware"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/controller/inventorycontroller"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/repository/inventoryrepository"
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
