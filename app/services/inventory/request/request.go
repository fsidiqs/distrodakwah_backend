package request

import (
	"github.com/zakiyfadhilmuhsin/distrodakwah_backend/app/helper/pagination"
)

type Preload []string

type FetchAllReq struct {
	Preload
	Metadata     *pagination.Metadata `json:"metadata"`
	ProductIDArr []int                `json:"product_id_arr"`
	SPIDArr      []int                `json:"sp_id_arr"`
	VPIDArr      []int                `json:"vp_id_arr"`
}