package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jess-monter/social/internal/store"
)

type application struct {
	config config
	store  store.Storage
}

type config struct {
	addr string
}

// We can set the return as http.Handler because chi.NewRouter() returns a value of this type.
// chi.Router also implements the http.Handler interface.
// This allows us to change the router implementation without changing the rest of the code.
func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	// This middleware will log the start and end of each request,
	// along with some useful data about what was requested and how long it took.
	// It will also recover from any panics and return a 500 error code.
	// This is a good set of middleware to include in all applications.
	// You can add more middleware to this stack as needed.
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Here we define a route group for version 1 of our API.
	// This is a good practice to allow for future versions of the API.
	// All routes for version 1 will be prefixed with /v1.
	// For example, /v1/health, /v1/endpoint, etc.
	// This allows us to maintain backward compatibility when we introduce new versions of the API.
	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.healthCheckHandler)
	})

	return r
}

func (app *application) run(mux http.Handler) error {

	server := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Printf("Starting server on %s", app.config.addr)

	return server.ListenAndServe()
}
