package api

import (
	"github.com/labstack/echo"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/database"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/middleware"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/controller/productcontroller"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/repository/productrepository"
)

var (
	productController productcontroller.ProductController
)

func Init() {
	productRepository := productrepository.ProductRepository{database.DB}
	productController = productcontroller.ProductController{&productRepository}
}

func productAuthGroup(g *echo.Group) {
	g.GET("", productController.Gets)
	g.POST("", productController.Post)
	g.POST("/create-product-basic-structure", productController.CreateProductBasicStructure)
	g.PUT("/:product_id/edit", productController.UpdateProduct)

}

func SetProductGroup(g *echo.Group) {
	authGroup := g.Group("", middleware.AdminRoleMiddleware)

	// authGroup := g.Group("", middleware.CheckAuthMiddleware, middleware.AdminRoleMiddleware)
	productAuthGroup(authGroup)

	g.GET("/generate-price-template", productController.GeneratePriceTemplate)
	g.POST("/import-prices", productController.ImportPrices)

	g.GET("/brands", productController.GetBrands)
	g.GET("/departments", productController.GetDepartments)
	g.GET("/subdepartments", productController.GetSubdepartments)
	g.GET("/categories", productController.GetCategories)
	g.GET("/product_types", productController.GetProductTypes)
	g.GET("/product_kinds", productController.GetProductKinds)
	// bycolumname
	g.GET("/get-products-by-columns", productController.GetProductsByColumn)
	g.POST("/brands", productController.PostBrand)
	g.POST("/departments", productController.PostDepartment)
	g.POST("/subdepartments", productController.PostSubdepartment)
	g.POST("/categories", productController.PostCategory)
}
