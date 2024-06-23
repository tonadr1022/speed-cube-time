package util

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/tonadr1022/speed-cube-time/internal/apperrors"
)

func WriteJson(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if v == nil {
		return nil
	}
	return json.NewEncoder(w).Encode(v)
}

var Validate = validator.New()

func ParseJson(r *http.Request, v any) error {
	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}
	return json.NewDecoder(r.Body).Decode(v)
}

func WriteApiError(w http.ResponseWriter, status int, message string) error {
	return WriteJson(w, status, ApiError{Error: message})
}

type ApiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string `json:"error"`
}

type MiddlewareFunc func(next http.HandlerFunc) http.HandlerFunc

type CountReponse struct {
	Count int `json:"count"`
}

func MakeHttpHandler(next ApiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // Use "*" to allow all origins
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		// Handle preflight requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		if err := next(w, r); err != nil {
			switch {
			case errors.Is(err, apperrors.ErrNotFound):
			case errors.Is(err, sql.ErrNoRows):
				WriteApiError(w, http.StatusNotFound, apperrors.ErrNotFound.Error())
			case errors.Is(err, apperrors.ErrAlreadyExists):
				WriteApiError(w, http.StatusBadRequest, apperrors.ErrAlreadyExists.Error())
			default:
				WriteApiError(w, http.StatusInternalServerError, err.Error())
			}
		}
	}
}

// defines a generic CRUD service that taks create and update payloads and returns
// a resource for each operation
type CRUDService[P any, R any, T any] interface {
	Get(ctx context.Context, id string) (*R, error)
	Create(ctx context.Context, req *P) (*R, error)
	Update(ctx context.Context, id string, req *T) error
	Delete(ctx context.Context, id string) error
}

type DeleteManyService interface {
	DeleteMany(ctx context.Context, ids []string) error
}
type QueryService[T any] interface {
	Query(ctx context.Context) ([]*T, error)
}

func GenericDeleteManyHandler(service DeleteManyService, w http.ResponseWriter, r *http.Request) error {
	var payload []string
	err := ParseJson(r, &payload)
	if err != nil {
		return WriteApiError(w, http.StatusBadRequest, apperrors.ErrMalformedRequest.Error())
	}
	Validate.Var(&payload, "required,dive,uuid")
	err = service.DeleteMany(r.Context(), payload)
	if err != nil {
		return err
	}
	return WriteJson(w, http.StatusNoContent, nil)
}

func GenericQueryHandler[T any](service QueryService[T], w http.ResponseWriter, r *http.Request) error {
	users, err := service.Query(r.Context())
	if err != nil {
		return err
	}
	return WriteJson(w, http.StatusOK, users)
}

func GenericDeleteHandler[P any, R any, T any](service CRUDService[P, R, T], w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]
	if _, err := uuid.Parse(id); err != nil {
		return WriteApiError(w, http.StatusNotFound, apperrors.ErrNotFound.Error())
	}
	err := service.Delete(r.Context(), id)
	if err != nil {
		return err
	}
	return WriteJson(w, http.StatusNoContent, nil)
}

func GenericCreateHandler[P any, R any, T any](service CRUDService[P, R, T], payload *P, w http.ResponseWriter, r *http.Request) error {
	isValid := ParseValidateWriteIfFailPayload(w, r, payload)
	if !isValid {
		return nil
	}
	response, err := service.Create(r.Context(), payload)
	if err != nil {
		return err
	}
	return WriteJson(w, http.StatusCreated, response)
}

// parses an id to update a resources, validates, then updates it
func GenericUpdateHandler[P any, R any, T any](service CRUDService[P, R, T], payload *T, w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]
	if _, err := uuid.Parse(id); err != nil {
		return WriteApiError(w, http.StatusNotFound, apperrors.ErrNotFound.Error())
	}
	isValid := ParseValidateWriteIfFailPayload(w, r, payload)
	if !isValid {
		return nil
	}
	err := service.Update(r.Context(), id, payload)
	if err != nil {
		return err
	}
	return WriteJson(w, http.StatusNoContent, nil)
}

// parses an id into a uuid to validate, then fetches the resource
func GenericGetByIdHandler[P any, R any, T any](service CRUDService[P, R, T], w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]
	if _, err := uuid.Parse(id); err != nil {
		return WriteApiError(w, http.StatusNotFound, apperrors.ErrNotFound.Error())
	}
	response, err := service.Get(r.Context(), id)
	if err != nil {
		return err
	}
	return WriteJson(w, http.StatusOK, response)
}

// returns true if payload is parsed and validated without error.
// otherwise returns false, writes the error to the response, and returns any unexpected errors
// that typically shouldn't happen
func ParseValidateWriteIfFailPayload(w http.ResponseWriter, r *http.Request, payload any) bool {
	err := ParseJson(r, &payload)
	if err != nil {
		WriteApiError(w, http.StatusBadRequest, apperrors.ErrMalformedRequest.Error())
		return false
	}
	err = Validate.StructCtx(r.Context(), payload)
	if err != nil {
		WriteApiError(w, http.StatusBadRequest, err.Error())
		return false
	}
	return true
}
