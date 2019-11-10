package database

import (
	"errors"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Conn struct {
	*gorm.DB
}

var Database Conn

func (c Conn) IsInit() bool {
	return c.DB != nil
}

func InitDb() (err error) {
	if Database.DB != nil {
		return
	}

	// TODO: ensure proper credentials mechanism (use env vars?)
	// TODO: set up UTC locale?
	db, err := gorm.Open("mysql", "root:pass@/hello?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return
	}
	Database.DB = db
	return
}

func GetDb() (c Conn, err error) {
	if Database.DB == nil {
		err = errors.New("database is not initialized")
		return
	}
	c = Database
	return
}

func CloseDb() {
	if Database.DB != nil {
		Database.Close()
	}
}
