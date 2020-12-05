package request

import (
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/helper/pagination"
)

type Preload []string

type FetchAllReq struct {
	Preload
	ProductIDArr []int               `json:"product_id_arr"`
	ItemIDArr    []int               `json:"item_id_arr"`
	Metadata     pagination.Metadata `json:"metadata"`
}
