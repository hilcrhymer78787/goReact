package user

import (
	"database/sql"
	"log"
	"net/http"

	request "sample/http/request/user/basic_authentication"
	response "sample/http/response/user/basic_authentication"
	"sample/pkg/db"

	"github.com/labstack/echo"
)

func BasicAuthentication(c echo.Context) error {
	conn := db.Connect()
	defer conn.Close()

	req := new(request.Req)
	if err := c.Bind(req); err != nil {
		log.Println("Bind error:", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	// ユーザー情報を取得
	row := conn.QueryRow("SELECT id, email, name, user_img, token FROM users WHERE email = ? AND password = ?", req.Email, req.Password)

	var res response.Res
	err := row.Scan(&res.Id, &res.Email, &res.Name, &res.User_img, &res.Token)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Email or password is incorrect"})
		}
		log.Println("DB error:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal server error"})
	}

	return c.JSON(http.StatusOK, res)
}
