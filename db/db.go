package db

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func New(dsn string) *sql.DB {
	db, err := sql.Open("pgx", dsn)

	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(10)

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}
