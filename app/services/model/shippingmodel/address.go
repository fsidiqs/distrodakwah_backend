package shippingmodel

type Province struct {
	ID   int    `gorm:"primaryKey;autoIncrement;not null"`
	Name string `gorm:"type:varchar(255);not null" json:"name"`
}

func (Province) TableName() string {
	return "tb_ro_provinces"
}

type City struct {
	ID         int    `gorm:"primaryKey;autoIncrement;not null"`
	Name       string `gorm:"type:varchar(255);not null" json:"name"`
	ProvinceID uint8  `gorm:"type:BIGINT;UNSIGNED;NOT NULL" json:"province_id"`
	PostalCode string `gorm:"type:varchar(255);not null" json:"postal_code"`
}

func (City) TableName() string {
	return "tb_ro_cities"
}

type Subdistrict struct {
	ID     int    `gorm:"primaryKey;autoIncrement;not null"`
	CityID uint8  `gorm:"type:BIGINT;UNSIGNED;NOT NULL" json:"city_id"`
	Name   string `gorm:"type:varchar(255);not null" json:"name"`
}

func (Subdistrict) TableName() string {
	return "tb_ro_subdistricts"
}
