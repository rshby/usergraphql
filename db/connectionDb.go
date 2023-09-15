package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

var Db *sql.DB

func InitDB() {
	dbUser := os.Getenv("DB_USER")
	dbHost := os.Getenv("DB_USER")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	connectionString := fmt.Sprintf("%v:@tcp(%v:%v)/%v?parseTime=True&loc=Local", dbUser, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Panic("cant connect database: " + err.Error())
	}

	log.Println("success connect to database")
	Db = db
}

func CloseDB() error {
	return Db.Close()
}

func Migrate() {
	if err := Db.Ping(); err != nil {
		log.Panic(err)
	}
}
