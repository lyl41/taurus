package logic

import (
	"taurus/src/model/dao"
	"taurus/src/model/dao/table"
)

func GetMemberInfo(merchant, phone, name string, page, pageSize int) (infos []*table.Members, err error) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = defaultPageSize
	}
	if pageSize > maxLimitPageSize {
		pageSize = maxLimitPageSize
	}
	offset := (page - 1) * pageSize
	infos, err = dao.GetMemberInfo(merchant, phone, name, offset, pageSize)
	if err != nil {
		return
	}
	return
}
