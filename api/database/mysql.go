package database

import (
	"fmt"
	"shunp/api/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	fmt.Println("[DATA INIT]")
	conf := config.Get()
	db, err := gorm.Open("mysql", conf.Dbhost)

	if err != nil {
		return nil, err
	}

	db.DB().SetMaxIdleConns(100)
	DB = db

	return db, nil
}
