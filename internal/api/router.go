package api

import (
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"

	"net/http"

  // Import the routers
	accountsRouter "github.com/Asterki/go-test/internal/api/accounts"

	log "github.com/sirupsen/logrus"
)

func Router() http.Handler {
	log.Info("Setting up the router...")

	r := mux.NewRouter()

	// Modular route registration
	accounts := r.PathPrefix("/api/accounts").Subrouter()
	accountsRouter.RegisterRouter(accounts)

	// Set up middleware with Negroni
	n := negroni.New()
  n.Use(negroni.NewRecovery())
  n.Use(negroni.NewLogger())
  n.Use(negroni.NewStatic(http.Dir("web/dist"))) // Serve the web directory
  n.UseHandler(r)

	return n
}
