package ordermodel

type OrderItem struct {
	// OrderID uint64
	Qty int
	// KindID uint8
	ItemID uint64
}

type OrderItemArr []*OrderItem

type OrderItemI interface {
	OrderItem() OrderItem
	GetVendorSubdistrictID() int
	GetSellingPrice() (float64, bool)
	GetWeight() int
	SetOrderShippingID(id uint64)
	ReturnCreate(order *OrderClass) (uint8, map[string]interface{}, error)
	ReturnModels() interface{}
}

// type OrderItemArrI interface {
// 	GetVendorSubdistrictIDArr() []int
// }

type OrderItemReq struct {
	Qty    int
	KindID uint8
	ItemID uint64
}

type OrderItemReqArr []*OrderItemReq

// type OrderItemWithItemID struct {
// 	OrderID uint64
// 	Qty     int
// 	// UnitWeight           int
// 	KindID uint8
// 	ItemID uint64
// 	// DropshipperItemPrice float64
// 	// RetailItemPrice      float64
// 	// OrderShippingID      uint64
// 	// SPInventory          *invModel.SPInventory `gorm:"foreignKey:SingleProductID;references:ItemID"`
// 	// VPInventory          *invModel.VPInventory `gorm:"foreignKey:VariantProductID;references:ItemID"`
// }

// func (orderItem *OrderItemWithItemID) ReturnModels() interface{} {
// 	if orderItem.KindID == 1 {
// 		return &OrderItemSingleProduct{}
// 	}
// 	return &OrderItemVariantProduct{}
// }

// func (orderItem *OrderItemWithItemID) ReturnCreate(order *OrderReq) (map[string]interface{}, error) {
// 	if orderItem.KindID == 1 {
// 		// singleProduct := &prodModel.SingleProduct{}
// 		// err = database.DB.Model(&prodModel.SingleProduct{}).
// 		// 	Where("id = ?", orderItem.ItemID).
// 		// 	Find(&singleProduct).Error
// 		// if err != nil {
// 		// 	return nil, err
// 		// }

// 		// spPrices := &prodModel.SingleProductsPriceArr{}
// 		// err = database.DB.Model(&prodModel.SingleProductsPrice{}).
// 		// 	Where("single_product_id = ?", singleProduct.ID).
// 		// 	Find(&spPrices).Error
// 		// if err != nil {
// 		// 	return nil, err
// 		// }

// 		// resPrice, ok := spPrices.GetSPPriceByName("reseller pro")
// 		// retailPrice, ok := spPrices.GetSPPriceByName("jual")

// 		// if !ok {
// 		// 	return nil, errors.New("price not found")
// 		// }
// 		return map[string]interface{}{
// 			"OrderID":              order.ID,
// 			"SingleProductID":      orderItem.ItemID,
// 			"Qty":                  orderItem.Qty,
// 			"UnitWeight":           orderItem.UnitWeight,
// 			"DropshipperItemPrice": orderItem.DropshipperItemPrice,
// 			"RetailItemPrice":      orderItem.RetailItemPrice,
// 			"OrderShippingID":      orderItem.OrderShippingID,
// 		}, nil
// 	} else if orderItem.KindID == 2 {
// 		// variantProduct := &prodModel.VariantProduct{}
// 		// err := database.DB.Model(&prodModel.VariantProduct{}).
// 		// 	Where("id = ?", orderItem.ItemID).
// 		// 	Find(&variantProduct).Error
// 		// if err != nil {
// 		// 	return nil, errors.New("variant product not found")
// 		// }
// 		// vpPrices := &prodModel.VariantProductsPriceArr{}
// 		// err = database.DB.Model(&prodModel.VariantProductsPrice{}).
// 		// 	Where("variant_product_id = ?", variantProduct.ID).
// 		// 	Find(&vpPrices).Error
// 		// if err != nil {
// 		// 	return nil, err
// 		// }

// 		// resPrice, ok := vpPrices.GetVPPriceByName("reseller pro")
// 		// retailPrice, ok := vpPrices.GetVPPriceByName("jual")
// 		// if !ok {
// 		// 	return nil, errors.New("price not found")
// 		// }

// 		return map[string]interface{}{
// 			"OrderID":              order.ID,
// 			"VariantProductID":     orderItem.ItemID,
// 			"Qty":                  orderItem.Qty,
// 			"UnitWeight":           orderItem.UnitWeight,
// 			"DropshipperItemPrice": orderItem.DropshipperItemPrice,
// 			"RetailItemPrice":      orderItem.RetailItemPrice,
// 			"OrderShippingID":      orderItem.OrderShippingID,
// 		}, nil
// 	}

// 	return nil, errors.New("couldn't create query")
// }

// type OrderItemI interface {
// 	ReturnModels() interface{}
// 	ReturnCreate(order *OrderReq) (map[string]interface{}, error)
// 	GetProductSubdistrictID() (int, bool)
// }

// func

// type OrderItemI interface {
// 	SubdistrictID() (int, error)
// }
// type OrderItemArrI []*OrderItem

// func (ois *OrderItemArrI) PopulateData() error {
// 	var err error

// }

// func (orderItem *OrderItemWithItemID) GetProductSubdistrictID() (int, bool) {
// 	if orderItem.KindID == 1 {

// 		// SPInventory := &invModel.SPInventory{}
// 		// err := database.DB.Model(&invModel.SPInventory{}).
// 		// 	Where("single_product_id = ?", orderItem.ItemID).
// 		// 	Preload("SPInventoryDetail.UserVendor").
// 		// 	Find(&SPInventory).
// 		// 	Error

// 		// if err != nil {
// 		// 	return 0, false
// 		// }

// 		// return SPInventory.SPInventoryDetail.UserVendor.SubdistrictID, true

// 	} else if orderItem.KindID == 2 {
// 		VPInventory := &invModel.VPInventory{}

// 		err := database.DB.Model(&invModel.VPInventory{}).
// 			Where("variant_product_id = ?", orderItem.ItemID).
// 			Preload("VPInventoryDetail.UserVendor").
// 			Find(&VPInventory).
// 			Error

// 		if err != nil {
// 			return 0, false
// 		}
// 		return VPInventory.VPInventoryDetail.UserVendor.SubdistrictID, true

// 	}
// 	return 0, false
// }

// func (orderItem *OrderItemWithItemID) GetProduct
