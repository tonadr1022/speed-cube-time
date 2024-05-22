package settings

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

	r.HandleFunc("/settings/{id}", authHandler(util.MakeHttpHandler(res.getById))).Methods(http.MethodGet)
	r.HandleFunc("/users/{userId}/settings", authHandler(util.MakeHttpHandler(res.getByUserId))).Methods(http.MethodGet)
	r.HandleFunc("/settings", authHandler(util.MakeHttpHandler(res.query))).Methods(http.MethodGet)
	r.HandleFunc("/settings", authHandler(util.MakeHttpHandler(res.create))).Methods(http.MethodPost)
	r.HandleFunc("/settings/{id}", authHandler(util.MakeHttpHandler(res.update))).Methods(http.MethodPatch)
	r.HandleFunc("/settings/{id}", authHandler(util.MakeHttpHandler(res.delete))).Methods(http.MethodDelete)
}

func (res resource) delete(w http.ResponseWriter, r *http.Request) error {
	return util.GenericDeleteHandler(res.service, w, r)
}

func (res resource) create(w http.ResponseWriter, r *http.Request) error {
	var payload entity.CreateSettingsPayload
	return util.GenericCreateHandler(res.service, &payload, w, r)
}

func (res resource) update(w http.ResponseWriter, r *http.Request) error {
	var payload entity.UpdateSettingsPayload
	return util.GenericUpdateHandler(res.service, &payload, w, r)
}

func (res resource) getById(w http.ResponseWriter, r *http.Request) error {
	return util.GenericGetByIdHandler(res.service, w, r)
}

func (res resource) query(w http.ResponseWriter, r *http.Request) error {
	return util.GenericQueryHandler(res.service, w, r)
}

func (res resource) getByUserId(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["userId"]
	session, err := res.service.GetByUserId(r.Context(), id)
	if err != nil {
		return err
	}
	return util.WriteJson(w, http.StatusOK, session)
}
