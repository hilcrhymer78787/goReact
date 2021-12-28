package task

import (
	"log"
	"net/http"
	"sample/pkg/db"

	"github.com/labstack/echo"
)

type task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// タスク読み込み
func Read(c echo.Context) error {

	db := db.Connect()
	defer db.Close()

	// res := new(response.Tasks)

	rows, err := db.Query("SELECT * FROM tasks")
	if err != nil {
		log.Fatal(err)
	}

	tasks := make([]task, 0)
	for rows.Next() {
		var task task
		err := rows.Scan(&task.ID, &task.Name)
		if err != nil {
			log.Fatal(err)
		}
		tasks = append(tasks, task)
	}
	return c.JSON(http.StatusOK, tasks)
}
