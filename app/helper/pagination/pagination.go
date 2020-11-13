package pagination

const (
	defaultPage  = 1
	defaultLimit = 10
)

type Metadata struct {
	Total  int64 `json:"total"`
	Limit  int   `json:"limit"`
	Offset int   `json:"offset"`
	Page   int   `json:"page"`
	Pages  int   `json:"pages"`
}

func (m *Metadata) UpdateTotal(t int64) {
	m.Total = t
}

type Pagination struct {
	Elements interface{} `json:"elements"`
	Metadata *Metadata   `json:"metadata"`
}

func (p *Pagination) UpdateElements(e interface{}) {
	p.Elements = e
}

// return page, limit, offset,
func BuildPagination(m *Metadata) (int, int, int) {
	page := (*m).Page
	if page < 1 {
		page = defaultPage
	}

	limit := (*m).Limit
	if limit == 0 {
		limit = defaultLimit
	}

	offset := (limit * page) - limit

	return page, limit, offset
}

func BuildPages(elems int64, limit int) int {
	var pages int
	total := int(elems)
	if pages = (total / limit); (total % limit) != 0 {
		pages++
	}
	return pages
}
