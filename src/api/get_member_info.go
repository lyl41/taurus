package api

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func GetMemberInfo(c echo.Context) (err error) {
	time.Sleep(time.Millisecond * 12)
	return c.JSON(http.StatusOK, "hello world!!")
}
