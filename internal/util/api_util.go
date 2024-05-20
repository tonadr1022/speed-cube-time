package util

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func WriteJson(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

var validate = validator.New()

type validationError struct {
	Namespace       string `json:"namespace"` // can differ when a custom TagNameFunc is registered or
	Field           string `json:"field"`     // by passing alt name to ReportError like below
	StructNamespace string `json:"structNamespace"`
	StructField     string `json:"structField"`
	Tag             string `json:"tag"`
	ActualTag       string `json:"actualTag"`
	Kind            string `json:"kind"`
	Type            string `json:"type"`
	Value           string `json:"value"`
	Param           string `json:"param"`
	Message         string `json:"message"`
}

func ParseJson(r *http.Request, v any) error {
	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}
	return json.NewDecoder(r.Body).Decode(v)
}

func WriteApiError(w http.ResponseWriter, status int, message string) error {
	return WriteJson(w, status, ApiError{Error: message})
}

func WriteMalformedRequest(w http.ResponseWriter) error {
	return WriteApiError(w, http.StatusBadRequest, "malformed request")
}

type ApiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string `json:"error"`
}

func MakeHttpHandler(next ApiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		if r.Method == http.MethodOptions {
			return
		}
		if err := next(w, r); err != nil {
			WriteJson(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

// returns true if payload is parsed and validated without error.
// otherwise returns false, writes the error to the response, and returns any unexpected errors
// that typically shouldn't happen
func ParseValidateWriteIfFailPayload(w http.ResponseWriter, r *http.Request, payload any) bool {
	err := ParseJson(r, &payload)
	if err != nil {
		WriteMalformedRequest(w)
		return false
	}
	err = validate.Struct(payload)
	if err != nil {
		WriteApiError(w, http.StatusBadRequest, err.Error())
		return false
	}
	return true
}
