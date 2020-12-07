package router

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/database"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/middleware"

var (
	authController user.
)

func InitAuthRoute() {
	authRepository := repository.UserRepository{database.DB}
	userController = controller.{&authRepository}
}

// func SetAuthMiddlewares(g *echo.Group)

func AuthGroup(g *echo.Group) {
	g.POST("/login", userController.Login)

	g.GET("", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "ok")
	}, middleware.CheckAuthMiddleware, middleware.AdminRoleMiddleware)
}
