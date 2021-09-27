package Conn

import (
	"database/sql"
	"fmt"
)

const (
	host = "localhost"
	port = 5432
	user = "postgres"
	pass = "abc123"
	dbname ="mydb"
)

func OpenConnection() *sql.DB {
	conn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, pass, dbname)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
