package model

type Customer struct {
	ID            uint64 `gorm:"primaryKey;autoIncrement;not null"`
	Name          string `gorm:"type:varchar(255);not null" json:"name"`
	Phone         string `gorm:"type:varchar(255);not null" json:"phone"`
	Email         string `gorm:"type:varchar(255);" json:"email"`
	AddressDetail string `gorm:"type:TEXT;not null" json:"address_detail"`
	SubdistrictID uint64 `gorm:"type:INT;not null" json:"subdistrict_id"`
	PostalCode    string `gorm:"type:varchar(255);not null" json:"postal_code"`
}
