package app

import (
	"arieansyah/golang-restful-api/helper"
	"database/sql"
	"os"
	"time"
)

func NewDB() *sql.DB {
	user := os.Getenv("DB_USERNAME")
	dbname := os.Getenv("DB_DATABASE")
	password := "password=" + os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	connection := os.Getenv("DB_CONNECTION")
	port := os.Getenv("DB_PORT")

	db, err := sql.Open(connection, ""+user+":"+password+"@tcp("+host+":"+port+")/"+dbname+"")

	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
