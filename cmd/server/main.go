package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/tonadr1022/speed-cube-time/internal/config"
	"github.com/tonadr1022/speed-cube-time/internal/db"
	"github.com/tonadr1022/speed-cube-time/internal/user"
)

//	func testHomeHandler(w http.ResponseWriter, r *http.Request) {
//		w.WriteHeader(http.StatusOK)
//		fmt.Fprintf(w, "Hello world\n")
//	}
var flagConfig = flag.String("config", "./config/local.yaml", "path to the config file")

func main() {
	flag.Parse()

	config, err := config.Load(*flagConfig)
	if err != nil {
		log.Fatalf("failed to load application config: %s", err)
	}

	// var err error
	dbConn, err := db.InitDB(config.DSN)
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
	// InitializeHandlers(router, dbConn)
	user.RegisterHandlers(router, user.NewService(user.NewRepository(dbConn)))

	loggedHandler := handlers.LoggingHandler(os.Stdout, router)
	fmt.Printf("Starting server at port %v\n", httpPort)
	log.Fatal(http.ListenAndServe(":"+httpPort, loggedHandler))
}

// func InitializeHandlers(r *mux.Router, dbConn *sql.DB) {
// 	userService := &services.UserService{DB: dbConn}
// 	apiHandler := &apihandlers.Handler{UserService: userService}
// 	r.HandleFunc("/", apihandlers.HomeHandler)
// 	r.HandleFunc("/user/{username}", apiHandler.GetUserByUsername).Methods("GET")
// }
