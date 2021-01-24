package productmodel

import (
	"time"

	"gorm.io/gorm"
)

type ProductStatus uint8

const (
	ProductKindSingle  = 1
	ProductKindVariant = 2
)

const (
	ProductTypeConsignment = 1
	ProductTypeVendor      = 2
)

// Product is base product table
type Product struct {
	ID            uint           `gorm:"primaryKey;autoIncrement;not null"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	BrandID       uint           `json:"brand_id"`
	CategoryID    uint           `json:"category_id"`
	ProductTypeID uint8          `gorm:"type:INT;UNSIGNED;NOT NULL" json:"product_type_id"`
	ProductKindID uint8          `gorm:"type:TINYINT;UNSIGNED;NOT NULL" json:"product_kind_id"`
	Name          string         `gorm:"type:varchar(255);not null" json:"name"`
	Description   string         `gorm:"type:text;not null" json:"description"`
	Status        string         `gorm:"type:NOT NULL;default:0" json:"status"`
	ProductImages []ProductImage `gorm:"foreignKey:ProductID;references:ID" json:"product_images"`
	// Items         []Item         `gorm:"foreignKey:ProductID;references:ID" json:"prices"`
	SingleProductItem   *SingleProductItem      `gorm:"foreignKey:ProductID;references:ID" json:"single_product_item"`
	VariantProductItems []VariantProductItem    `gorm:"foreignKey:ProductID;references:ID" json:"variant_product_item"`
	VariantProducts     []VariantProductVariant `gorm:foreignKey:ProductID;references:ID" json:"variant_products"`
}

type ProductResponse struct {
	ID              int                     `gorm:"primaryKey;autoIncrement;not null"`
	CreatedAt       time.Time               `json:"created_at"`
	UpdatedAt       time.Time               `json:"updated_at"`
	DeletedAt       gorm.DeletedAt          `gorm:"index" json:"deleted_at"`
	BrandID         uint                    `json:"brand_id"`
	Brand           *Brand                  `gorm:"foreignKey:BrandID" json:"brand,omitempty"`
	CategoryID      uint                    `json:"category_id"`
	Category        *Category               `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
	ProductTypeID   uint8                   `gorm:"type:INT;UNSIGNED;NOT NULL" json:"product_type_id"`
	ProductType     *ProductType            `gorm:"foreignKey:ProductTypeID;references:ID" json:"product_type,omitempty"`
	ProductKindID   uint8                   `gorm:"type:TINYINT;UNSIGNED;NOT NULL" json:"product_kind_id"`
	ProductKind     *ProductKind            `gorm:"foreignKey:ProductKindID;references:ID" json:"product_kind,omitempty"`
	Name            string                  `gorm:"type:varchar(255);not null" json:"name"`
	Description     string                  `gorm:"type:text;not null" json:"description"`
	Status          ProductStatus           `gorm:"type:TINYINT;UNSIGNED;NOT NULL;default:0" json:"status"`
	ProductImages   []ProductsProductImage  `gorm:"foreignKey:ProductID" json:"product_images"`
	Items           []Item                  `gorm:"foreignKey:ProductID" json:"items,omitempty"`
	VariantProducts []VariantProductVariant `gorm:"foreignKey:ProductID" json:"variant_products,omitempty"`
}

func (ProductResponse) TableName() string {
	return "products"
}

// ProductFromRequest.ProductDetail Contains only Harga Jual

type ProductSimpleInfo struct {
	ID            int            `gorm:"primaryKey;autoIncrement;not null"`
	ProductKindID uint8          `json:"product_kind_id"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (ProductSimpleInfo) TableName() string {
	return "products"
}

// type ProductInventory struct {
// 	ID            int              `gorm:"primaryKey;autoIncrement;not null"`
// 	Sku           string              `json:"sku"`
// 	ProductKindID uint8               `json:"product_kind_id"`
// 	SingleProduct *SingleProductStock `gorm:"foreignKey:ProductID;references:ID" json:"single_product,omitempty"`

// 	VariantProducts []*VariantProductStock `gorm:"foreignKey:ProductID;references:ID" json:"variant_product,omitempty"`
// }

// func (ProductInventory) TableName() string {
// 	return "products"
// }

// type ProductInventoryForFetch struct {
// 	ID            int              `gorm:"primaryKey;autoIncrement;not null"`
// 	Sku           string              `json:"sku"`
// 	ProductKindID uint8               `json:"product_kind_id"`
// 	SingleProduct *SingleProductStock `gorm:"foreignKey:ProductID;references:ID" json:"single_product,omitempty"`

// 	VariantProduct *VariantProductStock `gorm:"foreignKey:ProductID;references:ID" json:"variant_product,omitempty"`
// }

// func (ProductInventoryForFetch) TableName() string {
// 	return "products"
// }
