package data

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql",
		"root:root@tcp(127.0.0.1:3306)/fake")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	//defer db.Close()

	return db
}
