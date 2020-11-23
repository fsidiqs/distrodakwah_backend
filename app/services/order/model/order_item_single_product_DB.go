package model

import invModel "github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/inventory/model"

type OrderItemSingleProductDB struct {
	ID              uint64                `gorm:"primaryKey;autoIncrement;not null"`
	OrderID         uint64                `json:"order_id"`
	SingleProductID uint64                `json:"single_product_id"`
	SPInventory     *invModel.SPInventory `gorm:"foreignKey:SingleProductID" json:"sp_inventory"`
	Qty             int                   `json:"qty"`
	UnitWeight      int                   `json:"unit_weight"`
	// DropshipperItemPrice float64               `gorm:"type:decimal(19,2);not null;default:0.0" json:"dropshipper_item_price"`
	// RetailItemPrice      float64               `gorm:"type:decimal(19,2);not null;default:0.0" json:"retail_item_price"`
	Prices          *OrderItemSPPriceArr `gorm:"-" json:"prices"`
	OrderShippingID uint64               `json:"order_shipping_id"`
}
