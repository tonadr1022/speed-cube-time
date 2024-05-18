package user

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tonadr1022/speed-cube-time/internal/util"
)

type resource struct {
	service Service
}

func RegisterHandlers(r *mux.Router, service Service) {
	res := resource{service}
	r.HandleFunc("/user/{id}", util.MakeHttpHandleFunc(res.get))
}

func (res resource) get(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]
	userId, err := res.service.Get(id)
	if err != nil {
		return err
	}

	return util.WriteJson(w, http.StatusOK, userId)
}
