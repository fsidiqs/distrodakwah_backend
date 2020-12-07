package userhandler

import (
	"errors"

	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/class/userclass"
)

type UserReq struct {
	Name      string                  `json:"name"`
	Email     string                  `json:"email"`
	Password  string                  `json:"password"`
	Phone     string                  `json:"phone"`
	RoleID    uint8                   `json:"role_id"`
	Gender    uint8                   `json:"gender"`
	BirthDate userclass.UserBirthDate `json:"birthdate"`
}

const (
	EmptyFieldErr = "all fields must not be empty"
)

func (u UserReq) Validate() error {
	if len(u.Name) == 0 || len(u.Password) == 0 || len(u.Email) == 0 || len(u.Phone) == 0 || u.Gender == 0 {
		return errors.New(EmptyFieldErr)
	}

	return nil
}
