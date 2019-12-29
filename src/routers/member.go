package routers

import (
	"taurus/src/api"

	"github.com/labstack/echo"
)

func registerMember(e *echo.Echo) {
	g := e.Group("/member/")
	g.Use()
	g.POST("get-member-info", api.GetMemberInfo)
}
