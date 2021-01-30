package inventoryhandler

import (
	"distrodakwah_backend/app/helper/pagination"
)

type Preload []string

type FetchAllReq struct {
	Preload
	ProductIDArr []uint              `json:"product_id_arr"`
	ItemIDArr    []uint              `json:"item_id_arr"`
	Metadata     pagination.Metadata `json:"metadata"`
}
