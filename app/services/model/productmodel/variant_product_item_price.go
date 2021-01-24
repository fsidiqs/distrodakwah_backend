package productmodel

type VPItemPrice struct {
	ID       int     `gorm:"primaryKey;autoIncrement;not null"`
	VPItemID uint    `gorm:"column:VP_item_id;not null"json:"vp_item_id"`
	Name     string  `gorm:"type:varchar(255);not null" json:"name"`
	Value    float64 `gorm:"type:decimal(19,2);not null;default:0.0" json:"value"`
}

func (VPItemPrice) TableName() string {
	return "VP_item_prices"
}
