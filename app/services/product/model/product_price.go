package model

type ItemPrice struct {
	ID     uint64  `gorm:"primaryKey;autoIncrement;not null"`
	ItemID uint64  `json:"item_id"`
	Name   string  `gorm:"type:varchar(255);not null" json:"name"`
	Value  float64 `gorm:"type:decimal(19,2);not null;default:0.0" json:"value"`
}

type ItemPriceArr []ItemPrice
