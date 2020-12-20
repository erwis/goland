package database

import (
	"database/sql"
	"time"
)

func InitDB() *sql.DB {
	cadenaConexion := "root:12345@tcp(localhost:3307)/northwind"
	db, err := sql.Open("mysql", cadenaConexion)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db

}
