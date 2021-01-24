package productmodel

type SPItemPrice struct {
	ID       int     `gorm:"primaryKey;autoIncrement;not null"`
	SPItemID uint    `gorm:"column:SP_item_id"json:"sp_item_id"`
	Name     string  `gorm:"type:varchar(255);not null" json:"name"`
	Value    float64 `gorm:"type:decimal(19,2);not null;default:0.0" json:"value"`
}

func (SPItemPrice) TableName() string {
	return "SP_item_prices"
}
