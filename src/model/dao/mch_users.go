package dao

import (
	"taurus/src/common"
	"taurus/src/model/dao/table"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

func GetMchUsers(mchId int64, username, password string) (info *table.MchUsers, err error) {
	where := table.MchUsers{
		MchID:    mchId,
		Username: username,
		Password: password,
		Status:   common.StatusOk,
	}
	tmp := new(table.MchUsers)
	if err = GetDB().Where(where).First(tmp).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			err = nil
			return
		}
		err = errors.Wrap(err, "")
		return
	}
	info = tmp
	return
}
