package api

import (
	"net/http"
	"taurus/src/logic"
	"taurus/src/model/dao/table"

	"github.com/labstack/echo"
)

type GetMemberInfoReq struct {
	Merchant string `json:"merchant"`
	Phone    string `json:"phone"`
	Name     string `json:"name"`
	Pagination
}
type GetMemberInfoResp struct {
	Data []*table.Members
}

func GetMemberInfo(c echo.Context) (err error) {
	req := new(GetMemberInfoReq)
	if err = c.Bind(req); err != nil {
		return
	}

	res, err := logic.GetMemberInfo(req.Merchant, req.Phone, req.Name, req.Page, req.PageSize)
	if err != nil {
		return
	}

	return c.JSON(http.StatusOK, res)
}
