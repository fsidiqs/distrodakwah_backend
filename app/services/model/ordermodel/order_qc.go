package ordermodel

type OrderQC struct {
	OrderID uint64 `json:"order_id"`
	Bahan   bool   `gorm:"type:tinyint(1);not null" json:"bahan"`
	Desain  bool   `gorm:"type:tinyint(1);not null" json:"desain"`
	Qc      bool   `gorm:"type:tinyint(1);not null" json:"qc"`
	Packing bool   `gorm:"type:tinyint(1);not null" json:"packing"`
	Pickup  bool   `gorm:"type:tinyint(1);not null" json:"pickup"`
	Jurnal  bool   `gorm:"type:tinyint(1);not null" json:"jurnal"`
}

func (OrderQC) TableName() string {
	return "order_qcs"
}
