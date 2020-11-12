package api

import (
	"github.com/labstack/echo"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/database"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/product/controller"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/product/repository"
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
	g.POST("", productController.Post)
	g.POST("/create-product-basic-structure", productController.CreateProductBasicStructure)

	g.GET("/generate-price-template", productController.GeneratePriceTemplate)
	g.POST("/import-prices", productController.ImportPrices)

	g.GET("/brands", productController.GetBrands)
	g.GET("/departments", productController.GetDepartments)
	g.GET("/subdepartments", productController.GetSubdepartments)
	g.GET("/categories", productController.GetCategories)
	// bycolumname
	g.GET("/get-products-by-columns", productController.GetProductsByColumn)
	g.POST("/brands", productController.PostBrand)
	g.POST("/departments", productController.PostDepartment)
	g.POST("/subdepartments", productController.PostSubdepartment)
	g.POST("/categories", productController.PostCategory)

	// g.Group("/auth", appMid.CheckAuthMiddleware, appMid.AdminRoleMiddleware)
}
