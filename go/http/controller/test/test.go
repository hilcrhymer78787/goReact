package test

import (
	"net/http"

	"github.com/labstack/echo"
)

// ใในใ
func Test(c echo.Context) error {
	return c.JSON(http.StatusOK, "test success")
}
