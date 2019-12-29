package middlewares

import (
	"fmt"

	"github.com/labstack/echo"
)

func Recover(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = fmt.Errorf("系统繁忙，请稍后再试~")
				fmt.Println("panic err", r)
			}
		}()
		return next(c)
	}
}
