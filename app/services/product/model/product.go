package model

import (
	"time"

	"gorm.io/gorm"
)

type ProductStatus uint8

const (
	ProductKindSingle  = 1
	ProductKindVariant = 2
)

// Product is base product table
type Product struct {
	ID            uint64                  `gorm:"primaryKey;autoIncrement;not null"`
	CreatedAt     time.Time               `json:"created_at"`
	UpdatedAt     time.Time               `json:"updated_at"`
	DeletedAt     gorm.DeletedAt          `gorm:"index" json:"deleted_at"`
	BrandID       uint                    `json:"brand_id"`
	CategoryID    uint                    `json:"category_id"`
	ProductTypeID uint8                   `gorm:"type:INT;UNSIGNED;NOT NULL" json:"product_type_id"`
	ProductKindID uint8                   `gorm:"type:TINYINT;UNSIGNED;NOT NULL" json:"product_kind_id"`
	Name          string                  `gorm:"type:varchar(255);not null" json:"name"`
	Description   string                  `gorm:"type:text;not null" json:"description"`
	Status        ProductStatus           `gorm:"type:TINYINT;UNSIGNED;NOT NULL;default:0" json:"status"`
	ProductImages []*ProductsProductImage `gorm:"foreignKey:ProductID;references:ID" json:"product_images"`
	Sku           string                  `gorm:"type:varchar(255);not null" json:"sku"`
}

type ProductResponse struct {
	ID            uint64                  `gorm:"primaryKey;autoIncrement;not null"`
	CreatedAt     time.Time               `json:"created_at"`
	UpdatedAt     time.Time               `json:"updated_at"`
	DeletedAt     gorm.DeletedAt          `gorm:"index" json:"deleted_at"`
	BrandID       uint                    `json:"brand_id"`
	Brand         *Brand                  `gorm:"foreignKey:BrandID" json:"brand,omitempty"`
	CategoryID    uint                    `json:"category_id"`
	Category      *Category               `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
	ProductTypeID uint8                   `gorm:"type:INT;UNSIGNED;NOT NULL" json:"product_type_id"`
	ProductType   *ProductType            `gorm:"foreignKey:ProductTypeID;references:ID" json:"product_type,omitempty"`
	ProductKindID uint8                   `gorm:"type:TINYINT;UNSIGNED;NOT NULL" json:"product_kind_id"`
	ProductKind   *ProductKind            `gorm:"foreignKey:ProductKindID;references:ID" json:"product_kind,omitempty"`
	Name          string                  `gorm:"type:varchar(255);not null" json:"name"`
	Description   string                  `gorm:"type:text;not null" json:"description"`
	Status        ProductStatus           `gorm:"type:TINYINT;UNSIGNED;NOT NULL;default:0" json:"status"`
	ProductImages []*ProductsProductImage `gorm:"foreignKey:ProductID;references:ID" json:"product_images"`
	Sku           string                  `gorm:"type:varchar(255);not null" json:"sku"`

	SingleProduct *SingleProductWithPrices `gorm:"foreignKey:ProductID;references:ID" json:"single_product,omitempty"`

	VariantProduct []*VariantProductFetch `gorm:"foreignKey:ProductID;references:ID" json:"variant_product,omitempty"`
	Variants       []*Variant             `gorm:"foreignKey:ProductID;references:ID" json:"variants,omitempty"`
}

func (ProductResponse) TableName() string {
	return "products"
}

// ProductFromRequest.ProductDetail Contains only Harga Jual
type ProductFromRequestJSON struct {
	BrandID              uint          `json:"brand_id"`
	CategoryID           uint          `json:"category_id"`
	ProductTypeID        uint8         `json:"product_type_id"`
	ProductKindID        uint8         `json:"product_kind_id"`
	Name                 string        `json:"name"`
	Description          string        `json:"description"`
	Status               ProductStatus `json:"status"`
	ProductImages        ProductImages
	SingleProductDetail  string `json:"single_product_detail,omitempty"`
	VariantProductDetail string `json:"variant_product_detail,omitempty"`
	Sku                  string `json:"sku"`
}

type ProductSimpleInfo struct {
	ID            uint64         `gorm:"primaryKey;autoIncrement;not null"`
	ProductKindID uint8          `json:"product_kind_id"`
	Sku           string         `json:"sku"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (ProductSimpleInfo) TableName() string {
	return "products"
}

type ProductInventory struct {
	ID            uint64              `gorm:"primaryKey;autoIncrement;not null"`
	Sku           string              `json:"sku"`
	ProductKindID uint8               `json:"product_kind_id"`
	SingleProduct *SingleProductStock `gorm:"foreignKey:ProductID;references:ID" json:"single_product,omitempty"`

	VariantProducts []*VariantProductStock `gorm:"foreignKey:ProductID;references:ID" json:"variant_product,omitempty"`
}

func (ProductInventory) TableName() string {
	return "products"
}

type ProductInventoryForFetch struct {
	ID            uint64              `gorm:"primaryKey;autoIncrement;not null"`
	Sku           string              `json:"sku"`
	ProductKindID uint8               `json:"product_kind_id"`
	SingleProduct *SingleProductStock `gorm:"foreignKey:ProductID;references:ID" json:"single_product,omitempty"`

	VariantProduct *VariantProductStock `gorm:"foreignKey:ProductID;references:ID" json:"variant_product,omitempty"`
}

func (ProductInventoryForFetch) TableName() string {
	return "products"
}
