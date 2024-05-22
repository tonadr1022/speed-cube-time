package session

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tonadr1022/speed-cube-time/internal/entity"
	"github.com/tonadr1022/speed-cube-time/internal/util"
)

type resource struct {
	service Service
}

func RegisterHandlers(r *mux.Router, service Service, authHandler util.MiddlewareFunc) {
	res := resource{service}
	// get by id
	r.HandleFunc("/sessions/{id}", authHandler(util.MakeHttpHandler(res.getById))).Methods(http.MethodGet)
	// create
	r.HandleFunc("/sessions", authHandler(util.MakeHttpHandler(res.create))).Methods(http.MethodPost)
	// update
	r.HandleFunc("/sessions/{id}", authHandler(util.MakeHttpHandler(res.update))).Methods(http.MethodPatch)
	// delete
	r.HandleFunc("/sessions/{id}", authHandler(util.MakeHttpHandler(res.delete))).Methods(http.MethodDelete)
	// query by user
	r.HandleFunc("/users/{userId}/sessions", authHandler(util.MakeHttpHandler(res.queryByUser))).Methods(http.MethodGet)
	// query all
	r.HandleFunc("/sessions", authHandler(util.MakeHttpHandler(res.query))).Methods(http.MethodGet)
}

func (res resource) delete(w http.ResponseWriter, r *http.Request) error {
	return util.GenericDeleteHandler(res.service, w, r)
}

func (res resource) create(w http.ResponseWriter, r *http.Request) error {
	var payload entity.CreateSessionPayload
	return util.GenericCreateHandler(res.service, &payload, w, r)
}

func (res resource) update(w http.ResponseWriter, r *http.Request) error {
	var payload entity.UpdateSessionPayload
	return util.GenericUpdateHandler(res.service, &payload, w, r)
}

func (res resource) getById(w http.ResponseWriter, r *http.Request) error {
	return util.GenericGetByIdHandler(res.service, w, r)
}

func (res resource) query(w http.ResponseWriter, r *http.Request) error {
	return util.GenericQueryHandler(res.service, w, r)
}

func (res resource) queryByUser(w http.ResponseWriter, r *http.Request) error {
	userId := mux.Vars(r)["userId"]
	users, err := res.service.QueryByUser(r.Context(), userId)
	if err != nil {
		return err
	}
	return util.WriteJson(w, http.StatusOK, users)
}
