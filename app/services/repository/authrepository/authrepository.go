package authrepository

import "github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/model/usermodel"

func (r *Repository) GetUserByEmail(email string) (*usermodel.User, error) {
	user := &usermodel.User{}
	var err error

	err = r.DB.Model(&usermodel.User{}).Unscoped().Where("email = ?", email).Find(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}
