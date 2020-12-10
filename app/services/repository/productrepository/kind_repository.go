package productrepository

import "distrodakwah_backend/app/services/model/productmodel"

func (r *ProductRepository) FetchAllKind() ([]*productmodel.ProductKind, error) {
	productKinds := []*productmodel.ProductKind{}

	err := r.DB.Model(&productmodel.ProductType{}).Find(&productKinds).Error
	if err != nil {
		return nil, err
	}

	return productKinds, nil
}
