package model

type SingleProduct struct {
	SingleProductsPrices []*SingleProductsPrices `gorm:"foreignKey:ProductID;references:ID" json:"prices"`
}
