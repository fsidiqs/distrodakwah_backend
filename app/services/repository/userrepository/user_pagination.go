package userrepository

import "distrodakwah_backend/app/helper/pagination"

const (
	defaultPage  = 1
	defaultLimit = 10
)

type Pagination struct {
	Elements interface{}         `json:"elements"`
	Metadata pagination.Metadata `json:"metadata"`
}

func (p *Pagination) UpdateElements(e interface{}) {
	p.Elements = e
}
