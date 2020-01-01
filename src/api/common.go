package api

import "errors"

type Pagination struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

var (
	errParams = errors.New("参数有误，请仔细检查参数")
)
