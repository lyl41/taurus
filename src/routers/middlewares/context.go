package middlewares

import "github.com/labstack/echo"

func HookCtx(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c = &echoContext{c}
		return next(c)
	}
}

type echoContext struct {
	echo.Context
}

func (e *echoContext) SetOutResponse(i interface{}) {
	e.Set("set-out-response", i)
}

func GetOutResponse(c echo.Context) interface{} {
	return c.Get("set-out-response")
}

func (e *echoContext) Bind(i interface{}) error {
	e.SetOutResponse(i)
	return e.Context.Bind(i)
}

func (e *echoContext) JSON(code int, i interface{}) error {
	e.SetOutResponse(i)
	return e.Context.JSON(code, i)
}
func (e *echoContext) String(code int, s string) error {
	e.SetOutResponse(s)
	return e.Context.String(code, s)
}
func (e *echoContext) HTML(code int, html string) error {
	e.SetOutResponse(html)
	return e.Context.HTML(code, html)
}
func (e *echoContext) XML(code int, i interface{}) error {
	e.SetOutResponse(i)
	return e.Context.XML(code, i)
}
