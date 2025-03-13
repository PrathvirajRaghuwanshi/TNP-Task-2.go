package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// Global DB connection
var DB *sql.DB

// ConnectDB initializes the database connection
func ConnectDB() {
	connStr := "host=localhost port=5432 user=your_username password=your_password dbname=certificate_db sslmode=disable"

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(" Database not reachable:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal(" Database connection failed:", err)
	}

	fmt.Println(" Connected to PostgreSQL")
}
