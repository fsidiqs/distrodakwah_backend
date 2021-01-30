package productmodel

type SPItemPrice struct {
	ID       uint   `gorm:"primaryKey;autoIncrement;not null"`
	SPItemID uint   `gorm:"column:SP_item_id" json:"sp_item_id"`
	Name     string `gorm:"type:varchar(255);not null" json:"name"`
	Value    int    `json:"value"`
}

func (SPItemPrice) TableName() string {
	return "SP_item_prices"
}
