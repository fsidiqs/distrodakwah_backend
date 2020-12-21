package ordermodel

type OrderStatus struct {
	ID   uint64 `gorm:"primaryKey;autoIncrement;not null"`
	Name string `json:"name"`
}
