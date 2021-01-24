package productmodel

type Image struct {
	ID  int    `gorm:"primaryKey;autoIncrement;not null"`
	Url string `gorm:"type:varchar(255);not null" json:"url"`
}
