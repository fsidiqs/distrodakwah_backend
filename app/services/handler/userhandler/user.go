package userhandler

import (
	"errors"

	"distrodakwah_backend/app/services/class/userclass"
)

type UserReq struct {
	Gender    string              `gorm:"type:enum;UNSIGNED;NOT NULL" json:"gender"`
	Birthdate userclass.Birthdate `gorm:"type:date; column:birthdate" json:"birthdate"`
	Name      string              `gorm:"type:varchar(255);not null" json:"name"`
	Email     string              `gorm:"type:varchar(255);unique;not null" json:"email"`
	Phone     string              `gorm:"type:varchar(255);not null" json:"phone"`
	Password  string              `gorm:"type:varchar(255);not null" json:"password"`
	RoleID    uint8               `gorm:"type:TINYINT;UNSIGNED;NOT NULL;default:0" json:"role_id"`
}

const (
	EmptyFieldErr = "all fields must not be empty"
)

func (u UserReq) Validate() error {
	if len(u.Name) == 0 || len(u.Password) == 0 || len(u.Email) == 0 || len(u.Phone) == 0 || len(u.Gender) == 0 {
		return errors.New(EmptyFieldErr)
	}

	return nil
}

type UserResellerReq struct {
	Name           string              `gorm:"type:varchar(255);not null" json:"name"`
	Email          string              `gorm:"type:varchar(255);unique;not null" json:"email"`
	Phone          string              `gorm:"type:varchar(255);not null" json:"phone"`
	Password       string              `gorm:"type:varchar(255);not null" json:"password"`
	ResellerRoleID uint8               `gorm:"NOT NULL" json:"reseller_role_id"`
	LocationType   string              `json:"location_type"`
	LocationID     int                 `json:"location_id"`
	Address        string              `gorm:"type:text" json:"address"`
	PostalCode     string              `json:"postal_code"`
	Gender         string              `gorm:"type:enum;UNSIGNED;NOT NULL" json:"gender"`
	Birthdate      userclass.Birthdate `gorm:"type:date; column:birthdate" json:"birthdate"`
}
