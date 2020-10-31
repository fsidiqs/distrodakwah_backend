package model

type VariantProduct struct {
	ID                    uint                     `gorm:"primaryKey;autoIncrement;not null"`
	ProductID             uint64                   `gorm:"type:BIGINT;UNSIGNED;NOT NULL" json:"product_id"`
	Article               string                   `gorm:"type:varchar(255);not null" json:"article"`
	Variant               []*Variant               `gorm:"foreignKey:ProductID" json:"variants"`
	VariantProductsPrices []*VariantProductsPrices `gorm:"foreignKey:VariantProductID;references:ID" json:"prices"`
}
