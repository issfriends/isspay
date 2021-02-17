package scope

import (
	"fmt"

	"github.com/issfriends/isspay/internal/app/query"
	"gorm.io/gorm"
)

func Pagination(q query.Pagination) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page, perPage := q.Page, q.PerPage
		if page > 0 && perPage > 0 {
			return db.Offset(int((page - 1) * perPage)).Limit(int(perPage))
		}

		return db
	}
}

func Sort(q query.Sort) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		field, order := q.SortField, q.SortOrder
		if order == "" || (order != "ASC" && order != "DESC") {
			order = "DESC"
		}
		if field != "" && order != "" {
			return db.Order(fmt.Sprintf("%s %s", field, order))
		}
		return db
	}
}
