package ordermodel

import (
	"time"

	"gorm.io/gorm"
)

type OrderStatus struct {
	ID   uint8  `json:"id"`
	Name string `json:"name"`
}

type Order struct {
	ID               uint64         `gorm:"primaryKey;autoIncrement;not null"`
	Invoice          string         `json:"invoice"`
	OrderStatusID    uint8          `gorm:"type:TINYINT;UNSIGNED;NOT NULL" json:"order_status_id"`
	UserID           uint64         `json:"user_id"`
	Total            float64        `gorm:"type:decimal(19,2);not null;default:0.0" json:"total"`
	ShippingCost     float64        `gorm:"type:decimal(19,2);not null;default:0.0" json:"shipping_cost"`
	GrandTotal       float64        `gorm:"type:decimal(19,2);not null;default:0.0" json:"grand_total"`
	UniqueCode       float32        `gorm:"type:decimal(5,2);not null;default:0.0" json:"unique_code"`
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

//* exp

type OrderReq struct {
	ID                  uint64         `gorm:"primaryKey;autoIncrement;not null"`
	Invoice             string         `json:"invoice"`
	OrderStatusID       uint8          `gorm:"type:TINYINT;UNSIGNED;NOT NULL" json:"order_status_id"`
	UserID              uint64         `json:"user_id"`
	Total               float64        `gorm:"type:decimal(19,2);not null;default:0.0" json:"total"`
	ShippingCost        float64        `gorm:"type:decimal(19,2);not null;default:0.0" json:"shipping_cost"`
	GrandTotal          float64        `gorm:"type:decimal(19,2);not null;default:0.0" json:"grand_total"`
	UniqueCode          float32        `gorm:"type:decimal(5,2);not null;default:0.0" json:"unique_code"`
	CreatedAt           time.Time      `json:"created_at"`
	UpdatedAt           time.Time      `json:"updated_at"`
	DeletedAt           gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	StatusID1Expires    time.Time      `gorm:"column:status_id_1_expires" json:"status_id_1_expires"`
	OrderItems          *OrderItemReqArr
	OrderShippings      OrderShippings
	OrderCustomerDetail *OrderCustomerDetail `gorm:"foreignKey:OrderID" json:"customer_detail"`
}

type OrderClass struct {
	ID                  uint64         `gorm:"primaryKey;autoIncrement;not null"`
	Invoice             string         `json:"invoice"`
	OrderStatusID       uint8          `gorm:"type:TINYINT;UNSIGNED;NOT NULL" json:"order_status_id"`
	UserID              uint64         `json:"user_id"`
	Total               float64        `gorm:"type:decimal(19,2);not null;default:0.0" json:"total"`
	ShippingCost        float64        `gorm:"type:decimal(19,2);not null;default:0.0" json:"shipping_cost"`
	GrandTotal          float64        `gorm:"type:decimal(19,2);not null;default:0.0" json:"grand_total"`
	UniqueCode          float32        `gorm:"type:decimal(5,2);not null;default:0.0" json:"unique_code"`
	CreatedAt           time.Time      `json:"created_at"`
	UpdatedAt           time.Time      `json:"updated_at"`
	DeletedAt           gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	StatusID1Expires    time.Time      `gorm:"column:status_id_1_expires" json:"status_id_1_expires"`
	OrderItems          []OrderItemI
	OrderShippings      *OrderShippings
	OrderCustomerDetail *OrderCustomerDetail `gorm:"foreignKey:OrderID" json:"customer_detail"`
}

// func (o *OrderClass) PopulateData(orderItemsData OrderItemWithItemIDs) error {
// 	orderItemSPArr := &OrderItemSingleProductArr{}
// 	orderItemVPArr := &OrderItemVariantProductArr{}
// 	for _, oi := range orderItemsData {
// 		if
// 	}
// 	err := o.OrderItemSP.PopulateData()
// 	return nil
// }
