package ordermodel

import (
	"distrodakwah_backend/app/database"
	"distrodakwah_backend/app/helper/sliceutils"
	invModel "distrodakwah_backend/app/services/inventory/model"
	prodModel "distrodakwah_backend/app/services/product/model"
)

type OrderItemVariantProduct struct {
	ID               int                   `gorm:"primaryKey;autoIncrement;not null"`
	OrderID          int                   `json:"order_id"`
	VariantProductID int                   `json:"variant_product_id"`
	VPInventory      *invModel.VPInventory `gorm:"foreignKey:VariantProductID" json:"vp_inventory"`
	Qty              int                   `json:"qty"`
	UnitWeight       int                   `json:"unit_weight"`
	// DropshipperItemPrice float64               `gorm:"type:decimal(19,2);not null;default:0.0" json:"dropshipper_item_price"`
	// RetailItemPrice      float64               `gorm:"type:decimal(19,2);not null;default:0.0" json:"retail_item_price"`
	Prices          *OrderItemVPPriceArr `gorm:"-" json:"prices"`
	OrderShippingID int                  `json:"order_shipping_id"`
}

func (OrderItemVariantProduct) ReturnModels() interface{} {
	return &OrderItemVariantProduct{}
}

func (oi OrderItemVariantProduct) ReturnCreate(order *OrderClass) (uint8, map[string]interface{}, error) {

	return prodModel.ProductKindVariant, map[string]interface{}{
		"OrderID":          order.ID,
		"VariantProductID": oi.VariantProductID,
		"Qty":              oi.Qty,
		"UnitWeight":       oi.UnitWeight,
		"OrderShippingID":  oi.OrderShippingID,
	}, nil

}

func (oi *OrderItemVariantProduct) SetOrderShippingID(id int) {
	(*oi).OrderShippingID = id
}

func (oi OrderItemVariantProduct) GetSellingPrice() (float64, bool) {
	sellingPrice, ok := oi.Prices.GetOrderItemPriceByName("jual")
	if !ok {
		return 0, false
	}

	return sellingPrice, true
}

func (oi OrderItemVariantProduct) GetVendorSubdistrictID() int {
	return oi.VPInventory.VPInventoryDetail.UserVendor.SubdistrictID
}

func (oi OrderItemVariantProduct) OrderItem() OrderItem {
	return OrderItem{
		Qty:    oi.Qty,
		ItemID: oi.VariantProductID,
	}
}

func (oi OrderItemVariantProduct) GetWeight() int {
	return oi.UnitWeight
}

func (oi OrderItemVariantProduct) PopulateData() error {
	var err error
	vProduct := &prodModel.VariantProduct{}
	err = database.DB.Model(&prodModel.VariantProduct{}).
		Where("id = ?", oi.VariantProductID).
		Find(&vProduct).Error
	if err != nil {
		return err
	}

	vpPrices := &prodModel.VariantProductsPriceArr{}
	err = database.DB.Model(&prodModel.VariantProductsPriceArr{}).
		Where("variant_product_id = ?", vProduct.ID).
		Find(&vpPrices).Error
	if err != nil {
		return err
	}

	vpInventory := &invModel.VPInventory{}
	err = database.DB.Model(&invModel.VPInventory{}).
		Where("variant_product_id = ?", vProduct.ID).
		Preload("VPInventoryDetail.UserVendor").
		Find(&vpInventory).Error
	if err != nil {
		return err
	}

	// populate price data
	prices := make([]*OrderItemVPPrice, len(*vpPrices))
	for i, vpPrice := range *vpPrices {
		prices[i] = &OrderItemVPPrice{
			Name:  vpPrice.Name,
			Value: vpPrice.Value,
		}
	}
	return nil
}

func (OrderItemVariantProduct) SubdistrictID(itemID int) (int, error) {
	VPInventory := &invModel.VPInventory{}
	err := database.DB.Model(&invModel.VPInventory{}).
		Where("variant_product_id = ?", itemID).
		Preload("VPInventoryDetail.UserVendor").
		Find(&VPInventory).
		Error

	if err != nil {
		return 0, err
	}

	return VPInventory.VPInventoryDetail.UserVendor.SubdistrictID, err
}

func (OrderItemVariantProduct) TableName() string {
	return "order_item_variant_products"
}

type OrderItemVariantProductArr []*OrderItemVariantProduct

func (ois OrderItemVariantProductArr) GetVendorSubdistrictIDArr() []int {

	orderShippingOriginIDs := []int{}
	for _, oi := range ois {
		orderShippingOriginIDs = append(orderShippingOriginIDs,
			oi.VPInventory.VPInventoryDetail.UserVendor.SubdistrictID,
		)
	}

	uniqueOrderShippingOriginIDs := sliceutils.UniqueInts(orderShippingOriginIDs)
	return uniqueOrderShippingOriginIDs
}

func (os *OrderItemVariantProductArr) PopulateDataArr() error {
	var err error
	for _, oi := range *os {
		vProduct := &prodModel.VariantProduct{}
		err = database.DB.Model(&prodModel.VariantProduct{}).
			Where("id = ?", oi.VariantProductID).
			Find(&vProduct).Error
		if err != nil {
			return err
		}

		vpPrices := &prodModel.VariantProductsPriceArr{}
		err = database.DB.Model(&prodModel.VariantProductsPriceArr{}).
			Where("variant_product_id = ?", vProduct.ID).
			Find(&vpPrices).Error
		if err != nil {
			return err
		}

		vpInventory := &invModel.VPInventory{}
		err = database.DB.Model(&invModel.VPInventory{}).
			Where("variant_product_id = ?", vProduct.ID).
			Preload("VPInventoryDetail.UserVendor").
			Find(&vpInventory).Error
		if err != nil {
			return err
		}

		// populate price data

		prices := make([]*OrderItemVPPrice, len(*vpPrices))
		for i, vpPrice := range *vpPrices {
			prices[i] = &OrderItemVPPrice{
				Name:  vpPrice.Name,
				Value: vpPrice.Value,
			}
		}
		oi.VPInventory = vpInventory
		oi.UnitWeight = vProduct.Weight
		*oi.Prices = prices
		// for _, price := range *oi.Prices {
		// 	fmt.Printf("fajar sidiq salviro[sp] %+v \n", price)
		// }
	}
	return nil
}
