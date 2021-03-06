package ordermodel

// type OrderItemSingleProduct struct {
// 	ID              int                `gorm:"primaryKey;autoIncrement;not null"`
// 	OrderID         int                `json:"order_id"`
// 	SingleProductID int                `json:"single_product_id"`
// 	SPInventory     *invModel.SPInventory `gorm:"foreignKey:SingleProductID" json:"sp_inventory"`
// 	Qty             int                   `json:"qty"`
// 	UnitWeight      int                   `json:"unit_weight"`
// 	// DropshipperItemPrice float64               `gorm:"type:decimal(19,2);not null;default:0.0" json:"dropshipper_item_price"`
// 	// RetailItemPrice      float64               `gorm:"type:decimal(19,2);not null;default:0.0" json:"retail_item_price"`
// 	Prices          *OrderItemSPPriceArr `gorm:"-" json:"prices"`
// 	OrderShippingID int               `json:"order_shipping_id"`
// }

// func (OrderItemSingleProduct) ReturnModels() interface{} {
// 	return &OrderItemSingleProduct{}
// }

// func (oi OrderItemSingleProduct) ReturnCreate(order *OrderClass) (uint8, map[string]interface{}, error) {
// 	return prodModel.ProductKindSingle, map[string]interface{}{
// 		"OrderID":         order.ID,
// 		"SingleProductID": oi.SingleProductID,
// 		"Qty":             oi.Qty,
// 		"UnitWeight":      oi.UnitWeight,
// 		"OrderShippingID": oi.OrderShippingID,
// 	}, nil

// }

// func (oi *OrderItemSingleProduct) SetOrderShippingID(id int) {
// 	(*oi).OrderShippingID = id
// }

// func (oi OrderItemSingleProduct) GetVendorSubdistrictID() int {
// 	return oi.SPInventory.SPInventoryDetail.UserVendor.SubdistrictID
// }

// func (oi OrderItemSingleProduct) OrderItem() OrderItem {
// 	return OrderItem{
// 		Qty:    oi.Qty,
// 		ItemID: oi.SingleProductID,
// 	}
// }

// func (oi OrderItemSingleProduct) GetSellingPrice() (float64, bool) {
// 	dropshipPrice, ok := oi.Prices.GetOrderItemPriceByName("jual")
// 	if !ok {
// 		return 0, false
// 	}

// 	return dropshipPrice, true
// }

// func (oi OrderItemSingleProduct) GetWeight() int {
// 	return oi.UnitWeight
// }

// func (oi OrderItemSingleProduct) PopulateData() error {
// 	var err error
// 	sProduct := &prodModel.SingleProduct{}
// 	err = database.DB.Model(&prodModel.SingleProduct{}).
// 		Where("id = ?", oi.SingleProductID).
// 		Find(&sProduct).Error
// 	if err != nil {
// 		return err
// 	}

// 	spPrices := &prodModel.SingleProductsPriceArr{}
// 	err = database.DB.Model(&prodModel.SingleProductsPrice{}).
// 		Where("single_product_id = ?", sProduct.ID).
// 		Find(&spPrices).Error
// 	if err != nil {
// 		return err
// 	}

// 	spInventory := &invModel.SPInventory{}
// 	err = database.DB.Model(&invModel.SPInventory{}).
// 		Where("single_product_id = ?", sProduct.ID).
// 		Preload("SPInventory.SPInventoryDetail.UserVendor").
// 		Find(&spInventory).Error
// 	if err != nil {
// 		return err
// 	}

// 	// populate price data
// 	prices := make([]*OrderItemSPPrice, len(*spPrices))
// 	for i, spPrice := range *spPrices {
// 		(*prices[i]).Name = spPrice.Name
// 		(*prices[i]).Name = spPrice.Name
// 	}
// 	return nil
// }

// func (OrderItemSingleProduct) SubdistrictID(itemID int) (int, error) {
// 	SPInventory := &invModel.SPInventory{}
// 	err := database.DB.Model(&invModel.SPInventory{}).
// 		Where("single_product_id = ?", itemID).
// 		Preload("SPInventoryDetail.UserVendor").
// 		Find(&SPInventory).
// 		Error

// 	if err != nil {
// 		return 0, err
// 	}

// 	return SPInventory.SPInventoryDetail.UserVendor.SubdistrictID, err
// }

// func (OrderItemSingleProduct) TableName() string {
// 	return "order_item_single_products"
// }

// type OrderItemSingleProductArr []*OrderItemSingleProduct

// func (ois OrderItemSingleProductArr) GetVendorSubdistrictIDArr() []int {

// 	orderShippingOriginIDs := []int{}
// 	for _, oi := range ois {
// 		orderShippingOriginIDs = append(orderShippingOriginIDs,
// 			oi.SPInventory.SPInventoryDetail.UserVendor.SubdistrictID,
// 		)
// 	}

// 	uniqueOrderShippingOriginIDs := sliceutils.UniqueInts(orderShippingOriginIDs)
// 	return uniqueOrderShippingOriginIDs
// }

// func (ois *OrderItemSingleProductArr) PopulateDataArr() error {
// 	var err error
// 	// create temp oirderitem
// 	for _, oi := range *ois {
// 		sProduct := &prodModel.SingleProduct{}
// 		err = database.DB.Model(&prodModel.SingleProduct{}).
// 			Where("id = ?", oi.SingleProductID).
// 			Find(&sProduct).Error
// 		if err != nil {
// 			return err
// 		}

// 		spPrices := &prodModel.SingleProductsPriceArr{}
// 		err = database.DB.Model(&prodModel.SingleProductsPrice{}).
// 			Where("single_product_id = ?", sProduct.ID).
// 			Find(&spPrices).Error
// 		if err != nil {
// 			return err
// 		}

// 		spInventory := &invModel.SPInventory{}

// 		err = database.DB.Model(&invModel.SPInventory{}).
// 			Where("single_product_id = ?", sProduct.ID).
// 			Preload("SPInventoryDetail.UserVendor").
// 			Find(&spInventory).Error

// 		if err != nil {
// 			return err
// 		}

// 		// populate price data
// 		prices := make([]*OrderItemSPPrice, len(*spPrices))
// 		for i, spPrice := range *spPrices {
// 			prices[i] = &OrderItemSPPrice{
// 				Name:  spPrice.Name,
// 				Value: spPrice.Value,
// 			}

// 		}

// 		oi.SPInventory = spInventory
// 		oi.UnitWeight = sProduct.Weight

// 		*oi.Prices = prices
// 		// for _, price := range *oi.Prices {
// 		// 	fmt.Printf("fajar sidiq salviro[sp] %+v \n", price)
// 		// }

// 	}
// 	return nil
// }
