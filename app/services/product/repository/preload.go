package repository

import (
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/product/request"
	"gorm.io/gorm"
)

func HandlePreload(query *gorm.DB, pre *request.Preload) {
	for _, value := range *pre {
		if value == "belongs_to" {
			query = query.Preload("ProductType").
				Preload("Brand").
				Preload("Category.Subdepartment.Department").
				Preload("ProductType").
				Preload("ProductKind")
		} else if value == "product_image" {
			query = query.Preload("ProductImages.ProductImage")
		} else if value == "product_detail" {
			query = query.
				// Preload("Category.")
				Preload("SingleProduct.SingleProductsPrices").
				Preload("VariantProduct.VariantProductsPrices"). //VariantProduct
				Preload("Variants.Option")
		}
	}
}

// func PreloadCategoryHandler(query *gorm.DB, pre *request.Preload) {

// }
// func SelectByColumnName(query *gorm.DB, cols []string){
// 	for _, value :=
// }
