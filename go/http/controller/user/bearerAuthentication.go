package user

import (
	"database/sql"
	"log"
	"net/http"
	response "sample/http/response/user/basic_authentication"
	"sample/pkg/db"
	"strings"

	"github.com/labstack/echo"
)

func BearerAuthentication(c echo.Context) error {

	conn := db.Connect()
	defer conn.Close()

	authHeader := c.Request().Header.Get("Authorization")

	token := ""
	if strings.HasPrefix(authHeader, "Bearer ") {
		token = strings.TrimPrefix(authHeader, "Bearer ")
	} else {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "Authorization header missing or invalid",
		})
	}

	// ユーザー情報を取得
	row := conn.QueryRow("SELECT id, email, name, user_img, token FROM users WHERE token = ?", token)

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
