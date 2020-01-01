package api

import (
	"net/http"
	"taurus/src/logic"

	"github.com/labstack/echo"
)

type LoginReq struct {
	Merchant string `json:"merchant"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c echo.Context) (err error) {
	req := new(LoginReq)
	if err = c.Bind(req); err != nil {
		return
	}
	if req.Merchant == "" || req.Username == "" || req.Password == "" {
		return errParams
	}
	token, err := logic.Login(req.Merchant, req.Username, req.Password)
	if err != nil {
		return
	}

	return c.JSON(http.StatusOK, token)
}
