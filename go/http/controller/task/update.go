package task

import (
	"sample/pkg/db"

	"github.com/labstack/echo"
)

// タスク登録
func Update(c echo.Context) (err error) {
	db := db.Connect()
	defer db.Close()

	type Task struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}
	task := new(Task)
	if err := c.Bind(task); err != nil {
		return err
	}
	db.Query(`UPDATE tasks set name = ? WHERE id = ?`, task.Name, task.Id)
	return nil
}
