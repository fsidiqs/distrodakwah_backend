package productlibrary

type VariantProductVariant struct {
	ID                    uint                   `gorm:"primaryKey;autoIncrement;not null"`
	VPID                  uint                   `gorm:"column:VP_id;type:BIGINT;UNSIGNED;NOT NULL" json:"variant_product_id"`
	Name                  string                 `gorm:"type:varchar(255);not null" json:"name"`
	VariantProductOptions []VariantProductOption `gorm:"foreignKey:VPVariantID;references:ID" json:"variant_product_options,omitempty"`
}

func (VariantProductVariant) TableName() string {
	return "VP_variants"
}
