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
				Preload("ProductCharacter")
		} else if value == "product_image" {
			query = query.Preload("ProductImages.ProductImage")
		} else if value == "product_detail" {
			query = query.Preload("ProductSku").
				Preload("SingleProductsPrices.Price").
				Preload("VariantProduct.VariantProductsPrices.Price"). //VariantProduct
				Preload("Variants.Option")
		} else if value == "price" {
			query = query.Preload("SingleProduct").
				Preload("VariantProduct.VariantProductsPrices.Price")
		}
	}
}
