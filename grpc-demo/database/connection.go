package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetConnection(dsn string) (db *gorm.DB, err error) {
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
