package data

import (
	"gopkg.in/ini.v1"
	"log"
	"sync"
)

type DBConf struct {
	uri string
	driver string
	once sync.Once
}

var dbConf DBConf

func NewConf() *DBConf  {
	dbConf.once.Do(func() {
		cfg, err := ini.Load("config/db.ini")
		if err != nil {
			log.Fatalln("load conf error")
		}
		sec := cfg.Section("db")
		dbConf.uri = sec.Key("uri").String()
		dbConf.driver = sec.Key("driver").String()
	})
	return &dbConf
}
