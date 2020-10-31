package model

type Variant struct {
	ID        uint      `gorm:"primaryKey;autoIncrement;not null"`
	ProductID uint      `gorm:"not null" json:"product_id"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	Option    []*Option `gorm:"foreignKey:VariantID" json:"options,omitempty"`
}
