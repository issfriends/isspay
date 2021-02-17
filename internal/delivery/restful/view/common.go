package view

type Meta struct {
	Pagination
}

type Pagination struct {
	Page       int64
	PerPage    int64
	Total      int64
	TotalPages int64
}

func (view *Pagination) SetTotals(page, perPage, total int64) {
	view.Page = page
	view.PerPage = perPage
	view.Total = total

	if perPage > 0 && total > 0 {
		view.TotalPages = total / perPage
		if total%perPage > 0 {
			view.TotalPages++
		}
	}
}

type DataList struct {
	Data interface{} `json:"data"`
	Meta `json:"meta"`
}
