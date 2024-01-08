package core

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenGormConnection(host, user, password, db string, port int) (*gorm.DB, error) {
	dsn := "host=localhost user=picos password=picos dbname=picos port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
