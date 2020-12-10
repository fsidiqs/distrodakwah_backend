package ordermodel

import (
	"errors"

	"distrodakwah_backend/app/database"
	custModel "distrodakwah_backend/app/services/user/model"
)

type OrderCustomer struct {
	ID            uint64 `gorm:"primaryKey;autoIncrement;not null"`
	OrderID       uint64 `json:"order_id"`
	CustomerID    uint64 `json:"customer_id"`
	Name          string `gorm:"type:varchar(255);not null" json:"name"`
	AddressDetail string `json:"address_detail"`
	SubdistrictID uint64 `json:"subdistrict_id"`
	PostalCode    string `json:"postal_code"`
	Phone         string `json:"phone"`
	Email         string `jsonn:"email"`
}

func (oc *OrderCustomerDetail) PopulateData() error {
	customer := &custModel.Customer{}
	err := database.DB.Model(&custModel.Customer{}).
		Where("id = ?", oc.CustomerID).
		Find(&customer).Error
	if err != nil {
		return errors.New("Cant f")
	}
	oc.SubdistrictID = customer.SubdistrictID
	oc.AddressDetail = customer.AddressDetail
	oc.Email = customer.Email
	oc.Phone = customer.Phone
	oc.Name = customer.Name
	oc.PostalCode = customer.PostalCode

	return nil
}

func (oc *OrderCustomerDetail) UpdateOrderID(id uint64) {
	oc.OrderID = id
}
