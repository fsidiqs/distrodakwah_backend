package repository

import "github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/product/model"

func (r *ProductRepository) FetchAllKind() ([]*model.ProductKind, error) {
	productKinds := []*model.ProductKind{}

	err := r.DB.Model(&model.ProductType{}).Find(&productKinds).Error
	if err != nil {
		return nil, err
	}

	return productKinds, nil
}
