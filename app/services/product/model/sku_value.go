package model

type SkuValue struct {
	ID             uint   `gorm:"primaryKey"`
	ProductID      uint   `gorm:"not null" json:"product_id"`
	ProductImageID uint   `gorm:"type:BIGINT;not null" json:"product_image_id"`
	Sku            string `gorm:"type:varchar(255);not null" json:"sku"`
}

type SkuValuePrice struct {
	ID             uint     `gorm:"primaryKey"`
	ProductID      uint     `gorm:"not null" json:"product_id"`
	ProductImageID uint     `gorm:"type:BIGINT;not null" json:"product_image_id"`
	Sku            string   `gorm:"type:varchar(255);not null" json:"sku"`
	Prices         []*Price `gorm:"foreignKey:SkuValueID" json:"prices"`
}

func (SkuValuePrice) TableName() string {
	return "sku_values"
}
