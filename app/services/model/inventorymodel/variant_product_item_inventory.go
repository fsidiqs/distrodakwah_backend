package inventorymodel

type VPIInventory struct {
	ID                 int                 `gorm:"primaryKey;autoIncrement;not null"`
	VPItemID           uint                `gorm:"column:VP_item_id;type:BIGINT;UNSIGNED;NOT NULL" json:"item_id"`
	Stock              int                 `gorm:"type:INT;NOT NULL" json:"stock"`
	Keep               int                 `gorm:"type:INT;NOT NULL" json:"keep"`
	VPIInventoryDetail *VPIInventoryDetail `gorm:"foreignKey:VPItemInventoryID" json:"item_inventory_detail"`
}

func (VPIInventory) TableName() string {
	return "VP_item_inventories"
}
