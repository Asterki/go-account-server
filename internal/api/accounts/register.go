package accounts

import (
	"encoding/json"
	apiUtils "github.com/Asterki/go-test/internal/utils"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// Struct for the request body
type requestBody struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

// Response structure
type responseBody struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func AccountsRegisterController(w http.ResponseWriter, r *http.Request) {
	var unparsedBody requestBody
	if !apiUtils.ParseRequestBody(w, r, &unparsedBody) {
		return
	}

	if !apiUtils.ValidateRequestBody(w, unparsedBody) {
		return
	}

	// Process the parsed body (log the values)
	log.Infof("Email: %s", unparsedBody.Email)
	log.Infof("Password: %s", unparsedBody.Password)

	// Respond with a successful message
	successfulResponse := responseBody{
		Status:  "success",
		Message: "Account created",
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(successfulResponse)
}
