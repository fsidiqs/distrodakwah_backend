package inventorymodel

import "time"

type VPIInventoryAdjustment struct {
	ID             uint      `gorm:"primaryKey;autoIncrement;not null"`
	VPIInventoryID uint      `gorm:"column:VPI_inventory_id;type:BIGINT;UNSIGNED;NOT NULL" json:"item_inventory_id"`
	UserID         int       `gorm:"type:BIGINT;UNSIGNED;NOT NULL" json:"user_id"`
	StockBefore    int       `gorm:"type:INT;NOT NULL" json:"stock_before"`
	StockAfter     int       `gorm:"type:INT;NOT NULL" json:"stock_after"`
	CreatedAt      time.Time `json:"created_at"`
}

func (VPIInventoryAdjustment) TableName() string {
	return "VPI_inventory_adjustments"
}
