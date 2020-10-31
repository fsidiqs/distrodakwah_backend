package model

type SingleProduct struct {
	ID                   uint                    `gorm:"primaryKey;autoIncrement;not null"`
	ProductID            uint64                  `gorm:"type:BIGINT;UNSIGNED;NOT NULL" json:"product_id"`
	Article              string                  `gorm:"type:varchar(255);not null" json:"article"`
	SingleProductsPrices []*SingleProductsPrices `gorm:"foreignKey:SingleProductID;references:ID" json:"prices"`
}
