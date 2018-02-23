package config

import (
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
)

var DB *sql.DB

func Database() {
	dbinfo := fmt.Sprintf("user=%s db_name=%s", "postgres", "gospodar_development")
	DB, _ := sql.Open("postgres", dbinfo)
	DB.Ping()
}
