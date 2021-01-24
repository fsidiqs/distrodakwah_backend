package ordermodel

type OrderItem struct {
	ID              int    `gorm:"primaryKey;autoIncrement;not null"`
	ItemID          int    `json:"item_id"`
	Qty             int    `json:"qty"`
	UnitWeight      int    `json:"unit_weight"`
	Sku             string `gorm:"type:varchar(255);not null" json:"sku"`
	OrderShippingID int    `json:"order_shipping_id"`
}
