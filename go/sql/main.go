package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

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
}
