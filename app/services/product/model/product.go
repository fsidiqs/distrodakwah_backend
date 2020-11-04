package model

import (
	"time"

	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/product/request"
	"gorm.io/gorm"
)

type ProductStatus uint8

// Product is base product table
type Product struct {
	ID                 uint64                  `gorm:"primaryKey;autoIncrement;not null"`
	CreatedAt          time.Time               `json:"created_at"`
	UpdatedAt          time.Time               `json:"updated_at"`
	DeletedAt          gorm.DeletedAt          `gorm:"index" json:"deleted_at"`
	BrandID            uint                    `json:"brand_id"`
	CategoryID         uint                    `json:"category_id"`
	ProductTypeID      uint8                   `gorm:"type:INT;UNSIGNED;NOT NULL" json:"product_type_id"`
	ProductCharacterID uint8                   `gorm:"type:TINYINT;UNSIGNED;NOT NULL" json:"product_character_id"`
	Name               string                  `gorm:"type:varchar(255);not null" json:"name"`
	Description        string                  `gorm:"type:text;not null" json:"description"`
	Status             ProductStatus           `gorm:"type:TINYINT;UNSIGNED;NOT NULL;default:0" json:"status"`
	ProductImages      []*ProductsProductImage `gorm:"foreignKey:ProductID;references:ID" json:"product_images"`
	ProductSku         *ProductSku             `gorm:"foreignKey:ProductID;references:ID" json:"product_sku"`
}

type ProductResponse struct {
	ID                   uint64                  `gorm:"primaryKey;autoIncrement;not null"`
	CreatedAt            time.Time               `json:"created_at"`
	UpdatedAt            time.Time               `json:"updated_at"`
	DeletedAt            gorm.DeletedAt          `gorm:"index" json:"deleted_at"`
	BrandID              uint                    `json:"brand_id"`
	Brand                *Brand                  `gorm:"foreignKey:BrandID" json:"brand"`
	CategoryID           uint                    `json:"category_id"`
	Category             *Category               `gorm:"foreignKey:CategoryID" json:"category"`
	ProductTypeID        uint8                   `gorm:"type:INT;UNSIGNED;NOT NULL" json:"product_type_id"`
	ProductType          *ProductType            `gorm:"foreignKey:ProductTypeID;references:ID" json:"product_type,omitempty"`
	ProductCharacterID   uint8                   `gorm:"type:TINYINT;UNSIGNED;NOT NULL" json:"product_character_id"`
	ProductCharacter     *ProductCharacter       `gorm:"foreignKey:ProductCharacterID;references:ID" json:"product_character,omitempty"`
	Name                 string                  `gorm:"type:varchar(255);not null" json:"name"`
	Description          string                  `gorm:"type:text;not null" json:"description"`
	Status               ProductStatus           `gorm:"type:TINYINT;UNSIGNED;NOT NULL;default:0" json:"status"`
	ProductImages        []*ProductsProductImage `gorm:"foreignKey:ProductID;references:ID" json:"product_images"`
	ProductSku           *ProductSku             `gorm:"foreignKey:ProductID;references:ID" json:"product_sku"`
	SingleProductsPrices []*SingleProductsPrices `gorm:"foreignKey:ProductID;references:ID" json:"single_product,omitempty"`

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
	ProductCharacterID   uint8         `json:"product_character_id"`
	Name                 string        `json:"name"`
	Description          string        `json:"description"`
	Status               ProductStatus `json:"status"`
	ProductImages        string        `json:"product_images"`
	SingleProductDetail  string        `json:"single_product_detail,omitempty"`
	VariantProductDetail string        `json:"variant_product_detail,omitempty"`
	MainSku              string        `json:"main_sku"`
}

type ProductFromRequest struct {
	BrandID                          uint          `json:"brand_id"`
	CategoryID                       uint          `json:"category_id"`
	ProductTypeID                    uint8         `json:"product_type_id"`
	ProductCharacterID               uint8         `json:"product_character_id"`
	Name                             string        `json:"name"`
	Description                      string        `json:"description"`
	Status                           ProductStatus `json:"status"`
	ProductImages                    ProductImages `json:"product_images"`
	MainSku                          string        `json:"main_sku"`
	*request.SingleProductDetailReq  `json:"single_product_detail,omitempty"`
	*request.VariantProductDetailReq `json:"variant_product_detail,omitempty"`
	// SingleProduct  *SingleProduct    `gorm:"foreignKey:ProductID;references:ID" json:"single_product,omitempty"`
	// VariantProduct []*VariantProduct `gorm:"foreignKey:ProductID;references:ID" json:"variant_product,omitempty"`
}

// func (p *SaveProduct) UnmarshalJSON(data []byte) error {
// 	var v map[string]interface{}
// 	if err := json.Unmarshal(data, &v); err != nil {
// 		return err
// 	}

// 	fmt.Printf("fajar %+v\n", v)
// 	return nil
// }
