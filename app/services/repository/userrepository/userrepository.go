package userrepository

import (
	"distrodakwah_backend/app/helper/pagination"
	"distrodakwah_backend/app/services/handler/userhandler"
	"distrodakwah_backend/app/services/model/usermodel"
	"fmt"
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

func (r *Repository) CreateUserReseller(userReseller usermodel.UserWithChild) (*usermodel.UserWithChild, error) {
	var err error
	err = r.DB.Model(&usermodel.UserWithChild{}).Create(&userReseller).Error
	if err != nil {
		return nil, err
	}
	return &userReseller, nil
}

// GET ALL USER VENDOR
func (r *Repository) GetUserVendor(userVendor usermodel.UserWithChild) (*usermodel.UserWithChild, error) {

	err := r.DB.Model(&usermodel.UserWithChild{}).Preload("UserVendor").Find(&userVendor).Error
	if err != nil {
		return nil, err
	}

	return &userVendor, nil

}

// GET ALL USER RESELLER
func (r *Repository) GetUserReseller(userReseller []usermodel.UserReseller) (*[]usermodel.UserReseller, error) {

	err := r.DB.Model(&usermodel.UserReseller{}).Find(&userReseller).Error

	if err != nil {
		return nil, err
	}

	return &userReseller, nil
}

// GET ALL USERS
// func (r *Repository) GetUsers(user []usermodel.User) (*[]usermodel.User, error) {

// 	err := r.DB.Model(&usermodel.User{}).Find(&user).Error

// 	if err != nil {
// 		return nil, err
// 	}
// 	return &user, nil

// }
func (r *Repository) GetUsers(req *userhandler.FetchUsersReq) (*Pagination, error) {

	res := &Pagination{Metadata: pagination.Metadata{}}
	products := []*usermodel.User{}
	query := r.DB.Model(&usermodel.User{}).Unscoped()

	// if it has user_id_arr
	if len(req.UserIDarr) > 0 {
		query = query.Where("users.id IN (?)", req.UserIDarr)
	}
	var total int64

	err := query.Count(&total).Error

	if err != nil {
		fmt.Printf("Error counting Total : %v", err)
		return nil, err
	}
	// build metadata total
	res.Metadata.UpdateTotal(total)

	// query
	res.paginate(req.Metadata)
	// if req.Preload != nil { // check whether slice is empty
	// 	HandlePreload(query, &req.Preload)
	// }
	err = query.
		Offset(res.Metadata.Offset).
		Limit(res.Metadata.Limit).
		Find(&products).Error

	res.UpdateElements(products)

	if err != nil {
		fmt.Printf("Error fetching users\n")
		return nil, err
	}

	// build result

	return res, nil
}

// pagination
func (p *Pagination) paginate(m pagination.Metadata) {
	page, limit, offset := pagination.BuildPagination(m)
	p.Metadata = pagination.Metadata{
		Total:  p.Metadata.Total,
		Limit:  limit,
		Offset: offset,
		Page:   page,
		Pages:  pagination.BuildPages(p.Metadata.Total, limit),
	}

}

// SOFT DELETE USERS

func (r *Repository) DeleteUser(id int) error {

	// err := r.DB.Model(&usermodel.User{}).Where("id=?", id).Update("status", "I").Error
	err := r.DB.Delete(&usermodel.User{}, id).Error

	if err != nil {
		return err
	}

	return nil
}

// upgrade reseller to ekslusif
func (r *Repository) UpgradeAkun(id int) error {
	// err := r.DB.Model(&usermodel.UserWithChild{}).Preload("UserReseller").Where("id=?", id).Update("reseller_role_id", "2").Error
	err := r.DB.Model(&usermodel.UserReseller{}).Where("id=?", id).Update("reseller_role_id", "2").Error
	if err != nil {
		return err
	}
	return nil

}
