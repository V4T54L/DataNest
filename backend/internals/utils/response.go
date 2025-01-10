package utils

import (
	"backend/internals/schemas"
	"encoding/json"
	"net/http"
)

func jsonResponse(w http.ResponseWriter, statusCode int, obj schemas.Response) error {
	data, err := json.Marshal(obj)
	if err != nil {
		return err
	}

	w.WriteHeader(statusCode)
	_, err = w.Write(data)

	return err
}

func ErrorResponse(w http.ResponseWriter, statusCode int, error string) error {
	return jsonResponse(w, statusCode, schemas.Response{
		Success: false,
		Error:   error,
	})
}

func MessageResponse(w http.ResponseWriter, statusCode int, message string) error {
	return jsonResponse(w, statusCode, schemas.Response{
		Success: true,
		Message: message,
	})
}

func DataResponse(w http.ResponseWriter, statusCode int, data any) error {
	return jsonResponse(w, statusCode, schemas.Response{
		Success: true,
		Data:    data,
	})
}
