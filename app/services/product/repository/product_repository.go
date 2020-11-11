package repository

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/helper/pagination"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/product/model"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/product/request"
	"gorm.io/gorm"
)

func (r *ProductRepository) FetchByColumns(req *request.FetchByColumnReq) (*Pagination, error) {

	res := &Pagination{Metadata: &pagination.Metadata{}}
	products := []*model.ProductResponse{}
	query := r.DB.Model(&model.ProductResponse{}).
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
	products := []*model.ProductResponse{}
	query := r.DB.Model(&model.ProductResponse{}).
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

func (r *ProductRepository) SaveProduct(product *model.Product) (*model.Product, error) {
	// tx := r.DB.Begin()
	// err := tx.Model(&model.ProductImage{}).Create(product).Error

	// if err != nil {
	// 	tx.Rollback()
	// 	return nil, err
	// }

	// return product, tx.Commit().Error
	return nil, nil
}

func (r *ProductRepository) SaveProductTx(product *model.Product, tx *gorm.DB) error {

	// err := tx.Model(&model.Product{}).Create(&product).Error
	// fmt.Printf("product afters: %+v \n", product.ID)

	// if err != nil {
	// 	fmt.Printf("error creating product\n %+v \n", err)
	// 	tx.Rollback()
	// 	return nil
	// }

	return nil
}

func (r *ProductRepository) SaveProductBasicStructure(productReqJSON *model.ProductFromRequestJSON) error {

	var err error
	tx := r.DB.Begin()
	productImagesReq := &model.ProductImages{}
	// convert product image request array into db like struct

	//  STEP create product image
	err = json.NewDecoder(strings.NewReader(productReqJSON.ProductImages)).Decode(&productImagesReq)
	err = productImagesReq.Validate()
	if err != nil {
		fmt.Println("product images are invalid")
		return err
	}
	err = tx.Model(&model.ProductImage{}).Create(&productImagesReq).Error
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
	productRes := &model.Product{
		BrandID:       productReqJSON.BrandID,
		CategoryID:    productReqJSON.CategoryID,
		ProductTypeID: productReqJSON.ProductTypeID,
		ProductKindID: productReqJSON.ProductKindID,
		Status:        productReqJSON.Status,
		Name:          productReqJSON.Name,
		Description:   productReqJSON.Description,
		Sku:           productReqJSON.Sku,
	}

	err = tx.Model(&model.Product{}).Create(&productRes).Error

	if err != nil {
		fmt.Printf("error creating product \n %+v \n", err)
		tx.Rollback()
		return nil
	}

	//  Create ProductsProductImages

	var productsProductImages []*model.ProductsProductImage
	for _, pi := range *productImagesReq {
		productsProductImages = append(
			productsProductImages,
			&model.ProductsProductImage{
				ProductID:      productRes.ID,
				ProductImageID: pi.ID,
			},
		)
	}

	err = tx.Model(&model.ProductsProductImage{}).Create(&productsProductImages).Error
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

		SingleProductDetailReq := &request.SingleProductDetailReq{}
		err = json.NewDecoder(strings.NewReader(productReqJSON.SingleProductDetail)).Decode(&SingleProductDetailReq)
		// STEP ofCreating single product
		singleProduct := &model.SingleProduct{
			ProductID: productRes.ID,
			Weight:    SingleProductDetailReq.Weight,
		}

		err = tx.Model(&model.SingleProduct{}).Create(&singleProduct).Error

		if err != nil {
			fmt.Printf("Error Creating Single Product \n %+v \n", err)
			tx.Rollback()
			return err
		}

		// STEP Of creating singleProductPriceArr

		singleProductPriceArr := []*model.SingleProductsPrices{
			{
				SingleProductID: singleProduct.ID,
				Name:            request.HargaJualName,
				Value:           SingleProductDetailReq.Price,
			},
		}
		err = tx.Model(&model.SingleProductsPrices{}).Create(&singleProductPriceArr).Error
		if err != nil {
			fmt.Printf("error creating Single Product Prices\n %+v \n", err)
			tx.Rollback()
			return err
		}

	} else if productReqJSON.ProductKindID == 2 {
		variantProductDetailReqs := []*request.VariantProductDetailReq{}
		err = json.NewDecoder(strings.NewReader(productReqJSON.VariantProductDetail)).Decode(&variantProductDetailReqs)
		//STEP VariantProduct Create

		for _, variantProductDetailReq := range variantProductDetailReqs {
			// Creating VariantProduct
			var variantProduct *model.VariantProduct
			variantProduct = &model.VariantProduct{
				ProductID: productRes.ID,
				Sku:       variantProductDetailReq.Sku,
				Weight:    variantProductDetailReq.Weight,
			}

			err = tx.Model(&model.VariantProduct{}).Create(&variantProduct).Error
			if err != nil {
				fmt.Printf("Error Creating VariantProduct \n %+v \n", err)
				tx.Rollback()
				return err
			}

			//STEP variant_products_prices
			variantProductPrice := &model.VariantProductsPrices{
				VariantProductID: variantProduct.ID,
				Name:             request.HargaJualName,
				Value:            variantProductDetailReq.SellingPrice,
			}

			err = tx.Model(&model.VariantProductsPrices{}).Create(&variantProductPrice).Error
			if err != nil {
				fmt.Printf("Error Creating VariantProductsPrices \n %+v \n", err)
				tx.Rollback()
				return err
			}

			//STEP Variants Create
			for _, variantReq := range variantProductDetailReq.Variants {
				variant := &model.Variant{
					ProductID: productRes.ID,
					Name:      variantReq.VariantValue,
				}
				err = tx.Model(&model.Variant{}).Create(&variant).Error
				if err != nil {
					fmt.Printf("Error Creating Variant \n %+v \n", err)
					tx.Rollback()
					return err
				}

				option := &model.Option{
					VariantID:        variant.ID,
					Name:             variantReq.OptionValue,
					VariantProductID: variantProduct.ID,
				}
				err = tx.Model(&model.Option{}).Create(&option).Error
				if err != nil {
					fmt.Printf("Error Creating Option \n %+v \n", err)
					tx.Rollback()
					return err
				}

			}

		}

	}

	// 5. Create Single Product Prices

	return tx.Commit().Error
}
