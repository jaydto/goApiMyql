package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/jaydto/goApiMyql/types"
)

var Validate=validator.New()

func ParseJson(r *http.Request, payload any) error {
	if r.Body == nil {
		return fmt.Errorf("missing request Body")

	}
	return json.NewDecoder(r.Body).Decode(payload)

}

func WriteJson(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)

}
func WriteMessage(w http.ResponseWriter, status int, v any) {
	message:=types.Message{Data: v}
	WriteJson(w, status, message)
	// WriteJson(w, status, map[string]any{"data":v})
}

func WriteError(w http.ResponseWriter, status int, err error) {
	errorMessage:=types.Error{Error: err.Error()}
	// WriteJson(w, status, map[string]string{"error": err.Error()})
	WriteJson(w, status, errorMessage)

}
