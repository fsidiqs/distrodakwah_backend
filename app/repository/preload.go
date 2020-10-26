package repository

import "gorm.io/gorm"

type Preload []string

func HandlePreload(query *gorm.DB, pre *Preload) {
	for _, value := range *pre {
		if value == "variants" {
			query = query.Preload("Variants")
		} else if value == "options" {
			query = query.Preload("Variants.Options")
		} else if value == "sku_values" {
			query = query.Preload("SkuValues")
		} else if value == "prices" {
			query = query.Preload("SkuValues.Prices")
		}
	}
}
