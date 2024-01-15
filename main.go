package main

import (
	"log"

	"github.com/antunesleo/picos-api/core"
	"github.com/antunesleo/picos-api/spots/application"
	posts "github.com/antunesleo/picos-api/spots/infrastructure/repositories"
	"github.com/antunesleo/picos-api/spots/infrastructure/transactions"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type SpotResponse struct {
	Name string `json:"name"`
}

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
	transactionManager := transactions.SQLXTransactionManager{dbConnection}
	spotRepository := posts.SQLXSpotRepository{}
	spotsUserCases := application.SpotsUseCasesImpl{TransactionManager: &transactionManager, SpotRepository: &spotRepository}

	r := gin.Default()
	r.GET("/spots", func(c *gin.Context) {
		spots, err := spotsUserCases.List()
		if err != nil {
			c.JSON(500, gin.H{
				"error": "can't list spots",
			})
			return
		}
		spotsResponses := []SpotResponse{}
		for _, spot := range spots {
			spotResponse := SpotResponse{Name: spot.Name}
			spotsResponses = append(spotsResponses, spotResponse)
		}
		c.JSON(200, spotsResponses)
	})
	r.Run()
}
