package settings

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

	r.HandleFunc("/settings/{id}", auth.WithJWTAuth(util.MakeHttpHandler(res.get))).Methods(http.MethodGet)
	r.HandleFunc("/settings/user/{userId}", auth.WithJWTAuth(util.MakeHttpHandler(res.getByUserId))).Methods(http.MethodGet)
	r.HandleFunc("/settings", auth.WithJWTAuth(util.MakeHttpHandler(res.query))).Methods(http.MethodGet)
	r.HandleFunc("/settings", auth.WithJWTAuth(util.MakeHttpHandler(res.create))).Methods(http.MethodPost)
	// r.HandleFunc("/settings", auth.WithJWTAuth(util.MakeHttpHandler(res.query))).Methods(http.MethodGet)
	r.HandleFunc("/settings/{id}", auth.WithJWTAuth(util.MakeHttpHandler(res.update))).Methods(http.MethodPatch)
	r.HandleFunc("/settings/{id}", auth.WithJWTAuth(util.MakeHttpHandler(res.delete))).Methods(http.MethodDelete)
}

// func (res resource) query(w http.ResponseWriter, r *http.Request) error {
// 	users, err := res.service.Query(r.Context())
// 	if err != nil {
// 		return err
// 	}
// 	return util.WriteJson(w, http.StatusOK, users)
// }

func (res resource) delete(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]
	err := res.service.Delete(r.Context(), id)
	if err != nil {
		if err == ErrSettingsNotFound {
			return util.WriteApiError(w, http.StatusNotFound, ErrSettingsNotFound.Error())
		}
		return util.WriteApiError(w, http.StatusInternalServerError, err.Error())
	}
	return util.WriteJson(w, http.StatusNoContent, nil)
}

func (res resource) create(w http.ResponseWriter, r *http.Request) error {
	var payload entity.CreateSettingsPayload
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

func (res resource) update(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]
	var payload entity.UpdateSettingsPayload
	valid := util.ParseValidateWriteIfFailPayload(w, r, &payload)
	if !valid {
		return nil
	}
	response, err := res.service.Update(r.Context(), id, &payload)
	if err != nil {
		return err
	}
	return util.WriteJson(w, http.StatusCreated, response)
}

func (res resource) getByUserId(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["userId"]
	session, err := res.service.GetByUserId(r.Context(), id)
	if err != nil {
		return err
	}
	return util.WriteJson(w, http.StatusOK, session)
}

func (res resource) get(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]
	session, err := res.service.Get(r.Context(), id)
	if err != nil {
		return err
	}
	return util.WriteJson(w, http.StatusOK, session)
}

func (res resource) query(w http.ResponseWriter, r *http.Request) error {
	users, err := res.service.Query(r.Context())
	if err != nil {
		return err
	}
	return util.WriteJson(w, http.StatusOK, users)
}
