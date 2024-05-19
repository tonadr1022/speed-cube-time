package auth

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tonadr1022/speed-cube-time/internal/util"
)

type resource struct {
	service Service
}

func RegisterHandlers(r *mux.Router, service Service) {
	res := resource{service}
	r.HandleFunc("/login", util.MakeHttpHandler(res.login)).Methods(http.MethodPost)
	r.HandleFunc("/register", util.MakeHttpHandler(res.register)).Methods(http.MethodPost)
	r.HandleFunc("/user", util.MakeHttpHandler(res.delete)).Methods(http.MethodDelete)
}

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (res resource) delete(w http.ResponseWriter, r *http.Request) error {
	err := res.service.DeleteUser(r.Context())
	if err != nil {
		return util.WriteJson(w, http.StatusInternalServerError, util.ApiError{Error: "failed to delete user"})
	}
	return util.WriteJson(w, http.StatusNoContent, nil)
}

func (res resource) register(w http.ResponseWriter, r *http.Request) error {
	req := &RegisterUserRequest{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return util.MalformedRequest(w)
	}
	response, err := res.service.Register(req)
	if err != nil {
		return err
	}
	return util.WriteJson(w, http.StatusOK, response)
}

func (res resource) login(w http.ResponseWriter, r *http.Request) error {
	req := &loginRequest{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return util.MalformedRequest(w)
	}

	loginResponse, err := res.service.Login(req.Username, req.Password)
	if err != nil {
		return err
	}
	return util.WriteJson(w, http.StatusOK, loginResponse)
}
