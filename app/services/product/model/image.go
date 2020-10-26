package model

type Image struct {
	ID  uint   `gorm:"primaryKey;autoIncrement;notNull"`
	Url string `gorm:"type:varchar(255);not null" json:"url"`
}
