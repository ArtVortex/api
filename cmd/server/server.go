package main

import (
	"artvortex-api/https"
	"artvortex-api/internal/config"
	pg "artvortex-api/internal/postgres"
	"fmt"
	"log"
)

func main() {
	//local config
	config.InitializeConfig()
	nftConfig := config.NFTServiceConfig

	//set up postgres connection
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		nftConfig.DBHost, nftConfig.DBPort, nftConfig.DBUser, nftConfig.DBName, nftConfig.DBPassword,
	)
	_, db_error := pg.NewPGAdapter(dsn)

	if db_error != nil {
		log.Fatal(db_error)
	}

	// start server
	err := https.StartServer(nftConfig.Port)
	if err == nil {
		log.Fatalln("Failed to start server", err)
	}
}
