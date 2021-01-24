package productmodel

type Subdepartment struct {
	ID           uint8       `gorm:"primaryKey;autoIncrement;not null"`
	DepartmentID uint8       `gorm:"tinyint;not null" json:"department_id"`
	Department   *Department `gorm:"foreignKey:DepartmentID" json:"department,omitempty"`
	Name         string      `gorm:"type:varchar(255);not null" json:"name"`
	ImageID      int         `gorm:"type:bigint" json:"image_id"`
}
