package task

import (
	"log"
	"net/http"
	response "sample/http/response/task/read"
	"sample/pkg/db"

	"github.com/labstack/echo"
)

// タスク読み込み
func Read(c echo.Context) error {

	db := db.Connect()
	defer db.Close()

	date := c.QueryParam("date")
	userId := c.QueryParam("userId")

	res := new(response.Res)
	res.Date = date

	rows, err := db.Query("SELECT * FROM tasks WHERE task_user_id = ?", userId)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		task := response.Task{}
		err := rows.Scan(&task.Id, &task.UserId, &task.SortKey, &task.Name)
		if err != nil {
			log.Fatal(err)
		}
		res.Tasks = append(res.Tasks, task)
	}
	return c.JSON(http.StatusOK, res)
}
