package utils

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func ParseRequestBody(w http.ResponseWriter, r *http.Request, body interface{}) bool {
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(body); err != nil {
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Accept-Post", "application/json")
    SendResponse(w, http.StatusUnsupportedMediaType, map[string]interface{}{
      "status": "error",
      "message": "Unsupported media type",
    })
    return false
	}
	return true
}

func ValidateRequestBody(w http.ResponseWriter, body interface{}) bool {
	validate := validator.New()
	if err := validate.Struct(body); err != nil {
    errs := err.(validator.ValidationErrors)

    var validationErrors []string
    for _, e := range errs {
      validationErrors = append(validationErrors, e.Field() + " is " + e.Tag())
    }

    SendResponse(w, http.StatusBadRequest, map[string]interface{}{
      "status": "error",
      "message": "Validation error",
      "errors": validationErrors,
    })

    return false
	}
	return true
}

func SendResponse(w http.ResponseWriter, status int, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}
