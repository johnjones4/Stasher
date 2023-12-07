package api

import (
	"main/core"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
)

type API struct {
	runtime *core.RuntimeContext
	r       *chi.Mux
	log     *zap.SugaredLogger
}

func (a *API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.r.ServeHTTP(w, r)
}

func New(runtime *core.RuntimeContext, log *zap.SugaredLogger) *API {
	a := &API{
		runtime: runtime,
		r:       chi.NewRouter(),
		log:     log,
	}

	a.r.Use(middleware.RequestID)
	a.r.Use(middleware.RealIP)
	a.r.Use(middleware.Logger)
	a.r.Use(middleware.Recoverer)

	a.r.Route("/api", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})

		r.Get("/stash", a.stash)
		r.Post("/telegram", a.telegram)
	})

	return a
}
