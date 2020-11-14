package model

import "github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/inventory/model"

type VariantProduct struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement;not null"`
	ProductID uint64 `gorm:"type:BIGINT;UNSIGNED;NOT NULL" json:"product_id"`
	Sku       string `gorm:"type:varchar(255);not null" json:"sku"`
	Weight    int    `gorm:"type:INT;NOT NULL" json:"weight"`
}

type VariantProductFetch struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement;not null"`
	ProductID uint64 `gorm:"column:product_id;type:BIGINT;UNSIGNED;NOT NULL" json:"product_id"`
	Sku       string `gorm:"type:varchar(255);not null" json:"sku"`
	Weight    int    `gorm:"type:INT;NOT NULL" json:"weight"`

	VariantProductsPrices []*VariantProductsPrices `gorm:"foreignKey:VariantProductID;references:ID" json:"prices"`
}

func (VariantProductFetch) TableName() string {
	return "variant_products"
}

type VariantProductWithParent struct {
	ID        uint64   `gorm:"primaryKey;autoIncrement;not null"`
	ProductID uint64   `gorm:"column:product_id;type:BIGINT;UNSIGNED;NOT NULL" json:"product_id"`
	Product   *Product `gorm:"foreignKey:ProductID" json:"product,omitempty"`
	Sku       string   `gorm:"type:varchar(255);not null" json:"sku"`
	Weight    int      `gorm:"type:INT;NOT NULL" json:"weight"`
}

func (VariantProductWithParent) TableName() string {
	return "variant_products"
}

type VariantProductStock struct {
	ID          uint64             `gorm:"primaryKey;autoIncrement;not null"`
	Sku         string             `json:"sku"`
	ProductID   uint64             `gorm:"type:BIGINT;UNSIGNED;NOT NULL" json:"product_id"`
	Product     *ProductSimpleInfo `gorm:"foreignKey:ProductID" json:"product,omitempty"`
	VPInventory *model.VPInventory `gorm:"foreignKey:VariantProductID;references:ID" json:"inventory"`
}

func (VariantProductStock) TableName() string {
	return "variant_products"
}
