package inventorymodel

type Province struct {
	ID   int    `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	Name string `json:"name"`
}

func (Province) TableName() string {
	return "tb_ro_provinces"
}
