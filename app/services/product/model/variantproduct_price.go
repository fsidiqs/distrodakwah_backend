package model

type VariantProductsPrice struct {
	ID               uint64  `gorm:"primaryKey;autoIncrement;not null"`
	VariantProductID uint64  `gorm:"type:BIGINT;UNSIGNED;NOT NULL" json:"variant_product_id"`
	Name             string  `gorm:"type:varchar(255);not null" json:"name"`
	Value            float64 `gorm:"type:decimal(19,2);not null;default:0.0" json:"value"`
}

type VariantProductsPriceArr []*VariantProductsPrice

func (vp VariantProductsPriceArr) GetVPPriceByName(priceName string) (float64, bool) {
	for i := range vp {
		if vp[i].Name == priceName {
			return vp[i].Value, true
		}
	}
	return 0, false

}

type VariantProductsPriceWithParent struct {
	ID               uint64                    `gorm:"primaryKey;autoIncrement;not null"`
	VariantProductID uint64                    `gorm:"type:BIGINT;UNSIGNED;NOT NULL" json:"variant_product_id"`
	VariantProduct   *VariantProductWithParent `gorm:"foreignKey:VariantProductID" json:"variant_product,omitempty"`
	Name             string                    `gorm:"type:varchar(255);not null" json:"name"`
	Value            float64                   `gorm:"type:decimal(19,2);not null;default:0.0" json:"value"`
}

type VariantProductsPriceArrWithParent []*VariantProductsPriceWithParent

func (VariantProductsPriceWithParent) TableName() string {
	return "variant_products_prices"
}
