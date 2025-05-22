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
				task_id SERIAL NOT NULL PRIMARY KEY,
				task_user_id INT,
				task_sort_key INT,
				task_name VARCHAR(255)
			)`)
	db.Query(`TRUNCATE TABLE tasks`)
	db.Query(`INSERT INTO tasks (task_name,task_user_id,task_sort_key) VALUES ('Task1',1,1)`)
	db.Query(`INSERT INTO tasks (task_name,task_user_id,task_sort_key) VALUES ('Task2',1,2)`)
	db.Query(`INSERT INTO tasks (task_name,task_user_id,task_sort_key) VALUES ('Task3',1,3)`)
	db.Query(`INSERT INTO tasks (task_name,task_user_id,task_sort_key) VALUES ('Task4',1,4)`)
	db.Query(`INSERT INTO tasks (task_name,task_user_id,task_sort_key) VALUES ('Task2_1',2,1)`)
	db.Query(`INSERT INTO tasks (task_name,task_user_id,task_sort_key) VALUES ('Task2_2',2,2)`)

	db.Query(`CREATE TABLE IF NOT EXISTS users (
				id SERIAL NOT NULL PRIMARY KEY,
				name VARCHAR(255),
				email VARCHAR(255),
				password VARCHAR(255),
				user_img VARCHAR(255),
				token VARCHAR(255)
			)`)
	db.Query(`TRUNCATE TABLE users`)
	db.Query(`INSERT INTO users (name,email,password,user_img,token) VALUES ('user1','user1@gmail.com','password','user_img','tokentokentoken-user1@gmail.com')`)
	db.Query(`INSERT INTO users (name,email,password,user_img,token) VALUES ('user2','user2@gmail.com','password','user_img','tokentokentoken-user2@gmail.com')`)
	db.Query(`INSERT INTO users (name,email,password,user_img,token) VALUES ('user3','user3@gmail.com','password','user_img','tokentokentoken-user3@gmail.com')`)
	db.Query(`INSERT INTO users (name,email,password,user_img,token) VALUES ('user4','user4@gmail.com','password','user_img','tokentokentoken-user4@gmail.com')`)
	db.Query(`INSERT INTO users (name,email,password,user_img,token) VALUES ('user5','user5@gmail.com','password','user_img','tokentokentoken-user5@gmail.com')`)
	db.Query(`INSERT INTO users (name,email,password,user_img,token) VALUES ('user6','user6@gmail.com','password','user_img','tokentokentoken-user6@gmail.com')`)
	db.Query(`INSERT INTO users (name,email,password,user_img,token) VALUES ('user7','user7@gmail.com','password','user_img','tokentokentoken-user7@gmail.com')`)
}
