package main

import (
	"sample/pkg/db"
	"sample/route"

	"github.com/labstack/echo/middleware"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
)

func main() {

	initDb()

	e := echo.New()

	// ミドルウェア
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	route.SetAPI(e)

	e.Start(":1323")
}

// データベース初期化
func initDb() {
	db := db.Connect()
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
}
