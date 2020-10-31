package model

type ProductImage struct {
	ID  uint   `gorm:"primaryKey;autoIncrement;not null"`
	Url string `gorm:"type:varchar(255);not null" json:"url"`
}
