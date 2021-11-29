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
	db, err := sql.Open(os.Getenv("DB_CONNECTION"), os.Getenv("DB_USERNAME")+":"+os.Getenv("DB_PASSWORD")+"@tcp("+os.Getenv("DB_HOST")+":"+os.Getenv("DB_PORT")+")/"+os.Getenv("DB_DATABASE"))
	if err != nil {
		panic(err.Error())
	}
	return db
}
