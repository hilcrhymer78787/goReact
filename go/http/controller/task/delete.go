package task

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"

	request "sample/http/request/task/delete"
	"sample/pkg/db"
)

// タスク削除
func Delete(c echo.Context) error {
	db := db.Connect()
	defer db.Close()

	req := new(request.Req)
	if err := c.Bind(req); err != nil {
		return err
	}

	var err error
	_, err = db.Query(`DELETE FROM tasks WHERE task_id = (?)`, req.Id)

	if err != nil {
		log.Println("DB error:", err)
		return c.JSON(http.StatusInternalServerError, false)
	}
	return c.JSON(http.StatusOK, true)
}
