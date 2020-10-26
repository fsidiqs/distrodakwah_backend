package repository

import (
	"fmt"

	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/helper/pagination"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/product/model"
	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

type FetchAllReq struct {
	Preload
	Metadata     *pagination.Metadata `json:"metadata"`
	ProductIDArr []int                `json:"product_id_arr"`
}

func (r *ProductRepository) FetchAll(req *FetchAllReq) (*Pagination, error) {

	res := &Pagination{Metadata: &pagination.Metadata{}}
	products := []*model.Product{}
	query := r.DB.Model(&model.Product{}).
		Unscoped()
	// if it has product_id_arr
	if req.ProductIDArr != nil {
		query = query.Where("products.id IN (?)", req.ProductIDArr)
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
	if req.Preload != nil { // check wether slice is empty
		HandlePreload(query, &req.Preload)
	}

	err = query.
		Offset(res.Metadata.Offset).
		Limit(res.Metadata.Limit).
		Find(&products).Error

	res.UpdateElements(products)

	if err != nil {
		fmt.Printf("Error fetching products\n")
		return nil, err
	}

	// build result

	return res, nil
}

func (p *Pagination) paginate(m *pagination.Metadata) {

	page, limit, offset := pagination.BuildPagination(m)
	p.Metadata = &pagination.Metadata{
		Total:  p.Metadata.Total,
		Limit:  limit,
		Offset: offset,
		Page:   page,
		Pages:  pagination.BuildPages(p.Metadata.Total, limit),
	}
}