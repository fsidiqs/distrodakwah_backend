package repository

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/database"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/helper/pagination"
	invModel "github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/inventory/model"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/product/model"
	prodModel "github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/product/model"

	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/product/request"
	"gorm.io/gorm"
)

func (r *ProductRepository) FetchByColumns(req *request.FetchByColumnReq) (*Pagination, error) {

	res := &Pagination{Metadata: &pagination.Metadata{}}
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

	res := &Pagination{Metadata: &pagination.Metadata{}}
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

func (r *ProductRepository) SaveProduct(product *prodModel.Product) (*prodModel.Product, error) {
	// tx := r.DB.Begin()
	// err := tx.Model(&prodModel.ProductImage{}).Create(product).Error

	// if err != nil {
	// 	tx.Rollback()
	// 	return nil, err
	// }

	// return product, tx.Commit().Error
	return nil, nil
}

func (r *ProductRepository) SaveProductTx(product *prodModel.Product, tx *gorm.DB) error {

	// err := tx.Model(&prodModel.Product{}).Create(&product).Error
	// fmt.Printf("product afters: %+v \n", product.ID)

	// if err != nil {
	// 	fmt.Printf("error creating product\n %+v \n", err)
	// 	tx.Rollback()
	// 	return nil
	// }

	return nil
}

func (r *ProductRepository) SaveProductBasicStructure(productReqJSON *prodModel.ProductFromRequestJSON) error {

	var err error
	tx := r.DB.Begin()
	productImagesReq := &productReqJSON.ProductImages
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

	//  STEP create product prices

	// if productReqJSON.ProductKindID == 1 {

	// 	err = json.NewDecoder(strings.NewReader(productReqJSON.SingleProductDetail)).Decode(&productImagesReq)
	// 	err = productReqJSON.ProductDetail.SingleProductHargaJual.Validate()
	// } else if productReqJSON.ProductKindID == 2 {

	// }
	if err != nil {
		fmt.Printf("%+v \n", err)
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
		Sku:           productReqJSON.Sku,
	}

	err = tx.Model(&prodModel.Product{}).Create(&productRes).Error

	if err != nil {
		fmt.Printf("error creating product \n %+v \n", err)
		tx.Rollback()
		return nil
	}

	//  Create ProductsProductImages

	var productsProductImages []*prodModel.ProductsProductImage
	for _, pi := range *productImagesReq {
		productsProductImages = append(
			productsProductImages,
			&prodModel.ProductsProductImage{
				ProductID:      productRes.ID,
				ProductImageID: pi.ID,
			},
		)
	}

	err = tx.Model(&prodModel.ProductsProductImage{}).Create(&productsProductImages).Error
	if err != nil {
		fmt.Printf("error creating ProductsProductImages\n %+v \n", err)
		fmt.Println("test")

		tx.Rollback()
		return err
	}
	if err != nil {
		fmt.Println("product Repository error Creating ProductsProductImage")
		return err
	}

	// 4.a create singleProduct
	if productReqJSON.ProductKindID == 1 {

		// set uservendor id, if type is a vendor then find vendor, else use vendorid 1 (distrodakwah)
		brandDB := &model.Brand{}
		if productReqJSON.ProductTypeID == model.ProductTypeVendor {
			err = database.DB.Model(&model.Brand{}).Where("id = ?", productReqJSON.BrandID).Find(&brandDB).Error
		} else {
			brandDB.UserVendorID = model.ProductTypeConsignment
		}

		if err != nil {
			tx.Rollback()
			return err
		}

		SingleProductDetailReq := &request.SingleProductDetailReq{}
		err = json.NewDecoder(strings.NewReader(productReqJSON.SingleProductDetail)).Decode(&SingleProductDetailReq)
		// STEP ofCreating single product
		initSPInventory := &invModel.SPInventory{
			Keep:  0,
			Stock: 0,
			SPInventoryDetail: &invModel.SPInventoryDetail{
				VendorID: brandDB.UserVendorID,
			},
		}

		singleProduct := &prodModel.SingleProductStock{
			ProductID:   productRes.ID,
			Weight:      SingleProductDetailReq.Weight,
			SPInventory: initSPInventory,
		}

		err = tx.Debug().Model(&prodModel.SingleProductStock{}).Create(&singleProduct).Error

		if err != nil {
			fmt.Printf("Error Creating Single Product \n %+v \n", err)
			tx.Rollback()
			return err
		}

		// STEP Of creating singleProductPriceArr

		singleProductPriceArr := prodModel.SingleProductsPriceArr{
			{
				SingleProductID: singleProduct.ID,
				Name:            request.HargaJualName,
				Value:           SingleProductDetailReq.Price,
			},
		}
		err = tx.Model(&prodModel.SingleProductsPrice{}).Create(&singleProductPriceArr).Error
		if err != nil {
			fmt.Printf("error creating Single Product Prices\n %+v \n", err)
			tx.Rollback()
			return err
		}

	} else if productReqJSON.ProductKindID == 2 {
		variantProductDetailReqs := []*request.VariantProductDetailReq{}
		err = json.NewDecoder(strings.NewReader(productReqJSON.VariantProductDetail)).Decode(&variantProductDetailReqs)
		//STEP VariantProduct Create

		// set uservendor id, if type is a vendor then find vendor, else use vendorid 1 (distrodakwah)
		brandDB := &model.Brand{}
		if productReqJSON.ProductTypeID == model.ProductTypeVendor {
			err = database.DB.Model(&model.Brand{}).Where("id = ?", productReqJSON.BrandID).Find(&brandDB).Error
		} else {
			brandDB.UserVendorID = model.ProductTypeConsignment
		}

		for _, variantProductDetailReq := range variantProductDetailReqs {
			// Creating VariantProduct
			initVPInventory := &invModel.VPInventory{
				Keep:  0,
				Stock: 0,
				VPInventoryDetail: &invModel.VPInventoryDetail{
					VendorID: brandDB.UserVendorID,
				},
			}
			variantProduct := &prodModel.VariantProductStock{
				ProductID:   productRes.ID,
				Sku:         variantProductDetailReq.Sku,
				Weight:      variantProductDetailReq.Weight,
				VPInventory: initVPInventory,
			}

			// fmt.Printf("Error Before Creating VariantProduct \n %+v \n", variantProduct.VPInventory.VPInventoryDetail.VendorID)

			err = tx.Model(&prodModel.VariantProductStock{}).Create(&variantProduct).Error
			if err != nil {
				fmt.Printf("Error Creating VariantProduct \n %+v \n", err)
				tx.Rollback()
				return err
			}

			//STEP variant_products_prices
			variantProductPrice := &prodModel.VariantProductsPrice{
				VariantProductID: variantProduct.ID,
				Name:             request.HargaJualName,
				Value:            variantProductDetailReq.SellingPrice,
			}

			err = tx.Model(&prodModel.VariantProductsPrice{}).Create(&variantProductPrice).Error
			if err != nil {
				fmt.Printf("Error Creating VariantProductsPrices \n %+v \n", err)
				tx.Rollback()
				return err
			}

			//STEP Variants Create
			for _, variantReq := range variantProductDetailReq.Variants {
				variant := &prodModel.Variant{
					ProductID: productRes.ID,
					Name:      variantReq.VariantValue,
				}
				err = tx.Model(&prodModel.Variant{}).Create(&variant).Error
				if err != nil {
					fmt.Printf("Error Creating Variant \n %+v \n", err)
					tx.Rollback()
					return err
				}

				option := &prodModel.Option{
					VariantID:        variant.ID,
					Name:             variantReq.OptionValue,
					VariantProductID: variantProduct.ID,
				}
				err = tx.Model(&prodModel.Option{}).Create(&option).Error
				if err != nil {
					fmt.Printf("Error Creating Option \n %+v \n", err)
					tx.Rollback()
					return err
				}

			}

		}

	}

	return tx.Commit().Error
}
