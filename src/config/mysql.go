package config

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

// InitSql ...
func InitSql() {
	connectString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_HOST"),
		os.Getenv("APP_DB_PORT"),
		os.Getenv("APP_DB_NAME"))
	db, err = sql.Open("mysql", connectString)

	if err != nil {
		fmt.Println(err)
		panic("DB Connection Error")
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		panic(" Ping DB Connection Error")
	}

	// Max connection
	db.SetMaxOpenConns(100)                // Max connection
	db.SetMaxIdleConns(25)                 // Max idle connection
	db.SetConnMaxLifetime(3 * time.Second) // Max lifetime
}

// GetSqlDB ...
func GetSqlDB() *sql.DB {
	return db
}
