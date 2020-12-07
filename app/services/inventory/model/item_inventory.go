package model

var InitItemInventory = ItemInventory{
	Stock: 0,
	Keep:  0,
}

type ItemInventory struct {
	ID                  uint64               `gorm:"primaryKey;autoIncrement;not null"`
	ItemID              uint64               `gorm:"type:BIGINT;UNSIGNED;NOT NULL" json:"item_id"`
	Stock               int                  `gorm:"type:INT;NOT NULL" json:"stock"`
	Keep                int                  `gorm:"type:INT;NOT NULL" json:"keep"`
	ItemInventoryDetail *ItemInventoryDetail `gorm:"foreignKey:ItemInventoryID" json:"item_inventory_detail"`
}
