package ordermodel

type OrderItem struct {
	ID              uint64 `gorm:"primaryKey;autoIncrement;not null"`
	ItemID          uint64 `json:"item_id"`
	Qty             uint64 `json:"qty"`
	UnitWeight      int    `json:"unit_weight"`
	Sku             string `gorm:"type:varchar(255);not null" json:"sku"`
	OrderShippingID uint64 `json:"order_shipping_id"`
}
