package healthcheck

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Init will creat and attach all REST API route to given echo instance
func Init(e *echo.Echo) {
	e.GET("/", check)
}

func check(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}
