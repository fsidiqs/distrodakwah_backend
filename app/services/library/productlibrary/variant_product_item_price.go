package productlibrary

type VPItemPrice struct {
	ID       uint   `gorm:"primaryKey;autoIncrement;not null"`
	VPItemID uint   `gorm:"column:VP_item_id"json:"vp_item_id"`
	Name     string `gorm:"type:varchar(255);not null" json:"name"`
	Value    int    `json:"value"`
}
