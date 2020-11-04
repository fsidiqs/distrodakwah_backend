package model

type ProductSku struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement;not null"`
	ProductID uint64 `gorm:"type:BIGINT;UNSIGNED;NOT NULL" json:"product_id"`
	Sku       string `gorm:"type:varchar(255);not null" json:"sku"`
}

type ProductSkuSingleProductPrice struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement;not null"`
	ProductID uint64 `gorm:"type:BIGINT;UNSIGNED;NOT NULL" json:"product_id"`
	Sku       string `gorm:"type:varchar(255);not null" json:"sku"`
}

func (ProductSkuSingleProductPrice) TableName() string {
	return "product_skus"
}
