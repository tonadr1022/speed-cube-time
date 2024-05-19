package main

import (
	"log"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tonadr1022/speed-cube-time/internal/app"
	"github.com/tonadr1022/speed-cube-time/internal/db"
)

func main() {
	db, err := db.InitCockroaachDB()
	if err != nil {
		log.Fatalf("failed to initialize the store: %s", err)
	}
	defer db.Close()
	router := mux.NewRouter()
	apiServer := app.NewAPIServer(router, db)

	jwtSigningKey := os.Getenv("JWT_SECRET")
	jwtTokenExpirationMinutes, err := strconv.Atoi(os.Getenv("TOKEN_EXPIRATION_MINUTES"))
	if err != nil {
		log.Fatalf("Could not parse or find TOKEN_EXPIRATION_MINUTES")
	}
	apiServer.Initialize(jwtSigningKey, jwtTokenExpirationMinutes)

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}
	apiServer.Run(httpPort)
}
