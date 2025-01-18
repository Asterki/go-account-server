// The purpose of this file is to route all of the accounts routes via a single file
package accounts

import (
	"github.com/gorilla/mux"
)

func RegisterRouter(r *mux.Router) {
  r.HandleFunc("/register", AccountsRegisterController).Methods("POST")
}

