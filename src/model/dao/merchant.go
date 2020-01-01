package dao

import (
	"taurus/src/model/dao/table"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

func GetMerchant(mchId int64, mchName string) (info *table.Merchant, err error) {
	where := table.Merchant{
		ID:      mchId,
		MchName: mchName,
	}
	tmp := table.Merchant{}
	if err = GetDB().Where(where).First(&tmp).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			err = nil
			return
		}
		err = errors.Wrap(err, "")
		return
	}
	info = &tmp
	return
}
