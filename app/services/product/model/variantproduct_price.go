package model

type VariantProductsPrices struct {
	VariantProductID uint64 `gorm:"type:BIGINT;UNSIGNED;NOT NULL" json:"variant_product_id"`
	PriceID          uint64 `gorm:"type:BIGINT;UNSIGNED;NOT NULL" json:"price_id"`
	Price            *Price `gorm:"foreignKey:ID;references:PriceID" json:"prices"`
}
