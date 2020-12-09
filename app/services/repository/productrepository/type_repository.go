package productrepository

import "github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/model/productmodel"

func (r *ProductRepository) FetchAllType() ([]*productmodel.ProductType, error) {
	productTypes := []*productmodel.ProductType{}

	err := r.DB.Model(&productmodel.ProductType{}).Find(&productTypes).Error
	if err != nil {
		return nil, err
	}

	return productTypes, nil
}
