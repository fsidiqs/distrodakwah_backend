package productmodel

import (
	"time"

	"gorm.io/gorm"
)

type VariantProduct struct {
	ID                     uint                    `gorm:"primaryKey;autoIncrement;not null"`
	CreatedAt              time.Time               `json:"created_at"`
	UpdatedAt              time.Time               `json:"updated_at"`
	DeletedAt              gorm.DeletedAt          `gorm:"index" json:"deleted_at"`
	BrandID                uint                    `json:"brand_id"`
	CategoryID             uint                    `json:"category_id"`
	ProductTypeID          uint8                   `gorm:"type:INT;UNSIGNED;NOT NULL" json:"product_type_id"`
	Name                   string                  `gorm:"type:varchar(255);not null" json:"name"`
	Description            string                  `gorm:"type:text;not null" json:"description"`
	Status                 string                  `gorm:"type:NOT NULL;default:0" json:"status"`
	ProductImages          []VariantProductImage   `gorm:"foreignKey:VPID;references:ID" json:"variant_product_images"`
	VariantProductItems    []VariantProductItem    `gorm:"foreignKey:VPID;references:ID" json:"variant_product_items"`
	VariantProductVariants []VariantProductVariant `gorm:"foreignKey:VPID;references:ID" json:"variant_products"`
}
