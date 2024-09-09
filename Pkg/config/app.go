package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {

	d, err := gorm.Open("mysql", "avnadmin:AVNS_RFIGYSWDDEJ84pJeLnC@mysql-a94996d-aggarwalprabal1-dc76.e.aivencloud.com:16598/issue?ssl-mode=REQUIRED")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
