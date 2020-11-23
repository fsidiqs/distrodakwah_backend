package model

type SingleProductsPrice struct {
	ID              uint64  `gorm:"primaryKey;autoIncrement;not null"`
	SingleProductID uint64  `gorm:"type:BIGINT;UNSIGNED;NOT NULL" json:"single_product_id"`
	Name            string  `gorm:"type:varchar(255);not null" json:"name"`
	Value           float64 `gorm:"type:decimal(19,2);not null;default:0.0" json:"value"`
}

type SingleProductsPriceArr []*SingleProductsPrice

func (SingleProductsPriceArr) TableName() string {
	return "single_products_prices"
}

func (sp SingleProductsPriceArr) GetSPPriceByName(priceName string) (float64, bool) {
	for i := range sp {
		if sp[i].Name == priceName {
			return sp[i].Value, true
		}
	}
	return 0, false

}

type SingleProductsPriceWithParent struct {
	ID              uint64                   `gorm:"primaryKey;autoIncrement;not null"`
	SingleProductID uint64                   `gorm:"type:BIGINT;UNSIGNED;NOT NULL" json:"single_product_id"`
	SingleProduct   *SingleProductWithParent `gorm:"foreignKey:SingleProductID" json:"single_product,omitempty"`
	Name            string                   `gorm:"type:varchar(255);not null" json:"name"`
	Value           float64                  `gorm:"type:decimal(19,2);not null;default:0.0" json:"value"`
}

type SingleProductsPriceArrWithParent []*SingleProductsPriceWithParent

func (SingleProductsPriceWithParent) TableName() string {
	return "single_products_prices"
}

// func (SingleProductsPricesArrParent) TableName() string {
// 	return "single_products_prices"
// }
