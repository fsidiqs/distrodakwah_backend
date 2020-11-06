package model

type ProductKind struct {
	ID   uint8  `gorm:"primaryKey;autoIncrement;not null"`
	Name string `gorm:"type:varchar(255);not null" json:"name"`
}
