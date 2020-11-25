package router

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/database"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/middleware"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/user/controller"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/user/repository"
)

var (
	userRepository repository.UserRepository
	userController controller.UserController
)

func InitAuthRoute() {
	userRepository = repository.UserRepository{database.DB}
	userController = controller.UserController{&userRepository}
}

// func SetAuthMiddlewares(g *echo.Group)

func AuthGroup(g *echo.Group) {
	g.POST("/login", userController.Login)

	g.GET("", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "ok")
	}, middleware.CheckAuthMiddleware, middleware.AdminRoleMiddleware)
}
