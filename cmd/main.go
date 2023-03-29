package main

import (
	"github.com/joho/godotenv"
	"github.com/muhangga/config"
	"github.com/muhangga/internal/app"
	"github.com/rs/zerolog/log"
)

func main() {

	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal().Msgf(err.Error())
	}

	db, err := config.OpenConn()
	if err != nil {
		panic("Failed to connect database")
	}
	defer config.CloseConn(db)

	config := config.NewConfig()
	server := app.InitServer(config)

	server.RunServer()
}
