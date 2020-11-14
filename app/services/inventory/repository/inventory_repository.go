package repository

import (
	"sort"

	modelHelper "github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/helper/model"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/helper/pagination"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/inventory/request"
	prodModel "github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/product/model"
)

func (r *InventoryRepository) FetchAll(req *request.FetchAllReq) (*pagination.Pagination, error) {
	var err error
	var total int64

	res := &pagination.Pagination{Metadata: &pagination.Metadata{}}

	// inventories := []*model.InventoryResponse{}
	products := []*prodModel.ProductSimpleInfo{}
	productQ := r.DB.Model(&prodModel.ProductSimpleInfo{}).Unscoped()
	err = productQ.Count(&total).Error
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
	singleProducts := []*prodModel.SingleProductStock{}

	err = r.DB.Model(&prodModel.SingleProductStock{}).
		Where("single_products.product_id IN (?)", prodIDArr).
		Preload("Product").Preload("SPInventory.SPInventoryDetail.UserVendor").
		Find(&singleProducts).Error

	// STEP Of variantProducts
	variantProducts := []*prodModel.VariantProductStock{}
	err = r.DB.Model(&prodModel.VariantProductStock{}).
		Where("variant_products.product_id IN (?)", prodIDArr).
		Preload("Product").Preload("VPInventory.VPInventoryDetail.UserVendor").
		Find(&variantProducts).Error
	if err != nil {
		return nil, err
	}

	inventories := []*modelHelper.InventoryResponse{}

	for _, sp := range singleProducts {
		inventories = append(inventories, &modelHelper.InventoryResponse{
			SingleProduct: sp,
		})
	}

	for _, vp := range variantProducts {
		inventories = append(inventories, &modelHelper.InventoryResponse{
			VariantProduct: vp,
		})
	}
	res.UpdateElements(inventories)

	return res, nil
}

func (r *InventoryRepository) ExportInventory() ([]*prodModel.ProductInventory, error) {
	var err error

	products := []*prodModel.ProductSimpleInfo{}
	err = r.DB.Model(&prodModel.ProductSimpleInfo{}).Unscoped().Find(&products).Error

	prodIDArr := []uint64{}
	for _, product := range products {
		prodIDArr = append(prodIDArr, product.ID)
	}

	prodSPs := []*prodModel.ProductInventory{}
	err = r.DB.Model(&prodModel.ProductInventory{}).
		Where("id IN (?) AND product_kind_id = ?", prodIDArr, prodModel.ProductKindSingle).
		Preload("SingleProduct.SPInventory").
		Find(&prodSPs).Error

	prodVPs := []*prodModel.ProductInventory{}
	err = r.DB.Model(&prodModel.ProductInventory{}).
		Preload("VariantProducts.VPInventory").
		Where("id IN (?) AND product_kind_id = ?", prodIDArr, prodModel.ProductKindVariant).
		Find(&prodVPs).Error
	if err != nil {
		return nil, err
	}

	returnedData := []*prodModel.ProductInventory{}
	for _, prod := range prodSPs {
		returnedData = append(
			returnedData,
			&prodModel.ProductInventory{
				ID:            prod.ID,
				Sku:           prod.Sku,
				SingleProduct: prod.SingleProduct,
			},
		)
	}

	for _, prod := range prodVPs {
		returnedData = append(
			returnedData,
			&prodModel.ProductInventory{
				ID:              prod.ID,
				Sku:             prod.Sku,
				VariantProducts: prod.VariantProducts,
			},
		)
	}

	// STEP SORT
	sort.Slice(returnedData, func(i, j int) bool {
		return returnedData[i].ID < returnedData[j].ID
	})
	return returnedData, nil
}
