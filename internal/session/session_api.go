package session

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tonadr1022/speed-cube-time/internal/auth"
	"github.com/tonadr1022/speed-cube-time/internal/entity"
	"github.com/tonadr1022/speed-cube-time/internal/util"
)

type resource struct {
	service Service
}

func RegisterHandlers(r *mux.Router, service Service) {
	res := resource{service}
	r.HandleFunc("/session", auth.WithJWTAuth(util.MakeHttpHandler(res.create))).Methods(http.MethodPost)
	r.HandleFunc("/session/{id}", auth.WithJWTAuth(util.MakeHttpHandler(res.get))).Methods(http.MethodGet)
	r.HandleFunc("/session/{id}", auth.WithJWTAuth(util.MakeHttpHandler(res.delete))).Methods(http.MethodDelete)
}

func (res resource) delete(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]
	err := res.service.Delete(r.Context(), id)
	if err != nil {
		if err == ErrSessionNotFound {
			return util.WriteApiError(w, http.StatusNotFound, "session not found")
		}
		return util.WriteApiError(w, http.StatusInternalServerError, err.Error())
	}
	return util.WriteJson(w, http.StatusNoContent, nil)
}

func (res resource) create(w http.ResponseWriter, r *http.Request) error {
	var payload entity.CreateSessionPayload

	isValid := util.ParseValidateWriteIfFailPayload(w, r, &payload)
	if !isValid {
		return nil
	}
	response, err := res.service.Create(r.Context(), &payload)
	if err != nil {
		return err
	}
	return util.WriteJson(w, http.StatusCreated, response)
}

func (res resource) get(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]
	session, err := res.service.Get(id)
	if err != nil {
		return err
	}
	return util.WriteJson(w, http.StatusOK, session)
}
