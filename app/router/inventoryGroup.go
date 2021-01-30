package router

import (
	"distrodakwah_backend/app/database"
	"distrodakwah_backend/app/services/controller/inventorycontroller"
	"distrodakwah_backend/app/services/repository/inventoryrepository"

	"github.com/labstack/echo"
)

var (
	inventoryController inventorycontroller.InventoryController
)

func inventoryAdminRole(g *echo.Group) {
	// g.GET("", inventoryController.GetProductStocks)
	// g.GET("/:item_inventory_id", inventoryController.GetProductStock)
	// g.GET("/generate-import-inventory", inventoryController.ExportStocks)
	// g.POST("/update-stocks", inventoryController.ImportStocks)
}

func SetInventoryGroup(g *echo.Group) {
	inventoryRepository := inventoryrepository.InventoryRepository{database.DB}
	inventoryController = inventorycontroller.InventoryController{&inventoryRepository}
	g.GET("/export-stocks", inventoryController.ExportStocks)
	g.POST("/import-stocks", inventoryController.ImportStocks)

	g.GET("/locations/provinces", inventoryController.GetProvinces)
	g.GET("/locations/provinces/:province_id/cities", inventoryController.GetCitiesByProvinceID)
	g.GET("/locations/provinces/cities/:city_id/subdistricts", inventoryController.GetSubsByCityID)
	// adminRoleMiddleware := g.Group("", ddMiddleware.AdminRoleMiddleware)
	// inventoryAdminRole(adminRoleMiddleware)

}
