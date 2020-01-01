package logic

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"taurus/src/common"
	"taurus/src/common/util/crypto"
	"taurus/src/model/dao"
	"time"

	"github.com/pkg/errors"
)

func Login(merchant, username, password string) (tokenInfo string, err error) {
	mchInfo, err := dao.GetMerchant(0, merchant)
	if err != nil {
		return
	}
	if mchInfo == nil {
		err = fmt.Errorf("没有此商户%s", merchant)
		return
	}
	mchUsers, err := dao.GetMchUsers(mchInfo.ID, username, password)
	if err != nil {
		return
	}
	if mchUsers == nil {
		err = errors.New("用户名或密码错误")
		return
	}
	return genToken(int(mchInfo.ID), username, int(mchUsers.Role))
}

func genToken(mchId int, username string, role int) (tokenInfo string, err error) {
	token := common.Token{
		Username:    username,
		Role:        role,
		MchId:       mchId,
		ExpiredTime: time.Now().Add(time.Hour * 24).Unix(),
	}
	data, err := json.Marshal(token)
	if err != nil {
		err = errors.Wrap(err, "")
		return
	}
	encrypted, err := crypto.AesEncrypt(data, []byte(common.AesKEY))
	if err != nil {
		err = errors.Wrap(err, "")
		return
	}
	tokenInfo = hex.EncodeToString(encrypted)
	return
}
