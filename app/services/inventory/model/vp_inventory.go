package model

type VPInventory struct {
	ID                uint64             `gorm:"primaryKey;autoIncrement;not null"`
	VariantProductID  uint64             `gorm:"type:BIGINT;UNSIGNED;NOT NULL" json:"variant_product_id"`
	Stock             int                `gorm:"type:INT;NOT NULL" json:"stock"`
	Keep              int                `gorm:"type:INT;NOT NULL" json:"keep"`
	VPInventoryDetail *VPInventoryDetail `gorm:"foreignKey:VPInventoryID" json:"vp_inventory_detail"`
}

func (VPInventory) TableName() string {
	return "vp_inventory"
}
