package main

import (
	"github.com/labstack/echo/v4"
	"github.com/qsliu2017/devdb"
)

func main() {
	e := echo.New()

	devdb.RegisterFrontend(e)

	if err := e.Start(":8081"); err != nil {
		panic(err)
	}
}
