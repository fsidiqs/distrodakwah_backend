package model

type Variant struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement;not null"`
	ProductID uint64    `gorm:"not null" json:"product_id"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	Option    []*Option `gorm:"foreignKey:VariantID;references:ID" json:"options,omitempty"`
}

type VariantFetch struct {
	ID          uint64    `gorm:"primaryKey;autoIncrement;not null"`
	FKProductID uint64    `gorm:"column:product_id;not null" json:"product_id"`
	Name        string    `gorm:"type:varchar(255);not null" json:"name"`
	Option      []*Option `gorm:"foreignKey:VariantID;references:ID" json:"options,omitempty"`
}

func (VariantFetch) TableName() string {
	return "variants"
}

type VariantArr []Variant
