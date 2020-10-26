package api

import (
	"github.com/labstack/echo"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/controller"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/database"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/repository"
)

var (
	productRepository repository.ProductRepository
	productController controller.ProductController
)

func Init() {
	productRepository = repository.ProductRepository{database.DB}
	productController = controller.ProductController{&productRepository}
}

func ProductGroup(g *echo.Group) {
	g.GET("", productController.Gets)
}