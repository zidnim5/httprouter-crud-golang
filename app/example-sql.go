package app

import (
	"api-test/helper"
	"database/sql"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func NewDB() *sql.DB {
	db, err := sql.Open("pgx", "postgres://dbmasteruser:postgres123!@$@43.231.128.35:5432/db_test")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
