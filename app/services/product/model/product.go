package model

import "gorm.io/gorm"

type ProductStatus uint8

// Product is base product table
type Product struct {
	gorm.Model
	Name           string           `gorm:"type:varchar(255);not null" json:"name"`
	ProductImageID uint             `gorm:"type:BIGINT" json:"product_image_id"`
	Status         ProductStatus    `gorm:"char(1);default:0" json:"status"`
	Variants       []*VariantOption `gorm:"foreignKey:ProductID" json:"variants,omitempty"`
	SkuValues      []*SkuValuePrice `gorm:"foreignKey:ProductID" json:"sku_values"`
}

type ProductVariantOptionSkuValuePrice struct {
	*Product
	Variants  []*VariantOption `gorm:"foreignKey:ProductID" json:"variants"`
	SkuValues []*SkuValuePrice `gorm:"foreignKey:ProductID" json:"sku_values"`
}

func (ProductVariantOptionSkuValuePrice) TableName() string {
	return "products"
}
