package model

type SingleProductsPrices struct {
	SingleProductID uint64 `gorm:"type:BIGINT;UNSIGNED;NOT NULL" json:"single_product_id"`
	PriceID         uint64 `gorm:"type:BIGINT;UNSIGNED;NOT NULL" json:"price_id"`
	Price           *Price `gorm:"foreignKey:ID;references:PriceID" json:"prices"`
}
