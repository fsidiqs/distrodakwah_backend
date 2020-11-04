package model

type SingleProductsPrices struct {
	ProductID uint64 `gorm:"type:BIGINT;UNSIGNED;NOT NULL" json:"product_id"`
	PriceID   uint64 `gorm:"type:BIGINT;UNSIGNED;NOT NULL" json:"price_id"`
	Price     *Price `gorm:"foreignKey:ID;references:PriceID" json:"prices"`
}
