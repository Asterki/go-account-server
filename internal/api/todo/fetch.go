package todo

import (
	"encoding/json"
	apiUtils "github.com/Asterki/go-test/internal/utils"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// Struct for the request body
type requestBody struct {
  TodoName string `json:"todo_name" validate:"required"`
}

// Response structure
type responseBody struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func FetchTodosController(w http.ResponseWriter, r *http.Request) {
	var unparsedBody requestBody
	if !apiUtils.ParseRequestBody(w, r, &unparsedBody) {
		return
	}

	if !apiUtils.ValidateRequestBody(w, unparsedBody) {
		return
	}


}
