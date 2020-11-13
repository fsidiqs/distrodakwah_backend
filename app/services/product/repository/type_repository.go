package repository

import "github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/product/model"

func (r *ProductRepository) FetchAllType() ([]*model.ProductType, error) {
	productTypes := []*model.ProductType{}

	err := r.DB.Model(&model.ProductType{}).Find(&productTypes).Error
	if err != nil {
		return nil, err
	}

	return productTypes, nil
}
