package utils

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/ramdanariadi/grocery-product-service/main/helpers"
	"os"
	"time"
)

func NewDbConnection() (*sql.DB, error) {
	dbUsr := os.Getenv("DB_USR")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	connStr := fmt.Sprintf("postgres://%s:%s@localhost/%s?sslmode=disable", dbUsr, dbPass, dbName)
	db, err := sql.Open("postgres", connStr)
	helpers.PanicIfError(err)

	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db, err
}
