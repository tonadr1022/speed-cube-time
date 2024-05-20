package user

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tonadr1022/speed-cube-time/internal/auth"
	"github.com/tonadr1022/speed-cube-time/internal/util"
)

type resource struct {
	service Service
}

func RegisterHandlers(r *mux.Router, service Service) {
	res := resource{service}
	r.HandleFunc("/user", auth.WithJWTAuth(util.MakeHttpHandler(res.query))).Methods(http.MethodGet)
}

func (res resource) query(w http.ResponseWriter, r *http.Request) error {
	users, err := res.service.Query()
	if err != nil {
		return err
	}
	return util.WriteJson(w, http.StatusOK, users)
}
