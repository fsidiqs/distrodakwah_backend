package model

type SPInventory struct {
	ID                uint64             `gorm:"primaryKey;autoIncrement;not null"`
	SingleProductID   uint64             `gorm:"type:BIGINT;UNSIGNED;NOT NULL" json:"single_product_id"`
	Stock             int                `gorm:"type:INT;NOT NULL" json:"stock"`
	Keep              int                `gorm:"type:INT;NOT NULL" json:"keep"`
	SPInventoryDetail *SPInventoryDetail `gorm:"foreignKey:SPInventoryID" json:"sp_inventory_detail"`
}

func (SPInventory) TableName() string {
	return "sp_inventory"
}
