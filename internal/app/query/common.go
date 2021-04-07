package query

type Pagination struct {
	Page    int64 `json:"page" query:"page"`
	PerPage int64 `json:"perPage" query:"perPage"`
}

type Sort struct {
	SortField string `json:"sortField" query:"sortField"`
	SortOrder string `json:"sortOrder" query:"sortOrder"`
}

type LockType int8

const (
	ReadLock LockType = iota + 1
	WriteLock
)
