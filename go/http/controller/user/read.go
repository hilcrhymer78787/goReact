package user

import (
	"log"
	"net/http"
	"sample/pkg/db"

	"github.com/labstack/echo"
)

func BearerAuthentication(c echo.Context) error {

	db := db.Connect()
	defer db.Close()

	// res := new(response.Res)

	rows, err := db.Query("SELECT * FROM users WHERE id = 1 LIMIT 1;")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		// user := response.Res{}
		// err := rows.ToStruct(&user.Id, &user.Name, &user.Email, &user.User_img, &user.Token)
		// if err != nil {
		// 	log.Fatal(err)
		// }
	}

	return c.JSON(http.StatusOK, rows)

}
