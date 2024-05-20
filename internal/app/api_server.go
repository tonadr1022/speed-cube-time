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
	"github.com/tonadr1022/speed-cube-time/internal/user"
	"github.com/tonadr1022/speed-cube-time/internal/util"
)

type ApiServer struct {
	router *mux.Router
	db     *sql.DB
}

func NewAPIServer(router *mux.Router, db *sql.DB) *ApiServer {
	return &ApiServer{router, db}
}

type ApiServerConfig struct {
	JWTSigningKey             string
	JWTTokenExpirationMinutes int
}

func (s *ApiServer) Initialize(config *ApiServerConfig) {
	s.router.NotFoundHandler = http.HandlerFunc(notFoundHandler)
	s.router.MethodNotAllowedHandler = http.HandlerFunc(methodNotAllowdHandler)
	user.RegisterHandlers(s.router, user.NewUserService(user.NewRepository(s.db)))
	auth.RegisterHandlers(s.router, auth.NewService(config.JWTSigningKey, config.JWTTokenExpirationMinutes, auth.NewRepository(s.db)))
	s.router.Use(mux.CORSMethodMiddleware(s.router))
}

func (s *ApiServer) Run(httpPort string) {
	loggedHandler := handlers.LoggingHandler(os.Stdout, s.router)
	fmt.Printf("Starting server at port %v\n", httpPort)
	log.Fatal(http.ListenAndServe(":"+httpPort, loggedHandler))
}

func methodNotAllowdHandler(w http.ResponseWriter, r *http.Request) {
	util.WriteApiError(w, http.StatusMethodNotAllowed, "method not allowed")
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	util.WriteJson(w, http.StatusNotFound, util.ApiError{Error: "not found"})
}
