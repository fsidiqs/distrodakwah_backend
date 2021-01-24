package productlibrary

type SPItemPrice struct {
	Name  string `gorm:"type:varchar(255);not null" json:"name"`
	Value int    `gorm:"type:decimal(19,2);not null;default:0.0" json:"value"`
}

func NewSPItemRetailPrice() SPItemPrice {
	return SPItemPrice{
		Name:  "retail price",
		Value: 0,
	}
}
