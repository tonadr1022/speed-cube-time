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

func RegisterHandlers(r *mux.Router, service Service, authHandler util.MiddlewareFunc) {
	res := resource{service}
	// login the user
	r.HandleFunc("/login", util.MakeHttpHandler(res.login)).Methods(http.MethodPost)
	// register the user
	r.HandleFunc("/register", util.MakeHttpHandler(res.register)).Methods(http.MethodPost)
	// delete user
	r.HandleFunc("/unregister", authHandler(util.MakeHttpHandler(res.delete))).Methods(http.MethodDelete)
	// get current user from auth context
	r.HandleFunc("/users/me", authHandler(util.MakeHttpHandler(res.getByCurrentUser))).Methods(http.MethodGet)
	r.HandleFunc("/users/user-settings", authHandler(util.MakeHttpHandler(res.getUserSettings))).Methods(http.MethodGet)
	// get all users
	r.HandleFunc("/users", authHandler(util.MakeHttpHandler(res.query))).Methods(http.MethodGet)
	// get user by id
	r.HandleFunc("/users/{id}", authHandler(util.MakeHttpHandler(res.getById))).Methods(http.MethodGet)
	// update user by id
	r.HandleFunc("/users/{id}", authHandler(util.MakeHttpHandler(res.update))).Methods(http.MethodPatch)
}

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (res resource) getUserSettings(w http.ResponseWriter, r *http.Request) error {
	userCtx := CurrentUser(r.Context())
	actualUser, err := res.service.GetUserAndSettings(r.Context(), userCtx.GetID())
	if err != nil {
		return err
	}
	return util.WriteJson(w, http.StatusOK, actualUser)
}

func (res resource) getByCurrentUser(w http.ResponseWriter, r *http.Request) error {
	userCtx := CurrentUser(r.Context())
	actualUser, err := res.service.Get(r.Context(), userCtx.GetID())
	if err != nil {
		return err
	}
	return util.WriteJson(w, http.StatusOK, actualUser)
}

func (res resource) getById(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]
	response, err := res.service.Get(r.Context(), id)
	if err != nil {
		return err
	}
	return util.WriteJson(w, http.StatusOK, response)
}

func (res resource) delete(w http.ResponseWriter, r *http.Request) error {
	err := res.service.Delete(r.Context())
	if err != nil {
		return err
	}
	return util.WriteJson(w, http.StatusNoContent, nil)
}

func (res resource) query(w http.ResponseWriter, r *http.Request) error {
	users, err := res.service.Query(r.Context())
	if err != nil {
		return err
	}
	return util.WriteJson(w, http.StatusOK, users)
}

func (res resource) update(w http.ResponseWriter, r *http.Request) error {
	var user entity.UpdateUserPayload
	// return util.GenericUpdateHandler(res.service, &user, w, r)
	isValid := util.ParseValidateWriteIfFailPayload(w, r, &user)
	if !isValid {
		return nil
	}
	err := res.service.Update(r.Context(), &user)
	if err != nil {
		return err
	}
	return util.WriteJson(w, http.StatusNoContent, nil)
}

func (res resource) register(w http.ResponseWriter, r *http.Request) error {
	var user entity.RegisterUserPayload
	isValid := util.ParseValidateWriteIfFailPayload(w, r, &user)
	if !isValid {
		return nil // already handled
	}
	response, err := res.service.Register(r.Context(), &user)
	if err != nil {
		return err
	}
	return util.WriteJson(w, http.StatusCreated, response)
}

func (res resource) login(w http.ResponseWriter, r *http.Request) error {
	var user entity.LoginUserPayload
	isValid := util.ParseValidateWriteIfFailPayload(w, r, &user)
	if !isValid {
		return nil // already handled
	}
	loginResponse, err := res.service.Login(r.Context(), &user)
	if err != nil {
		return err
	}
	return util.WriteJson(w, http.StatusOK, loginResponse)
}
