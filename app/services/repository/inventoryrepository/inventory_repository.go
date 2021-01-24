package inventoryrepository

import (
	"distrodakwah_backend/app/helper/httphelper"
	"distrodakwah_backend/app/helper/pagination"
	"distrodakwah_backend/app/services/handler/inventoryhandler"
	"distrodakwah_backend/app/services/model/inventorymodel"
	"distrodakwah_backend/app/services/model/productmodel"
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

	prodIDArr := []int{}
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

	prodIDArr := []int{}
	for _, product := range products {
		prodIDArr = append(prodIDArr, product.ID)
	}

	if err != nil {
		return nil, err
	}

	items := []productmodel.Item{}
	err = r.DB.Model(&productmodel.Item{}).
		Preload("Product").Preload("ItemInventory.ItemInventoryDetail.Subdistrict").
		Find(&items).Error
	return items, nil
}

type FindReq struct {
	ItemInventoryID int
	Preload         httphelper.Preload
}

func (ir *InventoryRepository) Find(req FindReq) (*inventorymodel.SPIInventory, error) {
	return nil, nil
	// var err error

	// itemIntenvory := &inventorymodel.ItemInventory{}

	// err = ir.DB.Model(&inventorymodel.ItemInventory{}).
	// 	Preload("ItemInventoryDetail.Subdistrict").
	// 	First(&itemIntenvory, req.ItemInventoryID).Error

	// return itemIntenvory, err
}
