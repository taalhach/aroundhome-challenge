package forms

import "github.com/taalhach/aroundhome-challennge/pkg/items"

type BasicResponse struct {
	Success bool              `json:"success"`
	Message string            `json:"message"`
	Errors  map[string]string `json:"errors"`
}
type BasicList struct {
	Limit     int      `query:"limit"`
	Page      int      `query:"page"`
	SortBy    string   `query:"sort_by"`
	SortOrder string   `query:"sort_order"`
	Filters   []string `query:"filters"`
}

// AttachDefaults attaches items.DefaultPaginationLimit and 1st page if no query params are passed by client
func (bl *BasicList) AttachDefaults() {
	if bl.Limit == 0 {
		bl.Limit = items.DefaultPaginationLimit
	}

	if bl.Page == 0 {
		bl.Page = 1
	}
}
