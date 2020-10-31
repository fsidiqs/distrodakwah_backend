package model

import (
	"gorm.io/gorm"
)

type ProductStatus uint8

// Product is base product table
type Product struct {
	gorm.Model
	BrandID       uint                   `gorm:"type:INT;UNSIGNED;NOT NULL" json:"brand_id"`
	CategoryID    uint                   `gorm:"type:INT;UNSIGNED;NOT NULL" json:"category_id"`
	ProductTypeID uint8                  `gorm:"type:INT;UNSIGNED;NOT NULL" json:"product_type_id"`
	ProductType   *ProductType           `gorm:"foreignKey:ProductTypeID;references:ID" json:"product_type,omitempty"`
	Name          string                 `gorm:"type:varchar(255);not null" json:"name"`
	Description   string                 `gorm:"type:text;not null" json:"description"`
	Status        ProductStatus          `gorm:"type:TINYINT;UNSIGNED;NOT NULL;default:0" json:"status"`
	ProductImages []*ProductHasManyImage `gorm:"foreignKey:ProductID;references:ID" json:"product_images"`

	SingleProduct  *SingleProduct    `gorm:"foreignKey:ProductID;references:ID" json:"single_product,omitempty"`
	VariantProduct []*VariantProduct `gorm:"foreignKey:ProductID;references:ID" json:"variant_product,omitempty"`
}

type SaveProduct struct {
	gorm.Model
	BrandID       uint          `gorm:"type:INT;UNSIGNED;NOT NULL" json:"brand_id"`
	CategoryID    uint          `gorm:"type:INT;UNSIGNED;NOT NULL" json:"category_id"`
	ProductTypeID uint8         `gorm:"type:INT;UNSIGNED;NOT NULL" json:"product_type_id"`
	Name          string        `gorm:"type:varchar(255);not null" json:"name"`
	Description   string        `gorm:"type:text;not null" json:"description"`
	Status        ProductStatus `gorm:"type:TINYINT;UNSIGNED;NOT NULL;default:0" json:"status"`
	// ProductImages []*ProductHasManyImage `gorm:"foreignKey:ProductID;references:ID" json:"product_images"`

	// SingleProduct  *SingleProduct    `gorm:"foreignKey:ProductID;references:ID" json:"single_product,omitempty"`
	// VariantProduct []*VariantProduct `gorm:"foreignKey:ProductID;references:ID" json:"variant_product,omitempty"`
}

func (SaveProduct) TableName() string {
	return "products"
}
