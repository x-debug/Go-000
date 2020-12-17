package data

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"sync"
)

type DBConn struct {
	*sql.DB
	once sync.Once
}

var dbConn DBConn

func NewDB(conf *DBConf) (*DBConn, func()) {
	dbConn.once.Do(func() {
		db, err := sql.Open(conf.driver, conf.uri)
		if err != nil {
			log.Fatal(err)
		}

		err = db.Ping()
		if err != nil {
			log.Fatal(err)
		}

		dbConn.DB = db
	})

	return &dbConn, func(){
		err := dbConn.Close()
		if err != nil {
			log.Println("close dbConn error")
		}
	}
}
