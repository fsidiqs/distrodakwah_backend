package model

import "time"

type OrderStatusLeadTime struct {
	OrderID       uint64    `json:"order_id"`
	OrderStatusID uint8     `gorm:"type:TINYINT;UNSIGNED;NOT NULL" json:"order_status_id"`
	CreatedAt     time.Time `json:"created_at"`
}
