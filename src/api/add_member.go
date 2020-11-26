package api

import (
	"net/http"
	"taurus/src/logic"
	"taurus/src/model/dao/table"

	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

func AddMember(c echo.Context) (err error) {
	req := new(table.Members)
	if err = c.Bind(req); err != nil {
		return
	}
	if err = checkMemberInfo(req); err != nil {
		return
	}
	var data = []*table.Members{req}
	successCount, err := logic.AddMembers(data)
	if err != nil {
		return
	}
	return c.JSON(http.StatusOK, successCount)
}

func checkMemberInfo(req *table.Members) (err error) {
	if req == nil {
		err = errors.Wrap(err, "member info is nil")
		return
	}
	if req.MchName == "" {
		err = errors.New("商户号不能为空")
		return
	}
	if req.Phone == "" && req.Name == "" {
		err = errors.New("手机号和姓名请至少填写一个")
		return
	}
	return
}
