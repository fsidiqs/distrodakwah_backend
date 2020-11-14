package model

type SPInventoryAdjustment struct {
	ID            uint64 `gorm:"primaryKey;autoIncrement;not null"`
	SPInventoryID uint64 `gorm:"type:BIGINT;UNSIGNED;NOT NULL" json:"sp_inventory_id"`
	UserID        uint64 `gorm:"type:BIGINT;UNSIGNED;NOT NULL" json:"user_id"`
}

func (SPInventoryAdjustment) TableName() string {
	return "sp_inventory_adjustments"
}
