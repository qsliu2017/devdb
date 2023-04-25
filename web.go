package devdb

import (
	"embed"

	"github.com/labstack/echo/v4"
)

//go:embed web/build/*
var assets embed.FS

func RegisterFrontend(e *echo.Echo) {
	fs := echo.MustSubFS(assets, "web/build")
	e.StaticFS("/", fs)
}
