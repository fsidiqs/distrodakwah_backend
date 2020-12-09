package router

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/database"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/middleware"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/controller/authcontroller"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/repository/authrepository"
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
