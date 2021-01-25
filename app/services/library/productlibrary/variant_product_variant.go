package productlibrary

type VPVariant struct {
	ID        uint   `gorm:"primaryKey;autoIncrement;not null"`
	ProductID uint   `gorm:"type:BIGINT;UNSIGNED;NOT NULL" json:"product_id"`
	Name      string `json:"name"`
}
