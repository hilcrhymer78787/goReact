package task

import (
	"sample/pkg/db"

	"github.com/labstack/echo"
)

// タスク登録
func Create(c echo.Context) (err error) {
	db := db.Connect()
	defer db.Close()

	name := c.QueryParam("name")

	db.Query(`INSERT INTO tasks (name) VALUES (?)`, name)
	return nil
}
