package route

import (
	"sample/http/controller/test"
	"sample/http/controller/task"
	"sample/http/controller/user"

	"github.com/labstack/echo"
)

func SetAPI(e *echo.Echo) {
	e.GET("/", test.Test)

	e.GET("/user/bearer_authentication", user.BearerAuthentication)

	e.POST("/task/create", task.Create)
	e.GET("/task/read", task.Read)
	e.DELETE("/task/delete", task.Delete)
}
