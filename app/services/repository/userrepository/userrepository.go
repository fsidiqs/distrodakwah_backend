package userrepository

import (
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/model/usermodel"
)

func (r *Repository) GetUserByEmail(email string) (*usermodel.User, error) {
	user := &usermodel.User{}
	var err error

	err = r.DB.Model(&usermodel.User{}).Unscoped().Where("email = ?", email).Find(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *Repository) CreateUser(user usermodel.User) (*usermodel.User, error) {
	var err error
	err = r.DB.Model(&usermodel.User{}).Create(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *Repository) CreateUserReseller(userReseller usermodel.UserHasOneReseller) (*usermodel.UserHasOneReseller, error) {
	var err error
	err = r.DB.Model(&usermodel.UserHasOneReseller{}).Create(&userReseller).Error
	if err != nil {
		return nil, err
	}
	return &userReseller, nil
}
