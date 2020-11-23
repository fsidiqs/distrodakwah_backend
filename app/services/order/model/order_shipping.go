package model

import (
	"errors"

	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/helper/sliceutils"
)

type OrderShipping struct {
	ID                  uint64  `gorm:"primaryKey;autoIncrement;not null"`
	OrderID             uint64  `json:"order_id"`
	ShippingCost        float64 `gorm:"type:decimal(19,2);not null;default:0.0" json:"shipping_cost"`
	TotalCost           float64 `gorm:"type:decimal(19,2);not null;default:0.0" json:"total_cost"`
	Weight              int     `gorm:"not null" json:"weight"`
	ShippingCompanyID   uint8   `gorm:"UNSIGNED;NOT NULL" json:"shipping_company_id"`
	ShippingServiceName string  `json:"shipping_service_name"`
	SubdistrictIDOrigin int     `gorm:"not null" json:"subdistrict_id_origin"`
	Awb                 string  `json:"awb"`
	TypeID              uint8   `gorm:"not null;unsigned" json:"type_id"`
}

type OrderShippings []*OrderShipping

func (os OrderShippings) FindOrderShippingIndexByOriginID(originID int) (int, bool) {
	for i := range os {
		if os[i].SubdistrictIDOrigin == originID {
			return i, true
		}
	}
	return 0, false
}

func (OrderShippings) PrepareShippings(orderID uint64, orderItemIs []OrderItemI) (*OrderShippings, error) {
	// var orderShipping []*OrderShipping
	orderShippingArr := &OrderShippings{}

	orderShippingOriginIDs := []int{}
	for _, orderItemI := range orderItemIs {

		// test, _ := orderItemI.GetSellingPrice()
		// fmt.Printf("fajar sidiq salviro[%+v] %+v \n", i, test)
		orderShippingOriginIDs = append(orderShippingOriginIDs, orderItemI.GetVendorSubdistrictID())

	}

	uniqueOrderShippingOriginIDs := sliceutils.UniqueInts(orderShippingOriginIDs)
	// make orderShipping

	orderShipping := make([]*OrderShipping, len(uniqueOrderShippingOriginIDs))

	*orderShippingArr = OrderShippings(orderShipping)

	for i, originID := range uniqueOrderShippingOriginIDs {
		// shipdetails := rajaongkir.ShippingDetails{
		// 	OriginSubID: osid,
		// 	// ShipName: ,
		// }
		// abc := rajaongkir.GetCost()
		(*orderShippingArr)[i] = &OrderShipping{
			SubdistrictIDOrigin: originID,
			TypeID:              1,
			// calculate shipping cost
			ShippingCost:        5,
			ShippingCompanyID:   1,
			ShippingServiceName: "asd",
			OrderID:             orderID,
		}
	}

	for _, oi := range orderItemIs {
		// orderItemInterface = oi

		originSubID := oi.GetVendorSubdistrictID()

		ordershippingID, ok := orderShippingArr.FindOrderShippingIndexByOriginID(originSubID)
		if !ok {
			return nil, errors.New("Error Order Shippings")
		}
		itemSellingPrice, ok := oi.GetSellingPrice()
		if !ok {
			return nil, errors.New("Error Order Item Price")
		}

		(*orderShippingArr)[ordershippingID].TotalCost += float64(oi.OrderItem().Qty) * itemSellingPrice
		(*orderShippingArr)[ordershippingID].Weight += (oi.OrderItem().Qty) * oi.GetWeight()
	}

	return orderShippingArr, nil
}
