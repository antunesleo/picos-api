package main

import (
	"log"

	"github.com/antunesleo/picos-api/core"
	"github.com/gin-gonic/gin"
)

func main() {
	config, err := core.LoadConfig()
	if err != nil {
		log.Fatalf("Fail to load env vars, exiting %w", err)
	}

	_, err = core.OpenSQLXConnection(config.DBHost, config.DBUser, config.DBPassword, config.DBName, config.DBPort)
	if err != nil {
		log.Fatalf("Fail to open db connection, exiting: %w", err)
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
