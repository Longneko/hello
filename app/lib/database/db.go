package database

import (
	"errors"
	"time"

	"github.com/Longneko/lamp/app/lib/config"

	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
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

	appCfg := config.Get()
	dbCfg := mysql.NewConfig()
	dbCfg.Net = "tcp"
	dbCfg.Addr = appCfg.MySql.Addr
	dbCfg.User = appCfg.MySql.User
	dbCfg.Passwd = appCfg.MySql.Password
	dbCfg.DBName = appCfg.MySql.DbName
	dbCfg.Loc = time.UTC
	dbCfg.ParseTime = true
	dbCfg.Params = map[string]string{
		"time_zone": `"+00:00"`, // set "+00:00"
	}

	db, err := gorm.Open("mysql", dbCfg.FormatDSN())
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
