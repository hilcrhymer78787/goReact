package task

import (
	request "sample/http/request/task/delete"
	"sample/pkg/db"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
)

// タスク削除
func Delete(c echo.Context) error {
	db := db.Connect()
	defer db.Close()

	req := new(request.Req)
	if err := c.Bind(req); err != nil {
		return err
	}

	db.Query(`DELETE FROM tasks WHERE id = (?)`, req.Id)
	return nil
}
