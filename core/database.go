package core

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenGormConnection(host, user, password, db string, port int) (*gorm.DB, error) {
	dsn := "host=localhost user=picos password=picos dbname=picos port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func OpenSQLXConnection(host, user, password, db string, port int) (*sqlx.DB, error) {
	return sqlx.Connect("postgres", "host=localhost user=picos password=picos dbname=picos port=5432 sslmode=disable")
}
