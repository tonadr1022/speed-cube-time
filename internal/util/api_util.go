package util

import (
	"encoding/json"
	"net/http"
)

func WriteJson(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

type (
	apiFunc  func(http.ResponseWriter, *http.Request) error
	ApiError struct {
		Error string `json:"error"`
	}
)

func MakeHttpHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJson(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}
