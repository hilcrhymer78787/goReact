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

	db.Query(`CREATE TABLE IF NOT EXISTS users (
				id SERIAL NOT NULL PRIMARY KEY,
				name VARCHAR(255),
				email VARCHAR(255),
				password VARCHAR(255),
				user_img VARCHAR(255),
				token VARCHAR(255)
			)`)
	db.Query(`TRUNCATE TABLE users`)
	db.Query(`INSERT INTO users (name,email,password,user_img,token) VALUES ('user1','user1@gmail.com','password','user_img','user1@gmail.com')`)
	db.Query(`INSERT INTO users (name,email,password,user_img,token) VALUES ('user2','user2@gmail.com','password','user_img','user2@gmail.com')`)
	db.Query(`INSERT INTO users (name,email,password,user_img,token) VALUES ('user3','user3@gmail.com','password','user_img','user3@gmail.com')`)
	db.Query(`INSERT INTO users (name,email,password,user_img,token) VALUES ('user4','user4@gmail.com','password','user_img','user4@gmail.com')`)
	db.Query(`INSERT INTO users (name,email,password,user_img,token) VALUES ('user5','user5@gmail.com','password','user_img','user5@gmail.com')`)
	db.Query(`INSERT INTO users (name,email,password,user_img,token) VALUES ('user6','user6@gmail.com','password','user_img','user6@gmail.com')`)
	db.Query(`INSERT INTO users (name,email,password,user_img,token) VALUES ('user7','user7@gmail.com','password','user_img','user7@gmail.com')`)
}
