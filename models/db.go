package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

var (
	db *gorm.DB
)

func init() {
	var err error
	db, err = gorm.Open("sqlite3", "./systemmonitor.db")
	if err != nil {
		panic(err)
		panic("failed to connect database")
	}
	db.AutoMigrate(&Product{})
	db.AutoMigrate(&OS{})
	db.AutoMigrate(&SYS{})
	db.AutoMigrate(&NETIF{})
	db.AutoMigrate(&NETWORK{})
	db.AutoMigrate(&PARTITION{})
	db.AutoMigrate(&THERMAL{})
}
