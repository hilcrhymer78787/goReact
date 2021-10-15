package task

import (
	"log"
	"net/http"
	"sample/pkg/db"

	"github.com/labstack/echo"
)

type tasks struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// タスク読み込み
func Read(c echo.Context) error {

	db := db.Connect()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM tasks")
	if err != nil {
		log.Fatal(err)
	}

	tasksArray := make([]tasks, 0)
	for rows.Next() {
		var tasks tasks
		err := rows.Scan(&tasks.ID, &tasks.Name)
		if err != nil {
			log.Fatal(err)
		}
		tasksArray = append(tasksArray, tasks)
	}
	return c.JSON(http.StatusOK, tasksArray)
}
