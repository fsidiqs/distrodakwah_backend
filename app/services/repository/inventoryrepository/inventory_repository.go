package inventoryrepository

import (
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/helper/httphelper"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/helper/pagination"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/handler/inventoryhandler"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/model/inventorymodel"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/model/productmodel"
)

func (r *InventoryRepository) FetchAll(req inventoryhandler.FetchAllReq) (*pagination.Pagination, error) {
	var err error
	var total int64

	res := &pagination.Pagination{Metadata: pagination.Metadata{}}

	products := []*productmodel.ProductSimpleInfo{}
	productQ := r.DB.Model(&productmodel.ProductSimpleInfo{}).Unscoped()

	if len(req.ProductIDArr) > 0 {
		productQ = productQ.Where("products.id IN (?)", req.ProductIDArr)
	}
	// count
	err = productQ.Count(&total).Error
	if err != nil {
		return nil, err
	}
	res.Metadata.UpdateTotal(total)
	res.Paginate(req.Metadata)

	err = productQ.Offset(res.Metadata.Offset).
		Limit(res.Metadata.Limit).
		Find(&products).Error

	prodIDArr := []uint64{}
	for _, product := range products {
		prodIDArr = append(
			prodIDArr,
			product.ID,
		)
	}

	// STEP singleProducts
	items := []productmodel.Item{}

	err = r.DB.Model(&productmodel.Item{}).
		Where("product_id IN (?)", prodIDArr).
		Preload("Product").Preload("ItemInventory.ItemInventoryDetail.UserVendor").
		Find(&items).Error

	// STEP Of variantProducts
	res.UpdateElements(items)

	return res, nil
}

func (r *InventoryRepository) ExportInventory() ([]productmodel.Item, error) {
	var err error

	products := []*productmodel.ProductSimpleInfo{}
	err = r.DB.Model(&productmodel.ProductSimpleInfo{}).Unscoped().Find(&products).Error

	prodIDArr := []uint64{}
	for _, product := range products {
		prodIDArr = append(prodIDArr, product.ID)
	}

	if err != nil {
		return nil, err
	}

	items := []productmodel.Item{}
	err = r.DB.Model(&productmodel.Item{}).
		Preload("Product").Preload("ItemInventory.ItemInventoryDetail.UserVendor").
		Find(&items).Error
	return items, nil
}

type FindReq struct {
	ItemInventoryID uint64
	Preload         httphelper.Preload
}

func (ir *InventoryRepository) Find(req FindReq) (*inventorymodel.ItemInventory, error) {
	var err error

	itemIntenvory := &inventorymodel.ItemInventory{}

	err = ir.DB.Model(&inventorymodel.ItemInventory{}).
		Preload("ItemInventoryDetail.UserVendor").
		First(&itemIntenvory, req.ItemInventoryID).Error

	return itemIntenvory, err
}
