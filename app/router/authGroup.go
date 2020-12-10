package router

import (
	"net/http"

	"distrodakwah_backend/app/database"
	"distrodakwah_backend/app/middleware"
	"distrodakwah_backend/app/services/controller/authcontroller"
	"distrodakwah_backend/app/services/repository/authrepository"

	"github.com/labstack/echo"
)

var (
	authController authcontroller.Controller
)

// func SetAuthMiddlewares(g *echo.Group)

func AuthGroup(g *echo.Group) {
	authRepository := authrepository.Repository{database.DB}
	authController = authcontroller.Controller{&authRepository}

	g.POST("/login", authController.Login)

	g.GET("", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "ok")
	}, middleware.CheckAuthMiddleware, middleware.AdminRoleMiddleware)
}
