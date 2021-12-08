package main

import (
	internal "hr_api/internal"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/search", internal.HRSearch)

	e.Logger.Fatal(e.Start(":5000"))
}

