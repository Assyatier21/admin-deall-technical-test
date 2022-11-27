package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	var err error

	dsn := Username + ":" + Password + "@tcp(" + Host + ":" + Port + ")/" + Dbname
	DB, err = sql.Open("mysql", dsn)

	if err != nil {
		log.Println("[Database] Connection Failed")
		panic(err)
	}
}

func GetDB() (DB *sql.DB) {
	return DB
}
