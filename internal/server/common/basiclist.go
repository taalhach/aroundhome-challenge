package common

import (
	"fmt"
	"strings"

	"github.com/taalhach/aroundhome-challennge/pkg/items"
	"gorm.io/gorm"
)

const (
	eq        = "eq"
	neq       = "neq"
	gte       = "gte"
	gt        = "gt"
	lte       = "lte"
	lt        = "lt"
	startWith = "starts_with"
	endsWith  = "ends_with"
	between   = "between"
	contains  = "contains"
)

type BasicListRet struct {
	Page  int   `json:"page"`
	Pages int   `json:"pages"`
	Total int64 `json:"total"`
}

type Sort struct {
	By   string
	Desc bool
}

type BasicList struct {
	Limit     int      `query:"limit"`
	Page      int      `query:"page"`
	SortBy    string   `query:"sort_by"`
	SortOrder string   `query:"sort_order"`
	Filters   []string `query:"filters"`

	Columns map[string]string

	// pagination related
	SkipPagination bool

	// sort related
	SkipOrder bool
}

//PrepareSql molds query and applies filters and pagination
func (form *BasicList) PrepareSql(query *gorm.DB) (*gorm.DB, error) {
	if form.SortBy != "" {
		column, ok := form.Columns[form.SortBy]
		if !ok {
			return nil, fmt.Errorf("column mapping for %v not found", column)
		}

		sort := column
		if form.SortOrder != "" {
			sort = fmt.Sprintf("%s %s", column, form.SortOrder)
		} else {
			sort = fmt.Sprintf("%s", column)
		}

		query = query.Order(sort)
	}

	for _, filter := range form.Filters {
		subParts := strings.Split(filter, ":")
		if len(subParts) >= 3 {
			col := subParts[0]
			operation := subParts[1]
			val := subParts[2]

			column, ok := form.Columns[col]
			if !ok {
				return nil, fmt.Errorf("column mapping for %v not found", column)
			}

			switch strings.ToLower(operation) {
			case eq:
				query = query.Where(fmt.Sprintf("%s = ?", column), val)
			case gte:
				query = query.Where(fmt.Sprintf("%s >= ?", column), val)
			case gt:
				query = query.Where(fmt.Sprintf("%s > ?", column), val)
			case lte:
				query = query.Where(fmt.Sprintf("%s <= ?", column), val)
			case lt:
				query = query.Where(fmt.Sprintf("%s < ?", column), val)
			case neq:
				query = query.Where(fmt.Sprintf("%s <> ?", column), val)
			case startWith:
				query = query.Where(fmt.Sprintf("%s LIKE ?", column), fmt.Sprintf("%s%%", val))
			case endsWith:
				query = query.Where(fmt.Sprintf("%s LIKE ?", column), fmt.Sprintf("%%%s", val))
			case contains:
				query = query.Where(fmt.Sprintf("%s LIKE ?", column), fmt.Sprintf("%%%s%%", val))
			case between:
				//get second value
				var secVal string
				if len(subParts) > 3 {
					secVal = subParts[3]
				} else {
					secVal = val
				}
				query = query.Where(fmt.Sprintf("%s BETWEEN ? AND ?", column), val, secVal)
			}
		}
	}

	if !form.SkipPagination {
		query = form.Paginate(query)
	}

	return query, nil
}

//Paginate applies pagination if skipped in PrepareSql method
func (form *BasicList) Paginate(query *gorm.DB) *gorm.DB {
	// attach default pagination is limit and page are zero
	form.attachDefaultsPagination()
	// now paginate
	offset := (form.Page - 1) * form.Limit
	query = query.Limit(form.Limit).Offset(offset)

	return query
}

// attachDefaults attaches items.DefaultPaginationLimit and 1st page if no query params are passed by client
func (bl *BasicList) attachDefaultsPagination() {
	if bl.Limit == 0 {
		bl.Limit = items.DefaultPaginationLimit
	}

	if bl.Page == 0 {
		bl.Page = 1
	}
}
