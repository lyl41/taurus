package routers

import (
	"taurus/src/routers/middlewares"

	"github.com/labstack/echo"
)

func RegisterRouters(e *echo.Echo) {
	e.Use(middlewares.HookCtx, middlewares.Logger, middlewares.Recover)
	registerMember(e)
	registerSendSms(e)
}
