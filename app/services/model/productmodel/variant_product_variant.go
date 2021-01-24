package productmodel

type VariantProductVariant struct {
	ID                    uint                   `gorm:"primaryKey;autoIncrement;not null"`
	ProductID             uint                   `gorm:"type:BIGINT;UNSIGNED;NOT NULL" json:"product_id"`
	Name                  string                 `gorm:"type:varchar(255);not null" json:"name"`
	VariantProductOptions []VariantProductOption `gorm:"foreignKey:VariantProductVariantID;references:ID" json:"variant_product_options,omitempty"`
}

func (VariantProductVariant) TableName() string {
	return "VP_variants"
}

type VariantFetch struct {
	ID          int                    `gorm:"primaryKey;autoIncrement;not null"`
	FKProductID int                    `gorm:"column:product_id;not null" json:"product_id"`
	Name        string                 `gorm:"type:varchar(255);not null" json:"name"`
	Option      []VariantProductOption `gorm:"foreignKey:VariantProductVariantID;references:ID" json:"options,omitempty"`
}

func (VariantFetch) TableName() string {
	return "variants"
}

type VariantArr []VariantProductVariant
