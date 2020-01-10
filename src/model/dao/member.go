package dao

import (
	"taurus/src/common"
	"taurus/src/model/dao/table"

	"github.com/pkg/errors"
)

func GetMemberInfo(merchant, phone, name string, offset, limit int) (info []*table.Members, err error) {
	where := table.Members{
		MchName: merchant,
		Phone:   phone,
		Status:  common.StatusOk,
	}
	db := GetDB().Where(where)
	if name != "" {
		db = db.Where("name like ?", name+"%")
	}
	if err = db.Offset(offset).Limit(limit).Find(&info).Error; err != nil {
		err = errors.Wrap(err, "")
		return
	}
	return
}

func InsertMember(memberInfo *table.Members) (err error) {
	if err = GetDB().Create(memberInfo).Error; err != nil {
		err = errors.Wrap(err, "insert member err")
		return
	}
	return
}
