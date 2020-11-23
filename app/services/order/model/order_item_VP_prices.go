package model

type OrderItemVPPrice struct {
	ID       uint64  `gorm:"primaryKey;autoIncrement;not null"`
	Oitsp_ID uint64  `gorm:"type:BIGINT;UNSIGNED;NOT NULL" json:"oitsp_id"`
	Name     string  `gorm:"type:varchar(255);not null" json:"name"`
	Value    float64 `gorm:"type:decimal(19,2);not null;default:0.0" json:"value"`
}

func (OrderItemVPPrice) TableName() string {
	return "order_item_VP_prices"
}

type OrderItemVPPriceArr []*OrderItemVPPrice

func (vpPrices OrderItemVPPriceArr) GetOrderItemPriceByName(priceName string) (float64, bool) {
	for _, vpPrice := range vpPrices {
		if vpPrice.Name == priceName {
			return vpPrice.Value, true
		}
	}
	return 0, false

}
