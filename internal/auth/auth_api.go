package auth

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tonadr1022/speed-cube-time/internal/entity"
	"github.com/tonadr1022/speed-cube-time/internal/util"
)

type resource struct {
	service Service
}

func RegisterHandlers(r *mux.Router, service Service) {
	res := resource{service}
	r.HandleFunc("/login", util.MakeHttpHandler(res.login)).Methods(http.MethodPost)
	r.HandleFunc("/register", util.MakeHttpHandler(res.register)).Methods(http.MethodPost)
	r.HandleFunc("/user", WithJWTAuth(util.MakeHttpHandler(res.delete))).Methods(http.MethodDelete)
}

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (res resource) delete(w http.ResponseWriter, r *http.Request) error {
	err := res.service.DeleteUser(r.Context())
	if err != nil {
		if err == ErrUserNotFound {
			return util.WriteApiError(w, http.StatusNotFound, ErrUserNotFound.Error())
		}
		return util.WriteMalformedRequest(w)
	}
	return util.WriteJson(w, http.StatusNoContent, nil)
}

func (res resource) register(w http.ResponseWriter, r *http.Request) error {
	var user entity.RegisterUserPayload
	isValid := util.ParseValidateWriteIfFailPayload(w, r, &user)
	if !isValid {
		return nil
	}

	response, err := res.service.Register(&user)
	if err != nil {
		return err
	}
	return util.WriteJson(w, http.StatusOK, response)
}

func (res resource) login(w http.ResponseWriter, r *http.Request) error {
	var user entity.LoginUserPayload
	isValid := util.ParseValidateWriteIfFailPayload(w, r, &user)
	if !isValid {
		return nil
	}

	loginResponse, err := res.service.Login(&user)
	if err != nil {
		return err
	}
	return util.WriteJson(w, http.StatusOK, loginResponse)
}
