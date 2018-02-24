package config

import (
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
)

var DB *sql.DB

func Database() {
	dbinfo := fmt.Sprintf("user=%s host=%s dbname=%s sslmode=disable", "postgres", "gospodar_db", "gospodar_development")
	DB, _ = sql.Open("postgres", dbinfo)
	DB.Ping()
}
