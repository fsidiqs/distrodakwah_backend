package ordermodel

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID               uint64         `gorm:"primaryKey;autoIncrement;not null"`
	Invoice          string         `json:"invoice"`
	OrderStatusID    uint8          `gorm:"type:TINYINT;UNSIGNED;NOT NULL" json:"order_status_id"`
	UserResellerID   uint64         `json:"user_reseller_id"`
	Total            uint64         `gorm:"type:decimal(19,2);not null;default:0.0" json:"total"`
	GrandTotal       uint64         `gorm:"type:decimal(19,2);not null;default:0.0" json:"grand_total"`
	UniqueCode       uint           `gorm:"type:decimal(5,2);not null;default:0.0" json:"unique_code"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	StatusID1Expires time.Time      `gorm:"column:status_id_1_expires" json:"status_id_1_expires"`
}

type OrderAdditionalInfo struct {
	OrderID     uint64 `json:"order_id"`
	SenderName  string `json:"sender_name"`
	SenderPhone string `json:"sender_phone"`
}

func (OrderAdditionalInfo) TableName() string {
	return "order_additional_info"
}
