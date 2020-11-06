package model

type VariantProductsPrices struct {
	VariantProductID uint64  `gorm:"type:BIGINT;UNSIGNED;NOT NULL" json:"variant_product_id"`
	Name             string  `gorm:"type:varchar(255);not null" json:"name"`
	Value            float64 `gorm:"type:decimal(19,2);not null;default:0.0" json:"value"`
}

type VariantProductsPricesWithParent struct {
	VariantProductID uint64                    `gorm:"type:BIGINT;UNSIGNED;NOT NULL" json:"variant_product_id"`
	VariantProduct   *VariantProductWithParent `gorm:"foreignKey:VariantProductID" json:"variant_product,omitempty"`
	Name             string                    `gorm:"type:varchar(255);not null" json:"name"`
	Value            float64                   `gorm:"type:decimal(19,2);not null;default:0.0" json:"value"`
}

func (VariantProductsPricesWithParent) TableName() string {
	return "variant_products_prices"
}

type VariantProductsPricesArrWithParent []*VariantProductsPricesWithParent
