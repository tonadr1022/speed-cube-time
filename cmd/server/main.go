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

	jwtSigningKey := os.Getenv("JWT_SECRET")
	if jwtSigningKey == "" {
		log.Fatalln("Could not find JWT_SECRET")
	}
	jwtTokenExpirationMinutes, err := strconv.Atoi(os.Getenv("TOKEN_EXPIRATION_MINUTES"))
	if err != nil {
		log.Fatalf("Could not parse or find TOKEN_EXPIRATION_MINUTES")
	}
	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	config := &app.ApiServerConfig{
		JWTSigningKey:             jwtSigningKey,
		JWTTokenExpirationMinutes: jwtTokenExpirationMinutes,
	}

	apiServer := app.NewAPIServer(router, db, config)
	apiServer.Run(httpPort)
}
