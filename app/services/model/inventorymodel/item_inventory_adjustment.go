package inventorymodel

import "time"

type ItemInventoryAdjustment struct {
	ID              uint64    `gorm:"primaryKey;autoIncrement;not null"`
	ItemInventoryID uint64    `gorm:"type:BIGINT;UNSIGNED;NOT NULL" json:"item_inventory_id"`
	UserID          uint64    `gorm:"type:BIGINT;UNSIGNED;NOT NULL" json:"user_id"`
	StockBefore     int       `gorm:"type:INT;NOT NULL" json:"stock_before"`
	StockAfter      int       `gorm:"type:INT;NOT NULL" json:"stock_after"`
	CreatedAt       time.Time `json:"created_at"`
}
