package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func Database() {
	fmt.Println("Database initialize")
	fmt.Println(DatabaseUser, DatabaseHost, DatabaseName, DatabaseSslmode)

	dbinfo := fmt.Sprintf("user=%s host=%s dbname=%s sslmode=disable",
		DatabaseUser, DatabaseHost, DatabaseName, DatabaseSslmode)
	DB, _ = sql.Open("postgres", dbinfo)
	DB.Ping()
}
