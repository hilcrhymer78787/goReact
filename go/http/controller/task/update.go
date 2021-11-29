package task

import (
	"sample/pkg/db"

	"github.com/labstack/echo"
)

// タスク登録
func Update(c echo.Context) (err error) {
	db := db.Connect()
	defer db.Close()

	id := c.QueryParam("id")
	name := c.QueryParam("name")

	db.Query(`UPDATE tasks set name = ? WHERE id = ?`, name, id)
	return nil
}
