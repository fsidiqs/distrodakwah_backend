package router

import (
	"distrodakwah_backend/app/database"
	ddMiddleware "distrodakwah_backend/app/middleware"
	"distrodakwah_backend/app/services/controller/usercontroller"
	"distrodakwah_backend/app/services/repository/userrepository"

	"github.com/labstack/echo"
)

var (
	userController usercontroller.Controller
)

func userAdminRole(g *echo.Group) {
	g.POST("", userController.CreateUser)
	g.POST("/create-user-reseller", userController.CreateUserReseller)
}

func SetUserGroup(g *echo.Group) {
	userRepository := userrepository.Repository{database.DB}
	userController = usercontroller.Controller{&userRepository}

	adminRoleMiddleware := g.Group("", ddMiddleware.AdminRoleMiddleware)
	userAdminRole(adminRoleMiddleware)

}
