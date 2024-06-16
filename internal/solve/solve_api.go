package solve

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tonadr1022/speed-cube-time/internal/apperrors"
	"github.com/tonadr1022/speed-cube-time/internal/entity"
	"github.com/tonadr1022/speed-cube-time/internal/util"
)

type resource struct {
	service Service
}

func RegisterHandlers(r *mux.Router, service Service, authHandler util.MiddlewareFunc) {
	res := resource{service}
	// all solves
	r.HandleFunc("/solves", util.MakeHttpHandler(res.query)).Methods(http.MethodGet, http.MethodOptions)
	// number of solves
	r.HandleFunc("/solves/count", util.MakeHttpHandler(res.count)).Methods(http.MethodGet, http.MethodOptions)
	// get solve by id
	r.HandleFunc("/solves/{id}", authHandler(util.MakeHttpHandler(res.get))).Methods(http.MethodGet, http.MethodOptions)
	// create solve
	r.HandleFunc("/solves", authHandler(util.MakeHttpHandler(res.create))).Methods(http.MethodPost, http.MethodOptions)
	// batch delete solves
	r.HandleFunc("/solves", authHandler(util.MakeHttpHandler(res.deleteMany))).Methods(http.MethodDelete, http.MethodOptions)
	// solves of user
	r.HandleFunc("/users/{userId}/solves", authHandler(util.MakeHttpHandler(res.queryForUser))).Methods(http.MethodGet, http.MethodOptions)
	// solves of session
	r.HandleFunc("/sessions/{sessionId}/solves", authHandler(util.MakeHttpHandler(res.queryForSession))).Methods(http.MethodGet, http.MethodOptions)
	// update solve
	r.HandleFunc("/solves/{id}", authHandler(util.MakeHttpHandler(res.update))).Methods(http.MethodPatch, http.MethodOptions)
	// update many solves
	r.HandleFunc("/solves", authHandler(util.MakeHttpHandler(res.updateMany))).Methods(http.MethodPatch, http.MethodOptions)
	// delete solve
	r.HandleFunc("/solves/{id}", authHandler(util.MakeHttpHandler(res.delete))).Methods(http.MethodDelete, http.MethodOptions)
}

func (res resource) delete(w http.ResponseWriter, r *http.Request) error {
	return util.GenericDeleteHandler(res.service, w, r)
}

func (res resource) deleteMany(w http.ResponseWriter, r *http.Request) error {
	return util.GenericDeleteManyHandler(res.service, w, r)
}

func (res resource) updateMany(w http.ResponseWriter, r *http.Request) error {
	var payload []*entity.UpdateManySolvePayload
	err := util.ParseJson(r, &payload)
	if err != nil {
		util.WriteApiError(w, http.StatusBadRequest, apperrors.ErrMalformedRequest.Error())
	}
	err = util.Validate.Var(payload, "required,dive")
	if err != nil {
		return util.WriteApiError(w, http.StatusBadRequest, err.Error())
	}
	err = res.service.UpdateMany(r.Context(), payload)
	if err != nil {
		return err
	}
	return util.WriteJson(w, http.StatusNoContent, nil)
}

func (res resource) create(w http.ResponseWriter, r *http.Request) error {
	var payload entity.CreateSolvePayload
	return util.GenericCreateHandler(res.service, &payload, w, r)
}

func (res resource) update(w http.ResponseWriter, r *http.Request) error {
	var payload entity.UpdateSolvePayload
	return util.GenericUpdateHandler(res.service, &payload, w, r)
}

func (res resource) get(w http.ResponseWriter, r *http.Request) error {
	return util.GenericGetByIdHandler(res.service, w, r)
}

func (res resource) query(w http.ResponseWriter, r *http.Request) error {
	return util.GenericQueryHandler(res.service, w, r)
}

func (res resource) queryForSession(w http.ResponseWriter, r *http.Request) error {
	sessionId := mux.Vars(r)["sessionId"]
	users, err := res.service.QueryForSession(r.Context(), sessionId)
	if err != nil {
		return err
	}
	return util.WriteJson(w, http.StatusOK, users)
}

func (res resource) count(w http.ResponseWriter, r *http.Request) error {
	count, err := res.service.Count(r.Context())
	if err != nil {
		return err
	}
	return util.WriteJson(w, http.StatusOK, util.CountReponse{Count: count})
}

func (res resource) queryForUser(w http.ResponseWriter, r *http.Request) error {
	userId := mux.Vars(r)["userId"]
	users, err := res.service.QueryForUser(r.Context(), userId)
	if err != nil {
		return err
	}
	return util.WriteJson(w, http.StatusOK, users)
}
