package repository

import (
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/helper/pagination"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/inventory/model"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/inventory/request"
)

func (r *InventoryRepository) FetchAll(req *request.FetchAllReq) (*pagination.Pagination, error) {
	res := &pagination.Pagination{Metadata: &pagination.Metadata{}}
	inventories := []*model.InventoryResponse{}
	query := r.DB.Model(&model.InventoryResponse{}).Where(
		"INNER JOIN single_products ON single_products.id = sp_inventory"
	)

	if len(req.SPIDArr > 0 ){
		query = query.Where("sp_inventory.id IN (?)", req.SPIDArr)
	}

	if len(req.VPIDArr > 0) {
		query = query.Where("vp_inventory.id IN (?)", req.VPIDArr)
	}


}
