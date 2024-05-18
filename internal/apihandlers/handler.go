package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tonadr1022/speed-cube-time/internal/services"
	"github.com/tonadr1022/speed-cube-time/internal/util"
)

type Handler struct {
	UserService *services.UserService
}

func (h *Handler) HandleUserByID(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		userId := mux.Vars(r)["id"]
		user, err := h.UserService.GetUserById(userId)
		if err != nil {
			return err
		}
		return util.WriteJson(w, http.StatusOK, user)
	}
	// TODO delete
	return fmt.Errorf("method not allowed %s", r.Method)
}

func (h *Handler) HandleCreateUser(w http.ResponseWriter, r *http.Request) error {
	req := services.CreateUserRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return err
	}
	userId, err := h.UserService.CreateUser(req)
	if err != nil {
		return err
	}

	return util.WriteJson(w, http.StatusOK, userId)
}
