package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/headend/share-module/configuration"
	"time"
)

type Database struct {
	Db *gorm.DB
	Err error
}

func (e *Database) Connect(config *configuration.Conf) *Database {
	config.LoadConf()
	args := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", config.DB.User, config.DB.Password, config.DB.Host, config.DB.Port, config.DB.DBName)
	//log.Printf("Connection info: %v", args)
	e.Db, e.Err = gorm.Open("mysql", args)
	e.Db.DB().SetConnMaxLifetime(time.Minute * 5)
	e.Db.DB().SetMaxIdleConns(0)
	e.Db.DB().SetMaxOpenConns(10000)
	return e
}

