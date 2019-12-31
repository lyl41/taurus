package api

//func AddMember(c echo.Context) (err error) {
//	req := new(table.Members)
//	if err = c.Bind(req); err != nil {
//		return
//	}
//
//	res, err := logic.GetMemberInfo(req.MchName, req.Phone, req.Name)
//	if err != nil {
//		return
//	}
//
//	return c.JSON(http.StatusOK, &CommonResp{
//		Code: 0,
//		Msg:  "",
//		Data: res,
//	})
//}
