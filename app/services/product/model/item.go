package model

type Item struct {
	ID        uint64      `gorm:"primaryKey;autoIncrement;not null"`
	ProductID uint64      `gorm:"type:BIGINT;UNSIGNED;NOT NULL" json:"product_id"`
	Weight    int         `gorm:"type:INT;UNSIGNED;NOT NULL" json:"weight"`
	Sku       string      `gorm:"type:varchar(255);not null" json:"sku"`
	Prices    []ItemPrice `gorm:"foreignKey:ItemID" json:"prices"`
	Options   []Option    `gorm:"foreignKey:ItemID" json:"options"`
}

type ItemForPriceExport struct {
	ID        uint64      `gorm:"primaryKey;autoIncrement;not null"`
	ProductID uint64      `gorm:"type:BIGINT;UNSIGNED;NOT NULL" json:"product_id"`
	Sku       string      `gorm:"type:varchar(255);not null" json:"sku"`
	Prices    []ItemPrice `gorm:"foreignKey:ItemID" json:"prices"`
}

func (ItemForPriceExport) TableName() string {
	return "items"
}
