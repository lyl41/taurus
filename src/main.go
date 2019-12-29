package main

import (
	"taurus/src/routers"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	routers.RegisterRouters(e)
	if err := e.Start(":8989"); err != nil {
		panic(err)
	}
}
