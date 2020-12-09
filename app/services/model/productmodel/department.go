package productmodel

type Department struct {
	ID      uint8  `gorm:"primaryKey;autoIncrement;not null"`
	Name    string `gorm:"type:varchar(255);not null" json:"name"`
	ImageID uint64 `gorm:"type:bigint" json:"image_id"`
}
