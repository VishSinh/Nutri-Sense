package main

import (
	"user_service/server"

	"github.com/rs/zerolog/log"
)

func main() {
	db, err := server.SetupDatabase()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
	}

	r := server.SetupRouter(db)

	// Run the server
	if err := r.Run(":8080"); err != nil {
		log.Fatal().Err(err).Msg("Failed to start server")
	}
}
