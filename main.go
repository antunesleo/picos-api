package main

import (
	"fmt"
	"log"

	"github.com/antunesleo/picos-api/core"
)

func main() {
	config, err := core.LoadConfig()
	if err != nil {
		log.Fatalf("Fail to load env vars, exiting %w", err)
	}

	gorm, err := core.OpenGormConnection(config.DBHost, config.DBUser, config.DBPassword, config.DBName, config.DBPort)
	if err != nil {
		log.Fatalf("Fail to open db connection, exiting: %w", err)
	}
	fmt.Println("Connection opened: ", gorm.Config)
	fmt.Println("Hello World")
}
