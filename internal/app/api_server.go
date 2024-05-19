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

func (s *ApiServer) Initialize(jwtSigningKey string, jwtTokenExpirationMinutes int) {
	h := http.HandlerFunc(notFound)
	s.router.NotFoundHandler = h
	user.RegisterHandlers(s.router, user.NewUserService(user.NewRepository(s.db)))
	auth.RegisterHandlers(s.router, auth.NewService(jwtSigningKey, jwtTokenExpirationMinutes, auth.NewRepository(s.db)))
	s.router.Use(mux.CORSMethodMiddleware(s.router))
}

func (s *ApiServer) Run(httpPort string) {
	loggedHandler := handlers.LoggingHandler(os.Stdout, s.router)
	fmt.Printf("Starting server at port %v\n", httpPort)
	log.Fatal(http.ListenAndServe(":"+httpPort, loggedHandler))
}

func notFound(w http.ResponseWriter, r *http.Request) {
	util.WriteJson(w, http.StatusNotFound, util.ApiError{Error: "not found"})
}
