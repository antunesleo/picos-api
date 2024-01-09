package core

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func OpenSQLXConnection(host, user, password, db string, port int, sslmode string) (*sqlx.DB, error) {
	dataSource := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s", host, user, password, db, port, sslmode,
	)
	fmt.Println("dataSource", dataSource)
	return sqlx.Connect("postgres", dataSource)
}
