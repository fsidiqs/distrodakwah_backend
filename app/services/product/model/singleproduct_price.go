package model

type SingleProductsPrices struct {
	SingleProductID uint64  `gorm:"type:BIGINT;UNSIGNED;NOT NULL" json:"single_product_id"`
	Name            string  `gorm:"type:varchar(255);not null" json:"name"`
	Value           float64 `gorm:"type:decimal(19,2);not null;default:0.0" json:"value"`
}

type SingleProductsPricesWithParent struct {
	SingleProductID uint64                   `gorm:"type:BIGINT;UNSIGNED;NOT NULL" json:"single_product_id"`
	SingleProduct   *SingleProductWithParent `gorm:"foreignKey:SingleProductID" json:"single_product,omitempty"`
	Name            string                   `gorm:"type:varchar(255);not null" json:"name"`
	Value           float64                  `gorm:"type:decimal(19,2);not null;default:0.0" json:"value"`
}

func (SingleProductsPricesWithParent) TableName() string {
	return "single_products_prices"
}

type SingleProductsPricesArrWithParent []*SingleProductsPricesWithParent

// func (SingleProductsPricesArrParent) TableName() string {
// 	return "single_products_prices"
// }
