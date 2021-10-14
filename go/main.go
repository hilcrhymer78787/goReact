package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/labstack/echo/middleware"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
)

type Opening struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	// データベース接続処理
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/test")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.Query(`CREATE TABLE IF NOT EXISTS tasks (
		id SERIAL NOT NULL PRIMARY KEY,
		name VARCHAR(255)
	  )`)
	db.Query(`TRUNCATE TABLE tasks`)
	db.Query(`INSERT INTO tasks (name) VALUES ('Task1')`)
	db.Query(`INSERT INTO tasks (name) VALUES ('Task2')`)
	db.Query(`INSERT INTO tasks (name) VALUES ('Task3')`)
	db.Query(`INSERT INTO tasks (name) VALUES ('Task4')`)
	db.Query(`INSERT INTO tasks (name) VALUES ('Task5')`)

	e := echo.New()

	// 全てのリクエストで差し込みたいミドルウェア（ログとか）はここ
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	//   e.GET("/opening/read", read)
	e.GET("/opening/allread", allread)
	// e.GET("/test", test)
	e.Start("localhost:1323")
}

// func read(c echo.Context) error {

// 	// データベース接続処理
// 	db, err := sql.Open("mysql", "root@/test")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()

// 	id := c.QueryParam("id")
// 	var name string
// 	err = db.QueryRow("SELECT name FROM opening WHERE opening_id = ?", id).Scan(&name)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	data := map[string]string{
// 		"id":   id,
// 		"name": name,
// 	}

// 	return c.JSON(http.StatusOK, data)
// }

func allread(c echo.Context) error {

	// データベース接続処理
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/test")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM opening")
	if err != nil {
		log.Fatal(err)
	}

	openingArray := make([]Opening, 0)
	for rows.Next() {
		var opening Opening
		err := rows.Scan(&opening.ID, &opening.Name)
		if err != nil {
			log.Fatal(err)
		}
		openingArray = append(openingArray, opening)
	}
	return c.JSON(http.StatusOK, openingArray)
}

// func test(c echo.Context) error {

//  // データベース接続処理
//  db, err := sql.Open("mysql", "root@/test")
//  if err != nil {
//    log.Fatal(err)
//  }
//  defer db.Close()

//  db.Query("SELECT * FROM opening")
//  return c.JSON(http.StatusOK, "rows")

// }
