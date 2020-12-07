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

func adminRoleGroup(g *echo.Group) {
	g.POST("/update-stocks", userController.CreateUser)
}

func SetUserGroup(g *echo.Group) {
	userRepository := userrepository.Repository{database.DB}
	userController = controller.UserController{&userRepository}

	adminRole := g.Group("", ddMiddleware.AdminRoleMiddleware)
	adminRoleGroup(adminRole)

}
