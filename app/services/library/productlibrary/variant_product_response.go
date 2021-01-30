package productlibrary

import (
	"database/sql"
	"time"
)

type VariantProductResponse struct {
	ID                     uint                    `gorm:"primaryKey;autoIncrement;not null"`
	CreatedAt              time.Time               `json:"created_at"`
	UpdatedAt              time.Time               `json:"updated_at"`
	DeletedAt              sql.NullTime            `json:"deleted_at"`
	BrandID                uint                    `json:"brand_id"`
	CategoryID             uint                    `json:"category_id"`
	ProductTypeID          uint8                   `json:"product_type_id"`
	Name                   string                  `json:"name"`
	Description            string                  `json:"description"`
	Status                 string                  `json:"status"`
	ProductImages          []VariantProductImage   `gorm:"foreignKey:VPID;references:ID" json:"variant_product_images"`
	VariantProductItems    []VariantProductItem    `gorm:"foreignKey:VPID;references:ID" json:"variant_product_items"`
	VariantProductVariants []VariantProductVariant `gorm:"foreignKey:VPID;references:ID" json:"variant_products_variants"`
	ProductKindID          uint8                   `json:"product_kind_id"`
}
