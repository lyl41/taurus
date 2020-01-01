package routers

import (
	"taurus/src/api"
	"taurus/src/routers/middlewares"

	"github.com/labstack/echo"
)

func registerMember(e *echo.Echo) {
	g := e.Group("/member/")
	g.Use()
	g.POST("login", api.Login)
	g.POST("get-member-info", api.GetMemberInfo, middlewares.MustAuth)
}
