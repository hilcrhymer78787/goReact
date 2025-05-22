package task

import (
	"log"
	"net/http"

	"github.com/labstack/echo"

	request "sample/http/request/task/create"
	"sample/pkg/db"
)

// タスク登録
func Create(c echo.Context) error {
	db := db.Connect()
	defer db.Close()

	task := new(request.Req)
	if err := c.Bind(task); err != nil {
		log.Println("Bind error:", err)
		return c.JSON(http.StatusBadRequest, false)
	}

	var err error
	if task.Id != 0 {
		_, err = db.Query(`UPDATE tasks SET task_name = ? WHERE task_id = ?`, task.Name, task.Id)
	} else {
		_, err = db.Query(`INSERT INTO tasks (task_name, task_user_id) VALUES (?, ?)`, task.Name, 1)
	}

	if err != nil {
		log.Println("DB error:", err)
		return c.JSON(http.StatusInternalServerError, false)
	}
	return c.JSON(http.StatusOK, true)
}
