package routers

import (
	"taurus/src/api"

	"github.com/labstack/echo"
)

func registerSendSms(e *echo.Echo) {
	g := e.Group("/send-sms/")
	g.Use()
	g.POST("get-member-info", api.GetMemberInfo)
}
