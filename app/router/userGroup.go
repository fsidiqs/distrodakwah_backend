package router

import (
	"github.com/labstack/echo"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/database"
	ddMiddleware "github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/middleware"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/controller/usercontroller"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/repository/userrepository"
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
