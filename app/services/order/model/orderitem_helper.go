package model

// func (ois *OrderItemReqArr) PopulateData() error {
// 	var err error
// 	for _, oi := range *ois {
// 		if oi.KindID == 1 {
// 			singleProduct := &prodModel.SingleProduct{}
// 			err = database.DB.Model(&prodModel.SingleProduct{}).
// 				Where("id = ?", oi.ItemID).
// 				Find(&singleProduct).Error
// 			if err != nil {
// 				return err
// 			}

// 			spPrices := &prodModel.SingleProductsPriceArr{}
// 			err = database.DB.Model(&prodModel.SingleProductsPrice{}).
// 				Where("single_product_id = ?", singleProduct.ID).
// 				Find(&spPrices).Error
// 			if err != nil {
// 				return err
// 			}

// 			resPrice, ok := spPrices.GetSPPriceByName("reseller pro")
// 			retailPrice, ok := spPrices.GetSPPriceByName("jual")

// 			if !ok {
// 				return errors.New("price not found")
// 			}
// 			// populated data
// 			oi.DropshipperItemPrice = resPrice
// 			oi.RetailItemPrice = retailPrice
// 			oi.UnitWeight = singleProduct.Weight

// 		} else if oi.KindID == 2 {
// 			variantProduct := &prodModel.VariantProduct{}
// 			err := database.DB.Model(&prodModel.VariantProduct{}).
// 				Where("id = ?", oi.ItemID).
// 				Find(&variantProduct).Error
// 			if err != nil {
// 				return errors.New("variant product not found")
// 			}
// 			vpPrices := &prodModel.VariantProductsPriceArr{}
// 			err = database.DB.Model(&prodModel.VariantProductsPrice{}).
// 				Where("variant_product_id = ?", variantProduct.ID).
// 				Find(&vpPrices).Error
// 			if err != nil {
// 				return err
// 			}

// 			resPrice, ok := vpPrices.GetVPPriceByName("reseller pro")
// 			retailPrice, ok := vpPrices.GetVPPriceByName("jual")

// 			if !ok {
// 				return errors.New("price not found")
// 			}
// 			// populated data

// 			oi.DropshipperItemPrice = resPrice
// 			oi.RetailItemPrice = retailPrice
// 			oi.UnitWeight = variantProduct.Weight
// 		}
// 	}
// 	return nil
// }

// func (ois *OrderItemWithItemIDs) SumDropshipperItemPrice() float64 {
// 	var tempTotal float64
// 	for _, oi := range *ois {
// 		tempTotal += float64(oi.Qty) * oi.DropshipperItemPrice
// 	}
// 	return tempTotal

// }
