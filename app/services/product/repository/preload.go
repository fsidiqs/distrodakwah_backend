package repository

import "gorm.io/gorm"

type Preload []string

func HandlePreload(query *gorm.DB, pre *Preload) {
	for _, value := range *pre {
		if value == "product_type" {
			query = query.Preload("ProductType")
		} else if value == "product_image" {
			query = query.Preload("ProductImages.ProductImage")
		} else if value == "product_detail" {
			query = query.Preload("SingleProduct"). // query for single products
								Preload("VariantProduct.Variant.Option") // query for variant products
		} else if value == "price" {
			query = query.Preload("SingleProduct.SingleProductsPrices.Price").
				Preload("VariantProduct.VariantProductsPrices.Price")
		}
	}
}
