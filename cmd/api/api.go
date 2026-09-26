package main

import (
	"gopher/internal/store"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type application struct {
	config config
	store  store.Storage
}

type config struct {
	addr string
	db   dbConfig
	env  string
}
type dbConfig struct {
	addr         string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  string
}

func (app *application) mount() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.ClientIPFromRemoteAddr) // pick one ClientIPFrom* based on your infra, see below
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.healthChecker)
		r.Route("/posts", func(r chi.Router) {
			r.Post("/", app.createPostHanlder)
			r.Route("/{postID}", func(r chi.Router) {
				r.Get("/", app.getPostHandler)
			})
		})
	})
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	r.Get("/health", app.healthChecker)
	return r
}
func (app *application) run(mux *chi.Mux) error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}
	log.Printf("the server running at %s", app.config.addr)
	return srv.ListenAndServe()

}
