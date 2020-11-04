package model

type VariantProduct struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement;not null"`
	ProductID uint64 `gorm:"type:BIGINT;UNSIGNED;NOT NULL" json:"product_id"`
	Sku       string `gorm:"type:varchar(255);not null" json:"sku"`
}

type VariantProductFetch struct {
	ID                    uint64                   `gorm:"primaryKey;autoIncrement;not null"`
	ProductID             uint64                   `gorm:"column:product_id;type:BIGINT;UNSIGNED;NOT NULL" json:"product_id"`
	Sku                   string                   `gorm:"type:varchar(255);not null" json:"sku"`
	VariantProductsPrices []*VariantProductsPrices `gorm:"foreignKey:VariantProductID;references:ID" json:"prices"`
}

func (VariantProductFetch) TableName() string {
	return "variant_products"
}
