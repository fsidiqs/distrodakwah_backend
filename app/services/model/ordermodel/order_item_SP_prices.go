package ordermodel

type OrderItemSPPrice struct {
	ID       uint64  `gorm:"primaryKey;autoIncrement;not null"`
	Oitsp_ID uint64  `gorm:"type:BIGINT;UNSIGNED;NOT NULL" json:"oitsp_id"`
	Name     string  `gorm:"type:varchar(255);not null" json:"name"`
	Value    float64 `gorm:"type:decimal(19,2);not null;default:0.0" json:"value"`
}

func (OrderItemSPPrice) TableName() string {
	return "order_item_SP_prices"
}

type OrderItemSPPriceArr []*OrderItemSPPrice

func (spPrices OrderItemSPPriceArr) GetOrderItemPriceByName(priceName string) (float64, bool) {
	for _, spPrice := range spPrices {
		if spPrice.Name == priceName {
			return spPrice.Value, true
		}
	}
	return 0, false

}
