package ordermodel

import "time"

type OrderStatusLeadTime struct {
	OrderID       int       `json:"order_id"`
	OrderStatusID uint8     `gorm:"type:TINYINT;UNSIGNED;NOT NULL" json:"order_status_id"`
	CreatedAt     time.Time `json:"created_at"`
}
