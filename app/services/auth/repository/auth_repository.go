package repository

import (
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/user/model"
)

func (r *AuthRepository) GetUserByEmail(email string) (*model.User, error) {
	user := &model.User{}
	var err error

	err = r.DB.Model(&model.User{}).Unscoped().Where("email = ?", email).Find(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}
