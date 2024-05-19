package util

import "net/http"

func MalformedRequest(w http.ResponseWriter) error {
	return WriteJson(w, http.StatusBadRequest, ApiError{"malformed request"})
}
