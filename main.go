package main

import (
	"pair-co-slot-api/healthcheck"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	healthcheck.Init(e)
	e.Logger.Fatal(e.Start(":1323"))
}
