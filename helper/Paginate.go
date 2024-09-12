package helper

import (
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

type Pagination struct {
    Page     int
    PageSize int
}
func Paginate(r *http.Request) (func(db *gorm.DB) *gorm.DB, Pagination) {
    q := r.URL.Query()
    page, _ := strconv.Atoi(q.Get("page"))
    if page <= 0 {
        page = 1
    }

    pageSize, _ := strconv.Atoi(q.Get("page_size"))
    switch {
    case pageSize > 100:
        pageSize = 100
    case pageSize <= 0:
        pageSize = 15
    }

    offset := (page - 1) * pageSize

    pagination := Pagination{
        Page:     page,
        PageSize: pageSize,
    }
    return func(db *gorm.DB) *gorm.DB {
        return db.Offset(offset).Limit(pageSize)
    }, pagination
}
