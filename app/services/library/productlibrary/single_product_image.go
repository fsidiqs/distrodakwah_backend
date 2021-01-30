package productlibrary

type SingleProductImage struct {
	ID   int    `gorm:"primaryKey;autoIncrement;not null"`
	SPID uint   `gorm:"column:SP_id;type:BIGINT;UNSIGNED;NOT NULL" json:"single_product_id"`
	URL  string `gorm:"type:varchar(255);not null" json:"url"`
}

type VariantProductImage struct {
	ID   int    `gorm:"primaryKey;autoIncrement;not null"`
	VPID uint   `gorm:"column:VP_id;type:BIGINT;UNSIGNED;NOT NULL" json:"variant_product_id"`
	URL  string `gorm:"type:varchar(255);not null" json:"url"`
}
