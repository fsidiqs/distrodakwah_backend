package repository

import (
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/helper/httphelper"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/helper/pagination"
	inventoryModel "github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/inventory/model"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/inventory/request"
	productModel "github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/product/model"
)

func (r *InventoryRepository) FetchAll(req request.FetchAllReq) (*pagination.Pagination, error) {
	var err error
	var total int64

	res := &pagination.Pagination{Metadata: pagination.Metadata{}}

	products := []*productModel.ProductSimpleInfo{}
	productQ := r.DB.Model(&productModel.ProductSimpleInfo{}).Unscoped()

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
	items := []productModel.Item{}

	err = r.DB.Model(&productModel.Item{}).
		Where("product_id IN (?)", prodIDArr).
		Preload("Product").Preload("ItemInventory.ItemInventoryDetail.UserVendor").
		Find(&items).Error

	// STEP Of variantProducts
	res.UpdateElements(items)

	return res, nil
}

func (r *InventoryRepository) ExportInventory() ([]productModel.Item, error) {
	var err error

	products := []*productModel.ProductSimpleInfo{}
	err = r.DB.Model(&productModel.ProductSimpleInfo{}).Unscoped().Find(&products).Error

	prodIDArr := []uint64{}
	for _, product := range products {
		prodIDArr = append(prodIDArr, product.ID)
	}

	if err != nil {
		return nil, err
	}

	items := []productModel.Item{}
	err = r.DB.Model(&productModel.Item{}).
		Preload("Product").Preload("ItemInventory.ItemInventoryDetail.UserVendor").
		Find(&items).Error
	return items, nil
}

type FindReq struct {
	ItemInventoryID uint64
	Preload         httphelper.Preload
}

func (ir *InventoryRepository) Find(req FindReq) (*inventoryModel.ItemInventory, error) {
	var err error

	itemIntenvory := &inventoryModel.ItemInventory{}

	err = ir.DB.Model(&inventoryModel.ItemInventory{}).
		Preload("ItemInventoryDetail.UserVendor").
		First(&itemIntenvory, req.ItemInventoryID).Error

	return itemIntenvory, err
}
