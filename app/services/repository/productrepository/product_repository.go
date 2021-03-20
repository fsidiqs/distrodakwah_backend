package productrepository

import (
	"fmt"

	"distrodakwah_backend/app/helper/pagination"
	"distrodakwah_backend/app/services/handler/producthandler"
	"distrodakwah_backend/app/services/library/productlibrary"
	"distrodakwah_backend/app/services/library/query/productquery"
	"distrodakwah_backend/app/services/model/productmodel"

	"gorm.io/gorm"
)

func (r *ProductRepository) FetchByColumns(req *producthandler.FetchByColumnReq) (*Pagination, error) {

	res := &Pagination{Metadata: pagination.Metadata{}}
	products := []*productmodel.ProductResponse{}
	query := r.DB.Model(&productmodel.ProductResponse{}).
		Unscoped()
	if len(req.PTypeIDs) > 0 {
		query = query.Where("product_type_id in (?)", req.PTypeIDs)
	}
	if len(req.PKindIDs) > 0 {
		query = query.Where("product_kind_id in (?)", req.PKindIDs)
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

	// res.paginate(req.Metadata)
	if req.Preload != nil { // check whether slice is empty
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

func (r *ProductRepository) FetchAll(req *producthandler.FetchAllReq) (*Pagination, error) {

	res := &Pagination{Metadata: pagination.Metadata{}}
	products := []*productmodel.ProductResponse{}
	query := r.DB.Model(&productmodel.ProductResponse{}).
		Unscoped()

	// if it has product_id_arr
	if len(req.ProductIDArr) > 0 {
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

	// res.paginate(req.Metadata)
	if req.Preload != nil { // check whether slice is empty
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

func (r *ProductRepository) FetchByItemInventoryID(InventoryItemID [][]int) ([]interface{}, error) {
	var err error
	SPIInvIDs := InventoryItemID[0]
	VPIInvIDs := InventoryItemID[1]

	SPIInventories := []productlibrary.SPIInventory{}
	VPIInventories := []productlibrary.VPIInventory{}

	err = r.DB.Raw(productquery.SEL_SP_BY_ITEM_INVENTORY_ID, SPIInvIDs).
		Preload("SingleProductItem.SingleProduct.ProductImages").
		Find(&SPIInventories).
		Error

	if err != nil {
		fmt.Println("error fetching sp products")
		return nil, err
	}

	err = r.DB.Raw(productquery.SEL_VP_BY_ITEM_INVENTORY_ID, VPIInvIDs).
		Preload("VPItem.VariantProduct.ProductImages").
		Find(&VPIInventories).
		Error

	if err != nil {
		fmt.Println("error fetching sp products")
		return nil, err
	}
	var productResponseable []interface{}

	for _, sp := range SPIInventories {

		productResponseable = append(productResponseable, sp)
	}
	for _, vp := range VPIInventories {

		productResponseable = append(productResponseable, vp)
	}

	return productResponseable, nil
}

func (r *ProductRepository) FetchByItemID(ItemIDArr [][]int) ([]interface{}, error) {
	fmt.Printf("test")

	var err error
	SPItemIDs := ItemIDArr[0]
	VPItemIDs := ItemIDArr[1]

	SPIInventories := []productlibrary.SingleProductItem{}
	VPIInventories := []productlibrary.VariantProductItem{}

	err = r.DB.Raw(productquery.SEL_SP_BY_ITEM_ID, SPItemIDs).
		Preload("SingleProduct.ProductImages").
		Preload("SPIPrices").
		Find(&SPIInventories).
		Error

	if err != nil {
		fmt.Println("error fetching sp products")
		return nil, err
	}

	err = r.DB.Raw(productquery.SEL_VP_BY_ITEM_ID, VPItemIDs).
		Preload("VariantProduct.ProductImages").
		Preload("VariantProductOptions.VariantProductVariant").
		Preload("VPItemPrices").
		Find(&VPIInventories).
		Error

	if err != nil {
		fmt.Println("error fetching sp products")
		return nil, err
	}

	var productResponseable []interface{}
	for _, sp := range SPIInventories {

		productResponseable = append(productResponseable, sp)
	}
	for _, vp := range VPIInventories {

		productResponseable = append(productResponseable, vp)
	}
	return productResponseable, nil
}

func (r ProductRepository) TxUpdateProduct(tx *gorm.DB, productReq productmodel.Product) (*gorm.DB, error) {
	var err error

	err = tx.Model(&productmodel.Product{}).Where("id = ?", productReq.ID).Updates(productReq).Error
	if err != nil {
		fmt.Println("could not update product")
		return nil, err
	}

	return tx, nil
}
