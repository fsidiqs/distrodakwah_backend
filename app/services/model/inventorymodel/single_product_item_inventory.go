package inventorymodel

const MainSubdistrict = 1

type SPIInventory struct {
	ID                 uint                `gorm:"primaryKey;autoIncrement;not null"`
	SPItemID           uint                `gorm:"column:SP_item_id;type:BIGINT;UNSIGNED;NOT NULL" json:"item_id"`
	Stock              int                 `gorm:"type:INT;NOT NULL" json:"stock"`
	Keep               int                 `gorm:"type:INT;NOT NULL" json:"keep"`
	SPIInventoryDetail *SPIInventoryDetail `gorm:"foreignKey:SPItemInventoryID" json:"item_inventory_detail"`
}

func (SPIInventory) TableName() string {
	return "SP_item_inventories"
}
