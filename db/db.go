package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
)

var db *sql.DB
var err error

func StartDB() {

	if _, err := os.Create("data.db"); err != nil {
		log.Default().Fatal(err)
	}

	db, err = sql.Open("sqlite3", "data.db")

	if err != nil {
		log.Default().Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Default().Fatal(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)

	defer db.Close()

	if _, err = db.Exec("CREATE TABLE IF NOT EXISTS metrics (metric REAL)"); err != nil {
		log.Default().Fatal(err)
	}

	fmt.Println("DB started...")

}

func GetDB() *sql.DB {
	return db
}
