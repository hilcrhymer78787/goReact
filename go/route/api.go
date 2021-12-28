package route

import (
	"sample/http/controller/task"

	"github.com/labstack/echo"
)

func SetAPI(e *echo.Echo) {
	e.GET("/", task.Test)
	e.POST("/task/create", task.Create)
	e.GET("/task/read", task.Read)
	e.PUT("/task/update", task.Update)
	e.DELETE("/task/delete", task.Delete)
}
