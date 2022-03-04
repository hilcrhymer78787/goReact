package test

import (
	"net/http"

	"github.com/labstack/echo"
)

// テスト
func Test(c echo.Context) error {
	return c.JSON(http.StatusOK, "test success")
}
