package route

import (
	"sample/http/controller/task"
	"sample/http/controller/test"
	"sample/http/controller/user"

	"github.com/labstack/echo"
)

func SetAPI(e *echo.Echo) {
	e.GET("/api", test.Test)

	e.GET("/api/user/bearer_authentication", user.BearerAuthentication)

	e.POST("/api/task/create", task.Create)
	e.GET("/api/task/read", task.Read)
	e.DELETE("/api/task/delete", task.Delete)
}
