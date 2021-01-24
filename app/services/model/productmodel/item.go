package productmodel

type Item struct {
	ID        uint     `gorm:"primaryKey;autoIncrement;not null"`
	ProductID uint     `gorm:"type:BIGINT;UNSIGNED;NOT NULL" json:"product_id"`
	Product   *Product `gorm:"foreignKey:ProductID" json:"product,omitempty"`
	// ItemInventory []inventorymodel.ItemInventory `gorm:"foreignKey:itemable_id" json:"item"`
	Weight int           `gorm:"type:INT;UNSIGNED;NOT NULL" json:"weight"`
	Sku    string        `gorm:"type:varchar(255);not null" json:"sku"`
	Prices []SPItemPrice `gorm:"foreignKey:ItemID" json:"prices"`
	// Options []Option    `gorm:"foreignKey:ItemID" json:"options"`
}

type ItemForPriceExport struct {
	ID        int           `gorm:"primaryKey;autoIncrement;not null"`
	ProductID int           `gorm:"type:BIGINT;UNSIGNED;NOT NULL" json:"product_id"`
	Sku       string        `gorm:"type:varchar(255);not null" json:"sku"`
	Prices    []SPItemPrice `gorm:"foreignKey:ItemID" json:"prices"`
}

func (ItemForPriceExport) TableName() string {
	return "items"
}
