package router

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	product "distrodakwah_backend/app/router/product"
)

func Init() *echo.Echo {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	// productGroup.Use(appMid.CheckAuthMiddleware, appMid.AdminRoleMiddleware)
	authGroup := e.Group("/auth")
	AuthGroup(authGroup)

	userGroup := e.Group("/users")
	SetUserGroup(userGroup)

	productGroup := e.Group("/products")
	product.SetProductGroup(productGroup)

	InventoryGroup := e.Group("/inventories")
	SetInventoryGroup(InventoryGroup)

	orderGroup := e.Group("/orders")
	SetOrderGroup(orderGroup)

	return e
}
