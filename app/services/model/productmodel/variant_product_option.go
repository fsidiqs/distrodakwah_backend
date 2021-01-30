package productmodel

type VariantProductOption struct {
	ID          uint   `gorm:"primaryKey;autoIncrement;not null"`
	VPVariantID uint   `gorm:"column:VP_variant_id;type:BIGINT;UNSIGNED;not null" json:"variant_product_variant_id"`
	VPItemID    uint   `gorm:"column:VP_item_id" json:"variant_product_item_id"`
	Name        string `gorm:"type:varchar(255);not null" json:"name"`
}

func (VariantProductOption) TableName() string {
	return "VP_options"
}

type OptionArr []VariantProductOption
