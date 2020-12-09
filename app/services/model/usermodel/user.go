package usermodel

import (
	"errors"
	"time"

	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/auth"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/handler/userhandler"
	"gorm.io/gorm"
)

type User struct {
	ID        uint64         `gorm:"primaryKey;autoIncrement;not null"`
	CreatedAt time.Time      `json:"created_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Gender    string         `gorm:"type:enum;UNSIGNED;NOT NULL" json:"gender"`
	Birthdate time.Time      `gorm:"type:date; column:birthdate" json:"birthdate"`
	Name      string         `gorm:"type:varchar(255);not null" json:"name"`
	Email     string         `gorm:"type:varchar(255);unique;not null" json:"email"`
	Phone     string         `gorm:"type:varchar(255);not null" json:"phone"`
	Password  string         `gorm:"type:varchar(255);not null" json:"password"`
	RoleID    uint8          `gorm:"type:TINYINT;UNSIGNED;NOT NULL;default:0" json:"role_id"`
}

// 1 admin
// 10 vendor
//100 Reseller
const (
	EmptyFieldErr      = "all fields must not be empty"
	HashingPasswordErr = "error hashing password"
)

func NewUser(userReq userhandler.UserReq) (*User, error) {
	var err error
	if err := userReq.Validate(); err != nil {

		return nil, err
	}
	bPW, err := auth.Hash(userReq.Password)

	user := User{
		Gender:    userReq.Gender,
		Name:      userReq.Name,
		Email:     userReq.Email,
		Phone:     userReq.Phone,
		Password:  string(bPW),
		RoleID:    0,
		Birthdate: time.Time(userReq.Birthdate),
	}
	if err != nil {
		return nil, errors.New(HashingPasswordErr)
	}

	return &user, nil
}
func (u User) Validate() error {
	if len(u.Name) == 0 || len(u.Gender) == 0 || len(u.Email) == 0 || u.RoleID == 0 {
		return errors.New(EmptyFieldErr)
	}
	return nil
}

type UserHasOneReseller struct {
	ID           uint64         `gorm:"primaryKey;autoIncrement;not null"`
	CreatedAt    time.Time      `json:"created_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Gender       string         `gorm:"type:enum;UNSIGNED;NOT NULL" json:"gender"`
	Birthdate    time.Time      `gorm:"type:date; column:birthdate" json:"birthdate"`
	Name         string         `gorm:"type:varchar(255);not null" json:"name"`
	Email        string         `gorm:"type:varchar(255);unique;not null" json:"email"`
	Phone        string         `gorm:"type:varchar(255);not null" json:"phone"`
	Password     string         `gorm:"type:varchar(255);not null" json:"password"`
	RoleID       uint8          `gorm:"type:TINYINT;UNSIGNED;NOT NULL;default:0" json:"role_id"`
	UserReseller UserReseller   `gorm:"foreignKey:UserID" json:"user_reseller"`
}

func (UserHasOneReseller) TableName() string {
	return "users"
}
