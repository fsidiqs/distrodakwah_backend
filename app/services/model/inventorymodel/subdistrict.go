package inventorymodel

type Subdistrict struct {
	ID     int    `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	CityID int    `json:"city_id"`
	Name   string `json:"name"`
}

func (Subdistrict) TableName() string {
	return "tb_ro_subdistricts"
}
