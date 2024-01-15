package config

import (
	"database/sql"
	"log"

	// "time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnnectDB() {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/belajar_golang?parseTime=true")
	if err != nil {
		panic(err)
	}

	log.Println("Database Connected")
	DB = db

	// db.SetMaxIdleConns(10)
	// db.SetMaxOpenConns(100)
	// db.SetConnMaxIdleTime(5 * time.Minute)
	// db.SetConnMaxLifetime(60 * time.Minute)

	// // return db
	// defer db.Close()
}
