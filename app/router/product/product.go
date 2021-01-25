package api

import (
	"distrodakwah_backend/app/database"
	"distrodakwah_backend/app/services/controller/productcontroller"
	"distrodakwah_backend/app/services/repository/productrepository"

	"github.com/labstack/echo"
)

var (
	productController productcontroller.ProductController
)

func Init() {
	productRepository := productrepository.ProductRepository{database.DB}
	productController = productcontroller.ProductController{&productRepository}
}

func productAuthGroup(g *echo.Group) {
	// g.GET("", productController.Gets)
	// g.POST("", productController.Post)
	// g.PUT("/:product_id/edit", productController.UpdateProduct)

}

func SetProductGroup(g *echo.Group) {
	// authGroup := g.Group("", middleware.AdminRoleMiddleware)

	g.GET("", productController.GetAllProducts)
	// authGroup := g.Group("", middleware.CheckAuthMiddleware, middleware.AdminRoleMiddleware)
	// productAuthGroup(authGroup)
	g.POST("/create-product-basic-structure", productController.CreateProductBasicStructure)

	// g.POST("", productController.Post)
	// g.PUT("/:product_id/edit", productController.UpdateProduct)
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
