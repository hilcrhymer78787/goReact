package task

import (
	request "sample/http/request/task/create"
	"sample/pkg/db"

	"github.com/labstack/echo"
)

// タスク登録
func Create(c echo.Context) (err error) {
	db := db.Connect()
	defer db.Close()

	task := new(request.Req)
	if err := c.Bind(task); err != nil {
		return err
	}

	if task.Id != 0 {
		db.Query(`UPDATE tasks set name = ? WHERE id = ?`, task.Name, task.Id)
	} else {
		db.Query(`INSERT INTO tasks (name) VALUES (?)`, task.Name)
	}
	return nil
}
