package productmodel

type SingleProductItem struct {
	ID            uint           `gorm:"primaryKey;autoIncrement;not null"`
	SPID          uint           `gorm:"column:SP_id;type:BIGINT;UNSIGNED;NOT NULL" json:"single_product_id"`
	SingleProduct *SingleProduct `gorm:"foreignKey:SPID" json:"single_product,omitempty"`
	Weight        int            `gorm:"type:INT;UNSIGNED;NOT NULL" json:"weight"`
	Sku           string         `gorm:"type:varchar(255);not null" json:"sku"`
	SPIPrices     []SPItemPrice  `gorm:"foreignKey:SPItemID;references:ID" json:"single_product_item_prices"`
}

func (SingleProductItem) TableName() string {
	return "SP_items"
}
