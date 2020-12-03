package model

type Option struct {
	ID        uint   `gorm:"primaryKey;autoIncrement;not null"`
	VariantID uint64 `gorm:"type:BIGINT;UNSIGNED;not null" json:"variant_id"`
	ItemID    uint64 `json:"item_id"`
	Name      string `gorm:"type:varchar(255);not null" json:"name"`
}

type OptionArr []Option
