package usermodel

import (
	"errors"
	"time"

	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/auth"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/class/userclass"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/handler/userhandler"
	"gorm.io/gorm"
)

type User struct {
	ID        uint64                  `gorm:"primaryKey;autoIncrement;not null"`
	CreatedAt time.Time               `json:"created_at"`
	DeletedAt gorm.DeletedAt          `gorm:"index" json:"deleted_at"`
	Gender    uint8                   `gorm:"type:TINYINT;UNSIGNED;NOT NULL" json:"gender"`
	Birthdate userclass.UserBirthDate `json:"birthdate"`
	Name      string                  `gorm:"type:varchar(255);not null" json:"name"`
	Email     string                  `gorm:"type:varchar(255);unique;not null" json:"email"`
	Phone     string                  `gorm:"type:varchar(255);not null" json:"phone"`
	Password  string                  `gorm:"type:varchar(255);not null" json:"password"`
	RoleID    uint8                   `gorm:"type:TINYINT;UNSIGNED;NOT NULL;default:0" json:"role_id"`
}

const (
	EmptyFieldErr      = "all fields must not be empty"
	HashingPasswordErr = "error hashing password"
)

func NewUser(userreq userhandler.UserReq) (*User, error) {
	if err := userreq.Validate(); err != nil {
		return nil, err
	}
	hashedPassword, err := auth.Hash(userreq.Password)
	if err != nil {
		return nil, errors.New(HashingPasswordErr)
	}
	user := User{
		Name:      userreq.Name,
		Email:     userreq.Email,
		Phone:     userreq.Phone,
		Password:  string(hashedPassword),
		RoleID:    userreq.RoleID,
		Gender:    userreq.Gender,
		Birthdate: userreq.BirthDate,
	}
	return &user, nil
}
func (u User) Validate() error {
	if len(u.Name) == 0 || u.Gender == 0 || u.Birthdate.IsZero() || len(u.Email) == 0 || u.RoleID == 0 {
		return errors.New(EmptyFieldErr)
	}
	return nil
}
