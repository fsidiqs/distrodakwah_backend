package ordermodel

type OrderItemPrice struct {
	ID          int     `gorm:"primaryKey;autoIncrement;not null"`
	OrderItemID int     `json:"order_item_id"`
	Name        string  `gorm:"type:varchar(255);not null" json:"name"`
	Value       float64 `gorm:"type:decimal(19,2);not null;default:0.0" json:"value"`
}
