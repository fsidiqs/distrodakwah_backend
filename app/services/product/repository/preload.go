package repository

import "gorm.io/gorm"

type Preload []string

func HandlePreload(query *gorm.DB, pre *Preload) {
	for _, value := range *pre {
		if value == "belongs_to" {
			query = query.Preload("ProductType").
				Preload("Brand").
				Preload("Category").
				Preload("ProductType").
				Preload("ProductKind")
		} else if value == "product_image" {
			query = query.Preload("ProductImages.ProductImage")
		} else if value == "product_detail" {
			query = query.
				Preload("SingleProduct.SingleProductsPrices").
				Preload("VariantProduct.VariantProductsPrices"). //VariantProduct
				Preload("Variants.Option")
		}
	}
}
