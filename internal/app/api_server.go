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

func NewAPIServer(router *mux.Router, db *sql.DB, config *ApiServerConfig) *ApiServer {
	s := &ApiServer{router, db}

	s.router.NotFoundHandler = http.HandlerFunc(notFoundHandler)
	s.router.MethodNotAllowedHandler = http.HandlerFunc(methodNotAllowdHandler)
	settingsRepo := settings.NewRepository(s.db)
	sessionsRepo := session.NewRepository(s.db)
	auth.RegisterHandlers(s.router, auth.NewService(config.JWTSigningKey, config.JWTTokenExpirationMinutes, auth.NewRepository(s.db), settingsRepo, sessionsRepo), auth.WithJWTAuth)
	session.RegisterHandlers(s.router, session.NewService(session.NewRepository(s.db)), auth.WithJWTAuth)
	settings.RegisterHandlers(s.router, settings.NewService(settings.NewRepository(s.db)), auth.WithJWTAuth)
	solve.RegisterHandlers(s.router, solve.NewService(solve.NewRepository(s.db)), auth.WithJWTAuth)
	s.router.Use(mux.CORSMethodMiddleware(s.router))

	return s
}

func (s *ApiServer) Run(httpPort string) {
	loggedHandler := handlers.LoggingHandler(os.Stdout, s.router)
	fmt.Printf("Starting server at port %v\n", httpPort)
	log.Fatal(http.ListenAndServe(":"+httpPort, loggedHandler))
}

// writes error in json rather than default
func methodNotAllowdHandler(w http.ResponseWriter, r *http.Request) {
	util.WriteApiError(w, http.StatusMethodNotAllowed, "method not allowed")
}

// writes error in json rather than default
func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	util.WriteJson(w, http.StatusNotFound, util.ApiError{Error: "not found"})
}
