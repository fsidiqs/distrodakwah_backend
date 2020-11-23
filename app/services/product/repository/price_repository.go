package repository

import (
	"fmt"

	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/product/model"
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/services/product/model/aux"
)

type M map[string]interface{}

func (r *ProductRepository) GeneratePriceTemplate(productIDArr []int) (*aux.ProductPriceTemplate, error) {

	// ! reconstruct code with getting each table product
	singleProductPrices := model.SingleProductsPriceArrWithParent{}
	err := r.DB.Model(&model.SingleProductsPriceWithParent{}).
		Preload("SingleProduct.Product").
		Find(&singleProductPrices).Error

	variantProductPrices := model.VariantProductsPriceArrWithParent{}
	err = r.DB.Model(&model.VariantProductsPriceWithParent{}).
		Preload("VariantProduct.Product").
		Find(&variantProductPrices).Error

	if err != nil {
		fmt.Printf("err while fetching prices")
		return nil, err
	}
	// make prices template
	pricesXLSXTemplate := &aux.ProductPriceTemplate{}
	for _, price := range singleProductPrices {

		pricesXLSXTemplate.SingleProductPricesTemplate = append(
			pricesXLSXTemplate.SingleProductPricesTemplate,
			&model.SingleProductPriceTemplate{

				SingleProductID: price.SingleProduct.ID,
				Sku:             price.SingleProduct.Product.Sku,
				PriceName:       price.Name,
				PriceValue:      price.Value,
			},
		)

	}

	for _, price := range variantProductPrices {

		pricesXLSXTemplate.VariantProductPriceTemplate = append(
			pricesXLSXTemplate.VariantProductPriceTemplate,
			&model.VariantProductPriceTemplate{

				VariantProductID: price.VariantProduct.ID,
				Sku:              price.VariantProduct.Sku,
				PriceName:        price.Name,
				PriceValue:       price.Value,
			},
		)

	}
	return pricesXLSXTemplate, nil
}

func (r *ProductRepository) ImportPrices(priceTemplate *aux.ProductPriceTemplate) error {
	var err error
	tx := r.DB.Begin()

	if len(priceTemplate.SingleProductPricesTemplate) > 0 {
		singleProductPricesReq := model.SingleProductsPriceArr{}
		for _, sReq := range priceTemplate.SingleProductPricesTemplate {
			singleProductPricesReq = append(
				singleProductPricesReq,
				&model.SingleProductsPrice{
					SingleProductID: sReq.SingleProductID,
					Name:            sReq.PriceName,
					Value:           sReq.PriceValue,
				},
			)

		}
		// add validation here
		err = tx.Model(&model.SingleProductsPrice{}).Create(&singleProductPricesReq).Error
		if err != nil {
			fmt.Printf("error creating product_images\n %+v \n", err)
			tx.Rollback()
			return err
		}
	}

	if len(priceTemplate.VariantProductPriceTemplate) > 0 {
		fmt.Println(len(priceTemplate.VariantProductPriceTemplate))
		fmt.Println("salviro")
		variantProductPricesReq := model.VariantProductsPriceArr{}
		for _, sReq := range priceTemplate.VariantProductPriceTemplate {
			variantProductPricesReq = append(
				variantProductPricesReq,
				&model.VariantProductsPrice{
					VariantProductID: sReq.VariantProductID,
					Name:             sReq.PriceName,
					Value:            sReq.PriceValue,
				},
			)

		}

		// add validation here
		err = tx.Model(&model.VariantProductsPrice{}).Create(&variantProductPricesReq).Error
		if err != nil {
			fmt.Printf("error creating product_images\n %+v \n", err)
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}
