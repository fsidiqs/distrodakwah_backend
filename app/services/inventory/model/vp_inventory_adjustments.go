package model

type VPInventoryAdjustment struct {
	ID            uint64 `gorm:"primaryKey;autoIncrement;not null"`
	VPInventoryID uint64 `gorm:"type:BIGINT;UNSIGNED;NOT NULL" json:"vp_inventory_id"`
	UserID        uint64 `gorm:"type:BIGINT;UNSIGNED;NOT NULL" json:"user_id"`
}

func (VPInventoryAdjustment) TableName() string {
	return "vp_inventory_adjustments"
}
