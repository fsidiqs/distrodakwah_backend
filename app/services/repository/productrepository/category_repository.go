package productrepository

import (
	"distrodakwah_backend/app/services/handler/producthandler"
	"distrodakwah_backend/app/services/model/productmodel"
)

func (r *ProductRepository) FetchAllCategory() ([]*productmodel.Category, error) {
	categories := []*productmodel.Category{}
	err := r.DB.Model(&productmodel.Category{}).Find(&categories).Error

	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (r *ProductRepository) CreateCategory(req *producthandler.CategoryReq) error {
	image := &productmodel.Image{
		Url: "testing",
	}
	tx := r.DB.Begin()
	err := tx.Model(&productmodel.Image{}).Create(&image).Error

	if err != nil {
		tx.Rollback()
		return nil
	}

	category := &productmodel.Category{
		ImageID:         image.ID,
		SubdepartmentID: req.SubdepartmentID,
		ParentID:        req.ParentID,
		Name:            req.Name,
	}

	err = tx.Model(&productmodel.Category{}).Create(&category).Error
	if err != nil {
		tx.Rollback()
		return nil
	}

	return tx.Commit().Error
}

func (r *ProductRepository) FetchAllDepartments() ([]*productmodel.Department, error) {
	departments := []*productmodel.Department{}
	err := r.DB.Model(&productmodel.Department{}).Find(&departments).Error

	if err != nil {
		return nil, err
	}

	return departments, nil
}

func (r *ProductRepository) CreateDepartment(req *producthandler.DepartmentReq) error {
	image := &productmodel.Image{
		Url: "testing department image",
	}
	tx := r.DB.Begin()
	err := tx.Model(&productmodel.Image{}).Create(&image).Error

	if err != nil {
		tx.Rollback()
		return nil
	}

	department := &productmodel.Department{
		ImageID: image.ID,
		Name:    req.Name,
	}

	err = tx.Model(&productmodel.Department{}).Create(&department).Error
	if err != nil {
		tx.Rollback()
		return nil
	}

	return tx.Commit().Error
}

func (r *ProductRepository) FetchAllSubdepartments() ([]*productmodel.Subdepartment, error) {
	subdepartments := []*productmodel.Subdepartment{}
	err := r.DB.Model(&productmodel.Subdepartment{}).Find(&subdepartments).Error

	if err != nil {
		return nil, err
	}

	return subdepartments, nil
}

func (r *ProductRepository) CreateSubdepartment(req *producthandler.SubdepartmentReq) error {
	image := &productmodel.Image{
		Url: "testing subdepartment image",
	}
	tx := r.DB.Begin()
	err := tx.Model(&productmodel.Image{}).Create(&image).Error

	if err != nil {
		tx.Rollback()
		return nil
	}

	subdepartment := &productmodel.Subdepartment{
		ImageID:      image.ID,
		DepartmentID: req.DepartmentID,
		Name:         req.Name,
	}

	err = tx.Model(&productmodel.Subdepartment{}).Create(&subdepartment).Error
	if err != nil {
		tx.Rollback()
		return nil
	}

	return tx.Commit().Error
}
