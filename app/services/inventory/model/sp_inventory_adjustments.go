package model

import "time"

type SPInventoryAdjustment struct {
	ID            uint64    `gorm:"primaryKey;autoIncrement;not null"`
	SPInventoryID uint64    `gorm:"type:BIGINT;UNSIGNED;NOT NULL" json:"sp_inventory_id"`
	UserID        uint64    `gorm:"type:BIGINT;UNSIGNED;NOT NULL" json:"user_id"`
	StockBefore   int       `gorm:"type:INT;NOT NULL" json:"stock_before"`
	StockAfter    int       `gorm:"type:INT;NOT NULL" json:"stock_after"`
	CreatedAt     time.Time `json:"created_at"`
}

func (SPInventoryAdjustment) TableName() string {
	return "sp_inventory_adjustments"
}
