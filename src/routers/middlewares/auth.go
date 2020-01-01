package middlewares

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"net/http"
	"taurus/src/common"
	"taurus/src/common/util/crypto"
	"time"

	"github.com/labstack/echo"
)

func MustAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		tokenMsg := c.Request().Header.Get(common.HeaderTokenName)
		if tokenMsg == "" {
			d := StdResp{
				Code: common.ErrCodeMustAuth,
				Msg:  "请先登录",
				Data: nil,
			}
			return c.JSON(http.StatusOK, d)
		}
		// AES解密token
		token, err := DecryptToken(tokenMsg)
		if err != nil {
			return c.JSON(http.StatusOK, StdResp{
				Code: common.ErrCodeMustAuth,
				Msg:  "token解析失败，请重新登录。err:" + err.Error(),
				Data: nil,
			})
		}
		if err = VerifyToken(token); err != nil {
			return c.JSON(http.StatusOK, StdResp{
				Code: common.ErrCodeMustAuth,
				Msg:  "请重新登录" + err.Error(),
				Data: nil,
			})
		}
		c.Set(common.HeaderTokenName, token)
		return next(c)
	}
}

func VerifyToken(token *common.Token) (err error) {
	if token == nil {
		return errors.New("token is nil")
	}
	if token.ExpiredTime < time.Now().Unix() {
		return errors.New("token过期")
	}
	return
}

func DecryptToken(tokenString string) (token *common.Token, err error) {
	byteInfo, e := hex.DecodeString(tokenString)
	if e != nil {
		err = e
		return
	}
	bathing, e := crypto.AesDecrypt(byteInfo, []byte(common.AesKEY))
	if e != nil {
		err = e
		return
	}
	token = new(common.Token)
	e = json.Unmarshal(bathing, token)
	if e != nil {
		err = e
		return
	}
	return
}
