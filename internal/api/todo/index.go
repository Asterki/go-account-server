package todo 

import (
  "github.com/gorilla/mux"
)

func RegisterRouter(r *mux.Router) {
  r.HandleFunc("/register", AccountsRegisterController).Methods("POST")
}

