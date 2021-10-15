package route

import (
	"sample/http/controller/task"

	"github.com/labstack/echo"
)

func SetAPI(e *echo.Echo) {
	e.POST("/tasks/create", task.Create)
	e.GET("/tasks/read", task.Read)
	e.DELETE("/tasks/delete", task.Delete)
}
