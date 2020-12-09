package productrepository

import (
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/handler/producthandler"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/model/productmodel"
)

func (r *ProductRepository) FetchAllBrand() ([]*productmodel.Brand, error) {
	brands := []*productmodel.Brand{}
	err := r.DB.Model(&productmodel.Brand{}).Find(&brands).Error

	if err != nil {
		return nil, err
	}

	return brands, nil
}

func (r *ProductRepository) CreateBrand(req *producthandler.BrandReq) error {
	image := &productmodel.Image{
		Url: "Brand Image Testing",
	}
	tx := r.DB.Begin()
	err := tx.Model(&productmodel.Image{}).Create(&image).Error

	if err != nil {
		tx.Rollback()
		return nil
	}

	brand := &productmodel.Brand{
		ImageID:      image.ID,
		Name:         req.Name,
		UserVendorID: req.UserVendorID,
	}

	err = tx.Model(&productmodel.Brand{}).Create(&brand).Error
	if err != nil {
		tx.Rollback()
		return nil
	}
	return tx.Commit().Error
}
