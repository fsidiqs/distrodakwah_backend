package repository

import (
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/product/model"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/product/request"
)

func (r *ProductRepository) FetchAllBrand() ([]*model.Brand, error) {
	brands := []*model.Brand{}
	err := r.DB.Model(&model.Brand{}).Find(&brands).Error

	if err != nil {
		return nil, err
	}

	return brands, nil
}

func (r *ProductRepository) CreateBrand(req *request.BrandReq) error {
	image := &model.Image{
		Url: "Brand Image Testing",
	}
	tx := r.DB.Begin()
	err := tx.Model(&model.Image{}).Create(&image).Error

	if err != nil {
		tx.Rollback()
		return nil
	}

	brand := &model.Brand{
		ImageID:      image.ID,
		Name:         req.Name,
		UserVendorID: req.UserVendorID,
	}

	err = tx.Model(&model.Brand{}).Create(&brand).Error
	if err != nil {
		tx.Rollback()
		return nil
	}
	return tx.Commit().Error
}
