package producthandler

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/helper/pagination"
)

type Preload []string

type CategoryReq struct {
	SubdepartmentID uint   `json:"subdepartment_id"`
	Name            string `json:"name"`
	Image           string `json:"image"`
	ParentID        uint   `json:"parent_id,omitempty"`
}

type BrandReq struct {
	UserVendorID uint32 `json:"user_vendor_id"`
	Name         string `json:"name"`
	Image        string `json:"image"`
}

type DepartmentReq struct {
	Name  string `json:"name"`
	Image string `json:"image"`
}

type SubdepartmentReq struct {
	DepartmentID uint8  `json:"department_id"`
	Name         string `json:"name"`
	Image        string `json:"image"`
}

type FetchAllReq struct {
	Preload
	Metadata     pagination.Metadata `json:"metadata"`
	ProductIDArr []int               `json:"product_id_arr"`
}

type FetchByColumnReq struct {
	PKindIDs []int
	PTypeIDs []int
	Metadata pagination.Metadata `json:"metadata"`
	Preload
}

func (req *FetchByColumnReq) Mydecode(urlVal url.Values) error {
	if pKindIDReq, ok := urlVal["product_kind_id"]; ok && len(urlVal["product_kind_id"]) > 0 {
		fmt.Println(pKindIDReq[0])
		err := json.NewDecoder(strings.NewReader(pKindIDReq[0])).Decode(&req.PKindIDs)
		if err != nil {
			return err
		}
	}
	if pTypeIdRq, ok := urlVal["product_type_id"]; ok && len(urlVal["product_type_id"]) > 0 {
		err := json.NewDecoder(strings.NewReader(pTypeIdRq[0])).Decode(&req.PTypeIDs)
		if err != nil {
			return err
		}
	}
	return nil
}
