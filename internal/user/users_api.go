package user

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
	r.HandleFunc("/user/{id}", util.MakeHttpHandler(res.get))
	r.HandleFunc("/user", util.MakeHttpHandler(res.create)).Methods(http.MethodPost)
	r.HandleFunc("/user", util.MakeHttpHandler(res.query)).Methods(http.MethodGet)
}

func (res resource) get(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]
	userId, err := res.service.Get(id)
	if err != nil {
		return err
	}

	return util.WriteJson(w, http.StatusOK, userId)
}

func (res resource) query(w http.ResponseWriter, r *http.Request) error {
	users, err := res.service.Query()
	if err != nil {
		return err
	}
	return util.WriteJson(w, http.StatusOK, users)
}

func (res resource) create(w http.ResponseWriter, r *http.Request) error {
	req := &CreateUserRequest{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return err
	}
	response, err := res.service.Create(req)
	if err != nil {
		return err
	}
	return util.WriteJson(w, http.StatusOK, response)
}
