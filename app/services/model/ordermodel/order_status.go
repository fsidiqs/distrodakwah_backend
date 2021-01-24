package ordermodel

type OrderStatus struct {
	ID   int    `gorm:"primaryKey;autoIncrement;not null"`
	Name string `json:"name"`
}
