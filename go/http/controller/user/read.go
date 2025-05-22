package user

import (
	"net/http"

	"github.com/labstack/echo"
)

func BearerAuthentication(c echo.Context) error {

	return c.JSON(http.StatusInternalServerError, false)

	// db := db.Connect()
	// defer db.Close()

	// res := new(response.Res)

	// row := db.QueryRow("SELECT id, name, email, user_img, token FROM users WHERE id = 1 LIMIT 1;")

	// err := row.Scan(&res.Id, &res.Name, &res.Email, &res.User_img, &res.Token)

	// if err != nil {
	// 	log.Fatal(err)
	// 	return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
	// }

	// return c.JSON(http.StatusOK, res)

}
