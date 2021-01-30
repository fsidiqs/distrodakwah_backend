package inventorymodel

type City struct {
	ID         int    `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	ProvinceID int    `json:"province_id"`
	Name       string `json:"name"`
}

func (City) TableName() string {
	return "tb_ro_cities"
}
