package lib

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Gorm *gorm.DB

func InitGormDB() *gorm.DB {
	dsn := "root:123456@tcp(localhost:3306)/casbin_learning?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	mysqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	mysqlDB.SetMaxIdleConns(5)
	mysqlDB.SetMaxOpenConns(5)
	Gorm = db
	return db
}
