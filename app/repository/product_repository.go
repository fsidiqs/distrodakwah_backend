package repository

import (
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/helper/pagination"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/product/model"

	"fmt"

	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

// type FetchAllBody struct {
// 	Preload       []string
// 	Metadata      *paginationMetadata `json:"metadata"`
// 	CategoryIDArr []int     `json:"category_id_arr"`
// }

func (r *ProductRepository) FetchAll(m *pagination.Metadata) (*Pagination, error) {
	preloadArr := Preload{"variants", "options", "sku_values", "prices"} // ! udpate

	res := &Pagination{Metadata: &pagination.Metadata{}}
	products := []*model.Product{}
	query := r.DB.Model(&model.Product{}).
		Unscoped()

	var total int64
	err := query.Count(&total).Error
	if err != nil {
		fmt.Printf("Error counting Total : %v", err)
		return nil, err
	}
	// build metadata total
	res.Metadata.UpdateTotal(total)

	// query

	res.paginate(m)
	HandlePreload(query, &preloadArr)

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
