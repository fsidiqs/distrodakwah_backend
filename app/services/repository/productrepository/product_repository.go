package productrepository

import (
	"encoding/json"
	"fmt"
	"strings"

	"distrodakwah_backend/app/helper/pagination"
	"distrodakwah_backend/app/services/handler/producthandler"
	"distrodakwah_backend/app/services/model/inventorymodel"
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

func (r *ProductRepository) SaveProductBasicStructure(productReqJSON *producthandler.ProductFromRequestJSON) error {

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
	err = tx.Model(&productmodel.ProductImage{}).Create(&productImagesReq).Error
	if err != nil {
		fmt.Printf("error creating product_images\n %+v \n", err)
		tx.Rollback()
		return err
	}

	// STEP Create Product and prepare returned result
	productRes := &productmodel.Product{
		BrandID:       productReqJSON.BrandID,
		CategoryID:    productReqJSON.CategoryID,
		ProductTypeID: productReqJSON.ProductTypeID,
		ProductKindID: productReqJSON.ProductKindID,
		Status:        productReqJSON.Status,
		Name:          productReqJSON.Name,
		Description:   productReqJSON.Description,
	}

	err = tx.Model(&productmodel.Product{}).Create(&productRes).Error

	if err != nil {
		fmt.Printf("error creating product \n %+v \n", err)
		tx.Rollback()
		return nil
	}

	//  Create ProductsProductImages

	var productsProductImages []productmodel.ProductsProductImage
	for _, pi := range productImagesReq {
		productsProductImages = append(
			productsProductImages,
			productmodel.ProductsProductImage{
				ProductID:      productRes.ID,
				ProductImageID: pi.ID,
			},
		)
	}

	err = tx.Model(&productmodel.ProductsProductImage{}).Create(&productsProductImages).Error
	if err != nil {
		fmt.Printf("error creating ProductsProductImages\n %+v \n", err)
		tx.Rollback()
		return err
	}
	// STEP of creating Items
	itemReqs := []producthandler.ItemCreateBasicProduct{}
	err = json.NewDecoder(strings.NewReader(productReqJSON.Items)).Decode(&itemReqs)

	items := []productmodel.Item{}
	// variant or single
	if productReqJSON.ProductKindID == productmodel.ProductKindVariant {
		variantCreateReqs := []*productmodel.Variant{}
		err = json.NewDecoder(strings.NewReader(productReqJSON.Variants)).Decode(&variantCreateReqs)
		for _, v := range variantCreateReqs {
			v.ProductID = productRes.ID
		}
		err = tx.Model(&productmodel.Variant{}).Create(&variantCreateReqs).Error

		if err != nil {
			fmt.Println("product creating variants")
			return err
		}

		for idx, itemReq := range itemReqs {
			// STEP of creating options
			optionCreateReqs := []productmodel.Option{}
			err = json.NewDecoder(strings.NewReader(itemReq.Options)).Decode(&optionCreateReqs)

			// populate option itemID
			for i := 0; i < len(optionCreateReqs); i++ {
				optionCreateReqs[i].VariantID = variantCreateReqs[idx].ID
			}

			// prepare item inventory
			itemInventoryReqslice := []producthandler.ItemInventoryRequestCreateBasicProduct{}
			err = json.NewDecoder(strings.NewReader(itemReq.ItemInventories)).Decode(&itemInventoryReqslice)

			itemInventory := []inventorymodel.ItemInventory{}
			for _, itemInventoryReq := range itemInventoryReqslice {
				itemInventory = append(
					itemInventory,
					inventorymodel.ItemInventory{
						ItemInventoryDetail: &inventorymodel.ItemInventoryDetail{
							SubdistrictID: itemInventoryReq.SubdistrictID,
						},
					},
				)
			}

			items = append(items,
				productmodel.Item{
					ProductID: productRes.ID,
					Weight:    itemReq.Weight,
					Sku:       itemReq.Sku,
					Options:   optionCreateReqs,
					Prices: []productmodel.ItemPrice{
						{
							Name:  producthandler.RetailPriceName,
							Value: itemReq.Price,
						},
					},
					ItemInventory: itemInventory,
				},
			)

		}

	} else if productReqJSON.ProductKindID == productmodel.ProductKindSingle {
		for _, itemReq := range itemReqs {
			// prepare item inventory
			itemInventoryReqslice := []producthandler.ItemInventoryRequestCreateBasicProduct{}
			err = json.NewDecoder(strings.NewReader(itemReq.ItemInventories)).Decode(&itemInventoryReqslice)

			// loop itemreq
			itemInventory := []inventorymodel.ItemInventory{}
			for _, itemInventoryReq := range itemInventoryReqslice {
				itemInventory = append(
					itemInventory,
					inventorymodel.ItemInventory{
						ItemInventoryDetail: &inventorymodel.ItemInventoryDetail{
							SubdistrictID: itemInventoryReq.SubdistrictID,
						},
					},
				)
			}

			items = append(
				items,
				productmodel.Item{
					ProductID: productRes.ID,
					Sku:       itemReq.Sku,
					Weight:    itemReq.Weight,
					Prices: []productmodel.ItemPrice{
						{
							Name:  producthandler.RetailPriceName,
							Value: itemReq.Price,
						},
					},
					ItemInventory: itemInventory,
				},
			)
		}

	}

	err = tx.Model(&productmodel.Item{}).Create(&items).Error

	if err != nil {
		fmt.Printf("Error Creating Product \n %+v \n", err)
		tx.Rollback()
		return err
	}

	// STEP creating iteminventory

	return tx.Commit().Error
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
