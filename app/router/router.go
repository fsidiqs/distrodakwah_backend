package router

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	product "github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/router/product"
)

func Init() *echo.Echo {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	// productGroup.Use(appMid.CheckAuthMiddleware, appMid.AdminRoleMiddleware)
	InitAuthRoute()
	authGroup := e.Group("/auth")
	AuthGroup(authGroup)

	productGroup := e.Group("/products")
	product.SetProductGroup(productGroup)

	InventoryGroup := e.Group("/inventory")
	SetInventoryGroup(InventoryGroup)
	return e
}
