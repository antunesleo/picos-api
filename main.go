package main

import (
	"fmt"
	"log"

	"github.com/antunesleo/picos-api/core"
	"github.com/antunesleo/picos-api/spots/infrastructure/persistence"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No env file, skipping envvar loading")
	}

	config, err := core.LoadConfig()
	if err != nil {
		log.Fatalf("Fail to load env vars, exiting %w", err)
	}

	dbConnection, err := core.OpenSQLXConnection(config.DBHost, config.DBUser, config.DBPassword, config.DBName, config.DBPort, config.DBSSLMode)
	if err != nil {
		log.Fatalf("Fail to open db connection, exiting: %w", err)
	}

	transactionManager := persistence.SQLXTransactionManager{dbConnection}
	tx, err := transactionManager.Begin()
	if err != nil {
		log.Fatalf("Fail to open transaction, exiting: %w", err)
	}
	spotRepository := persistence.SQLXSpotRepository{}
	spots, err := spotRepository.ListAll(tx)
	if err != nil {
		log.Fatalf("Failed to list spots, exiting: %w", err)
		transactionManager.Rollback(tx)
	}
	transactionManager.Commit(tx)

	for _, spot := range spots {
		fmt.Println("spot name", spot.Name)
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
