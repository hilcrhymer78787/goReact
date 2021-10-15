package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func Connect() *sql.DB {
	// envファイル読み込み
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	// データベース接続処理
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/"+os.Getenv("DB_NAME"))
	if err != nil {
		panic(err.Error())
	}
	return db
}
