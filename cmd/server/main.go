package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/tonadr1022/speed-cube-time/internal/apihandlers"
	"github.com/tonadr1022/speed-cube-time/internal/db"
	"github.com/tonadr1022/speed-cube-time/internal/services"
)

// func testHomeHandler(w http.ResponseWriter, r *http.Request) {
// 	w.WriteHeader(http.StatusOK)
// 	fmt.Fprintf(w, "Hello world\n")
// }

func main() {
	// var err error
	dbConn, err := db.InitDB()
	if err != nil {
		log.Fatalf("failed to initialize the store: %s", err)
	}
	defer dbConn.Close()

	router := mux.NewRouter()

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}
	// router.HandleFunc("/", testHomeHandler)
	InitializeHandlers(router, dbConn)

	loggedHandler := handlers.LoggingHandler(os.Stdout, router)
	fmt.Printf("Starting server at port %v\n", httpPort)
	log.Fatal(http.ListenAndServe(":"+httpPort, loggedHandler))
}

func InitializeHandlers(r *mux.Router, dbConn *sql.DB) {
	userService := &services.UserService{DB: dbConn}
	apiHandler := &apihandlers.Handler{UserService: userService}
	r.HandleFunc("/", apihandlers.HomeHandler)
	r.HandleFunc("/user/{username}", apiHandler.GetUserByUsername).Methods("GET")
}
