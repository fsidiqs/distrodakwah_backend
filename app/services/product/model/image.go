package model

type Image struct {
	ID  uint64 `gorm:"primaryKey;autoIncrement;not null"`
	Url string `gorm:"type:varchar(255);not null" json:"url"`
}
