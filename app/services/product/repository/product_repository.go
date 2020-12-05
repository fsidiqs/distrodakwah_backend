package repository

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/database"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/helper/pagination"
	invModel "github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/inventory/model"
	prodModel "github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/product/model"
	"gorm.io/gorm"

	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/product/request"
)

func (r *ProductRepository) FetchByColumns(req *request.FetchByColumnReq) (*Pagination, error) {

	res := &Pagination{Metadata: pagination.Metadata{}}
	products := []*prodModel.ProductResponse{}
	query := r.DB.Model(&prodModel.ProductResponse{}).
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

	res.paginate(req.Metadata)
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

func (r *ProductRepository) FetchAll(req *request.FetchAllReq) (*Pagination, error) {

	res := &Pagination{Metadata: pagination.Metadata{}}
	products := []*prodModel.ProductResponse{}
	query := r.DB.Model(&prodModel.ProductResponse{}).
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

	res.paginate(req.Metadata)
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

func (p *Pagination) paginate(m pagination.Metadata) {

	page, limit, offset := pagination.BuildPagination(m)
	p.Metadata = pagination.Metadata{
		Total:  p.Metadata.Total,
		Limit:  limit,
		Offset: offset,
		Page:   page,
		Pages:  pagination.BuildPages(p.Metadata.Total, limit),
	}
}

func (r *ProductRepository) SaveProductBasicStructure(productReqJSON *request.ProductFromRequestJSON) error {

	var err error
	tx := r.DB.Begin()
	productImagesReq := productReqJSON.ProductImages
	// convert product image request array into db like struct

	//  STEP create product image
	err = productImagesReq.Validate()
	if err != nil {
		fmt.Println("product images are invalid")
		return err
	}
	err = tx.Model(&prodModel.ProductImage{}).Create(&productImagesReq).Error
	if err != nil {
		fmt.Printf("error creating product_images\n %+v \n", err)
		tx.Rollback()
		return err
	}

	// STEP Create Product and prepare returned result
	productRes := &prodModel.Product{
		BrandID:       productReqJSON.BrandID,
		CategoryID:    productReqJSON.CategoryID,
		ProductTypeID: productReqJSON.ProductTypeID,
		ProductKindID: productReqJSON.ProductKindID,
		Status:        productReqJSON.Status,
		Name:          productReqJSON.Name,
		Description:   productReqJSON.Description,
	}

	err = tx.Model(&prodModel.Product{}).Create(&productRes).Error

	if err != nil {
		fmt.Printf("error creating product \n %+v \n", err)
		tx.Rollback()
		return nil
	}

	//  Create ProductsProductImages

	var productsProductImages []prodModel.ProductsProductImage
	for _, pi := range productImagesReq {
		productsProductImages = append(
			productsProductImages,
			prodModel.ProductsProductImage{
				ProductID:      productRes.ID,
				ProductImageID: pi.ID,
			},
		)
	}

	err = tx.Model(&prodModel.ProductsProductImage{}).Create(&productsProductImages).Error
	if err != nil {
		fmt.Printf("error creating ProductsProductImages\n %+v \n", err)
		tx.Rollback()
		return err
	}
	// STEP of creating Items
	itemReqs := []request.ItemCreateBasicProduct{}
	err = json.NewDecoder(strings.NewReader(productReqJSON.Items)).Decode(&itemReqs)
	items := []prodModel.Item{}

	if productReqJSON.ProductKindID == prodModel.ProductKindVariant {
		variantCreateReqs := []*prodModel.Variant{}
		err = json.NewDecoder(strings.NewReader(productReqJSON.Variants)).Decode(&variantCreateReqs)
		for _, v := range variantCreateReqs {
			v.ProductID = productRes.ID
		}
		err = tx.Model(&prodModel.Variant{}).Create(&variantCreateReqs).Error
		if err != nil {
			fmt.Println("product creating variants")
			return err
		}

		for idx, itemReq := range itemReqs {
			// STEP of creating options
			optionCreateReqs := []prodModel.Option{}
			err = json.NewDecoder(strings.NewReader(itemReq.Options)).Decode(&optionCreateReqs)

			// populate option itemID
			for i := 0; i < len(optionCreateReqs); i++ {
				optionCreateReqs[i].VariantID = variantCreateReqs[idx].ID
			}
			items = append(items,
				prodModel.Item{
					ProductID: productRes.ID,
					Weight:    itemReq.Weight,
					Sku:       itemReq.Sku,
					Options:   optionCreateReqs,
					Prices: []prodModel.ItemPrice{
						{
							Name:  request.RetailPriceName,
							Value: itemReq.Price,
						},
					},
				},
			)

		}

	} else if productReqJSON.ProductKindID == prodModel.ProductKindSingle {
		for _, itemReq := range itemReqs {
			items = append(
				items,
				prodModel.Item{
					ProductID: productRes.ID,
					Sku:       itemReq.Sku,
					Weight:    itemReq.Weight,
					Prices: []prodModel.ItemPrice{
						{
							Name:  request.RetailPriceName,
							Value: itemReq.Price,
						},
					},
				},
			)
		}

	}

	err = tx.Model(&prodModel.Item{}).Create(&items).Error

	if err != nil {
		fmt.Printf("Error Creating Single Product \n %+v \n", err)
		tx.Rollback()
		return err
	}

	brandDB := &prodModel.Brand{}

	if productReqJSON.ProductTypeID == prodModel.ProductTypeVendor {
		err = database.DB.Model(&prodModel.Brand{}).Where("id = ?", productReqJSON.BrandID).Find(&brandDB).Error
	} else {
		brandDB.UserVendorID = prodModel.ProductTypeConsignment
	}
	// STEP ofCreating single product

	itemInventories := []invModel.ItemInventory{}
	for _, item := range items {
		itemInventories = append(
			itemInventories,
			invModel.ItemInventory{
				Keep:   0,
				Stock:  0,
				ItemID: item.ID,
				ItemInventoryDetail: &invModel.ItemInventoryDetail{
					VendorID: brandDB.UserVendorID,
				},
			},
		)
	}

	err = tx.Debug().Model(&invModel.ItemInventory{}).Create(&itemInventories).Error
	if err != nil {
		fmt.Printf("Error Creating Single Product \n %+v \n", err)
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r ProductRepository) TxUpdateProduct(tx *gorm.DB, productReq prodModel.Product) (*gorm.DB, error) {
	var err error

	err = tx.Model(&prodModel.Product{}).Where("id = ?", productReq.ID).Updates(productReq).Error
	if err != nil {
		fmt.Println("could not update product")
		return nil, err
	}

	return tx, nil
}
