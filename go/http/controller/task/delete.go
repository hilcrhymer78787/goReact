package task

import (
	"sample/pkg/db"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
)

// タスク削除
func Delete(c echo.Context) error {
	db := db.Connect()
	defer db.Close()

	id := c.QueryParam("id")

	db.Query(`DELETE FROM tasks WHERE id = (?)`, id)
	return nil
}
