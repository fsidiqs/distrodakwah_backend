package usermodel

import (
	"errors"
	"time"

	"distrodakwah_backend/app/auth"
	"distrodakwah_backend/app/services/handler/userhandler"

	"gorm.io/gorm"
)

type UserReseller struct {
	ID             uint32         `gorm:"primaryKey;autoIncrement;not null"`
	UserID         uint64         `gorm:"type:BIGINT;UNSIGNED;NOT NULL" json:"user_id"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	ResellerRoleID uint8          `gorm:"not null" json:"reesller_role_id"`
	LocationType   string         `json:"location_type"`
	LocationID     int            `gorm:"type:int;not null" json:"location_id"`
	Address        string         `gorm:"type:text;not null" json:"address"`
	PostalCode     string         `json:"postal_code"`
	Status         uint8          `gorm:"type:TINYINT;UNSIGNED;NOT NULL;default:0" json:"status"`
}

func NewUserReseller(userResellerReq userhandler.UserResellerReq) (*UserWithChild, error) {
	var err error

	// validate
	bPW, err := auth.Hash(userResellerReq.Password)
	if err != nil {
		return nil, errors.New(HashingPasswordErr)
	}

	userReseller := UserWithChild{
		Gender:    userResellerReq.Gender,
		Name:      userResellerReq.Name,
		Email:     userResellerReq.Email,
		Phone:     userResellerReq.Phone,
		Password:  string(bPW),
		RoleID:    Reseller,
		Birthdate: time.Time(userResellerReq.Birthdate),
		UserReseller: UserReseller{
			LocationType:   userResellerReq.LocationType,
			LocationID:     userResellerReq.LocationID,
			Address:        userResellerReq.Address,
			PostalCode:     userResellerReq.PostalCode,
			ResellerRoleID: userResellerReq.ResellerRoleID,
			Status:         0,
		},
	}

	return &userReseller, nil
}
