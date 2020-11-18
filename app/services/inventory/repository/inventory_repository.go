package repository

import (
	"sort"

	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/helper/httphelper"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/helper/pagination"
	invModel "github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/inventory/model"
	invModelAux "github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/inventory/model/aux"
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

	if len(req.ProductIDArr) > 0 {
		productQ = productQ.Where("products.id IN (?)", req.ProductIDArr)
	}
	// count
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

	inventories := []*prodModel.ProductInventoryForFetch{}

	for _, sp := range singleProducts {
		inventories = append(inventories, &prodModel.ProductInventoryForFetch{
			ID:            sp.ProductID,
			ProductKindID: sp.Product.ProductKindID,
			Sku:           sp.Product.Sku,
			SingleProduct: sp,
		})
	}

	for _, vp := range variantProducts {
		inventories = append(inventories, &prodModel.ProductInventoryForFetch{
			ID:             vp.ProductID,
			ProductKindID:  vp.Product.ProductKindID,
			Sku:            vp.Product.Sku,
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
				ProductKindID: prod.ProductKindID,
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
				ProductKindID:   prod.ProductKindID,
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

type FindReq struct {
	RelatedID     uint64
	ProductKindID uint8
	Preload       httphelper.Preload
}

func (ir *InventoryRepository) Find(req FindReq) (*invModelAux.InventoryResponse, error) {
	var err error
	var inv *invModelAux.InventoryResponse

	if req.ProductKindID == prodModel.ProductKindSingle {
		inv = &invModelAux.InventoryResponse{
			SPInventory: &invModel.SPInventory{},
		}
		query := ir.DB.Model(&invModel.SPInventory{}).
			Where("sp_inventory.single_product_id = ?", req.RelatedID)

		if req.Preload != nil { // check whether slice is empty
			HandlePreload(query, req.Preload, prodModel.ProductKindSingle)
		}
		err = query.First(&inv.SPInventory).Error

	} else if req.ProductKindID == prodModel.ProductKindVariant {
		inv = &invModelAux.InventoryResponse{
			VPInventory: &invModel.VPInventory{},
		}

		query := ir.DB.Debug().Model(&invModel.VPInventory{}).
			Where("vp_inventory.variant_product_id = ?", req.RelatedID)

		if req.Preload != nil { // check whether slice is empty
			HandlePreload(query, req.Preload, prodModel.ProductKindVariant)
		}
		err = query.First(&inv.VPInventory).Error
	}

	return inv, err
}
