package model

type SPInventory struct {
	ID              uint64 `gorm:"primaryKey;autoIncrement;not null"`
	SingleProductID uint64 `gorm:"type:BIGINT;UNSIGNED;NOT NULL" json:"single_product_id"`
	Stock           int    `gorm:"type:INT;NOT NULL" json:"stock"`
	Keep            int    `gorm:"type:INT;NOT NULL" json:"keep"`
}

func (SPInventory) TableName() string {
	return "sp_inventory"
}
