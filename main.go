package main

import (
	"fmt"

	"github.com/antunesleo/picos-api/core"
)

func main() {
	gorm := core.OpenGormConnection()
	fmt.Println("Connection opened: ", gorm.Config)
	fmt.Println("Hello World")
}
