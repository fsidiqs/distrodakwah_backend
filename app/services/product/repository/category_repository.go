package repository

import (
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/product/model"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/product/request"
)

func (r *ProductRepository) FetchAllCategory() ([]*model.Category, error) {
	categories := []*model.Category{}
	err := r.DB.Model(&model.Category{}).Find(&categories).Error

	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (r *ProductRepository) CreateCategory(req *request.CategoryReq) error {
	image := &model.Image{
		Url: "testing",
	}
	tx := r.DB.Begin()
	err := tx.Model(&model.Image{}).Create(&image).Error

	if err != nil {
		tx.Rollback()
		return nil
	}

	category := &model.Category{
		ImageID:         image.ID,
		SubdepartmentID: req.SubdepartmentID,
		ParentID:        req.ParentID,
		Name:            req.Name,
	}

	err = tx.Model(&model.Category{}).Create(&category).Error
	if err != nil {
		tx.Rollback()
		return nil
	}

	return tx.Commit().Error
}

func (r *ProductRepository) FetchAllDepartments() ([]*model.Department, error) {
	departments := []*model.Department{}
	err := r.DB.Model(&model.Department{}).Find(&departments).Error

	if err != nil {
		return nil, err
	}

	return departments, nil
}

func (r *ProductRepository) CreateDepartment(req *request.DepartmentReq) error {
	image := &model.Image{
		Url: "testing department image",
	}
	tx := r.DB.Begin()
	err := tx.Model(&model.Image{}).Create(&image).Error

	if err != nil {
		tx.Rollback()
		return nil
	}

	department := &model.Department{
		ImageID: image.ID,
		Name:    req.Name,
	}

	err = tx.Model(&model.Department{}).Create(&department).Error
	if err != nil {
		tx.Rollback()
		return nil
	}

	return tx.Commit().Error
}

func (r *ProductRepository) FetchAllSubdepartments() ([]*model.Subdepartment, error) {
	subdepartments := []*model.Subdepartment{}
	err := r.DB.Model(&model.Subdepartment{}).Find(&subdepartments).Error

	if err != nil {
		return nil, err
	}

	return subdepartments, nil
}

func (r *ProductRepository) CreateSubdepartment(req *request.SubdepartmentReq) error {
	image := &model.Image{
		Url: "testing subdepartment image",
	}
	tx := r.DB.Begin()
	err := tx.Model(&model.Image{}).Create(&image).Error

	if err != nil {
		tx.Rollback()
		return nil
	}

	subdepartment := &model.Subdepartment{
		ImageID:      image.ID,
		DepartmentID: req.DepartmentID,
		Name:         req.Name,
	}

	err = tx.Model(&model.Subdepartment{}).Create(&subdepartment).Error
	if err != nil {
		tx.Rollback()
		return nil
	}

	return tx.Commit().Error
}
