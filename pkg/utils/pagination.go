package utils

import "math"

type Pagination struct {
	TotalCount  int `json:"total_count"`
	PageSize    int `json:"page_size"`
	CurrentPage int `json:"current_page"`
}

func (p *Pagination) TotalPages() int {
	return int(math.Ceil(float64(p.TotalCount) / float64(p.PageSize)))
}

func (p *Pagination) Offset() int {
	if p.CurrentPage < 1 {
		p.CurrentPage = 1
	}
	return (p.CurrentPage - 1) * p.PageSize
}

func (p *Pagination) PageRange() (start, end int) {
	start = p.Offset()
	end = start + p.PageSize
	if end > p.TotalCount {
		end = p.TotalCount
	}
	return
}
