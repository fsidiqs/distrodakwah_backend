package model

type Option struct {
	ID         uint   `gorm:"primaryKey;autoIncrement;not null"`
	VariantID  uint   `gorm:"not null" json:"variant_id"`
	SkuValueID string `gorm:"type:varchar(255);not null" json:"sku_value_id"`
	Name       string `gorm:"type:varchar(255);not null" json:"name"`
}
