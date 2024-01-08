package core

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenGormConnection() *gorm.DB {
	dsn := "host=localhost user=picos password=picos dbname=picos port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db
}
