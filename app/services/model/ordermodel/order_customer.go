package ordermodel

type OrderCustomer struct {
	ID             int    `gorm:"primaryKey;autoIncrement;not null"`
	OrderID        int    `json:"order_id"`
	Name           string `gorm:"type:varchar(255);not null" json:"name"`
	Phone          string `gorm:"type:varchar(255);not null" json:"phone"`
	Email          string `gorm:"type:varchar(255);" json:"email"`
	SubdistrictID  int    `gorm:"type:INT;not null" json:"subdistrict_id"`
	Address        string `gorm:"type:TEXT;not null" json:"address"`
	PostalCode     string `gorm:"type:varchar(255);not null" json:"postal_code"`
	UserResellerID int    `json:"user_reseller_id"`
	CustomerID     int    `json:"customer_id"`
}
