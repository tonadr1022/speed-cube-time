package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/tonadr1022/speed-cube-time/internal/auth"
	"github.com/tonadr1022/speed-cube-time/internal/session"
	"github.com/tonadr1022/speed-cube-time/internal/settings"
	"github.com/tonadr1022/speed-cube-time/internal/solve"
	"github.com/tonadr1022/speed-cube-time/internal/util"
)

type ApiServer struct {
	router *mux.Router
	db     *sql.DB
}

type ApiServerConfig struct {
	JWTSigningKey             string
	JWTTokenExpirationMinutes int
}

// CORS middleware
func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set the necessary headers
		w.Header().Set("Access-Control-Allow-Origin", "*") // Use "*" to allow all origins
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func NewAPIServer(router *mux.Router, db *sql.DB, config *ApiServerConfig) *ApiServer {
	router.Use(cors)
	router.NotFoundHandler = http.HandlerFunc(notFoundHandler)
	router.MethodNotAllowedHandler = http.HandlerFunc(methodNotAllowdHandler)
	settingsRepo := settings.NewRepository(db)
	sessionsRepo := session.NewRepository(db)
	auth.RegisterHandlers(router, auth.NewService(config.JWTSigningKey, config.JWTTokenExpirationMinutes, auth.NewRepository(db), settingsRepo, sessionsRepo), auth.WithJWTAuth)
	session.RegisterHandlers(router, session.NewService(session.NewRepository(db)), auth.WithJWTAuth)
	settings.RegisterHandlers(router, settings.NewService(settings.NewRepository(db)), auth.WithJWTAuth)
	solve.RegisterHandlers(router, solve.NewService(solve.NewRepository(db)), auth.WithJWTAuth)
	s := &ApiServer{router, db}

	return s
}

func (s *ApiServer) Run(httpPort string) {
	loggedHandler := handlers.LoggingHandler(os.Stdout, s.router)

	// credentials := handlers.AllowCredentials()
	// origins := handlers.AllowedOrigins([]string{"http://localhost:8000", "http://127.0.0.1:3000"})
	// methods := handlers.AllowedMethods([]string{"POST", "GET", "PATCH", "PUT", "DELETE", "OPTIONS"})

	fmt.Printf("Starting server at port %v\n", httpPort)
	log.Fatal(http.ListenAndServe(":"+httpPort, loggedHandler))
	// log.Fatal(http.ListenAndServe(":"+httpPort, handlers.CORS(credentials, methods, origins)(loggedHandler)))
}

// writes error in json rather than default
func methodNotAllowdHandler(w http.ResponseWriter, r *http.Request) {
	util.WriteApiError(w, http.StatusMethodNotAllowed, "method not allowed")
}

// writes error in json rather than default
func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	util.WriteJson(w, http.StatusNotFound, util.ApiError{Error: "not found"})
}
