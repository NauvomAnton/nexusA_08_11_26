package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/iliaatmazhitov/nexus/internal/httpx"
)

func handleNotFound(w http.ResponseWriter, r *http.Request) {
	httpx.WriteError(w, http.StatusNotFound, "not_found", "route not found")
}

func handleMethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	httpx.WriteError(w, http.StatusMethodNotAllowed, "method_not_allowed", "method not allowed")
}

// mock func
func handleMe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("WWW-Authenticate", "Bearer")
	httpx.WriteError(w, http.StatusUnauthorized, "unauthenticated", "missing or invalid token")
}

// TODO: насчет RealIP надо подумать и чем нибудь его заменить, он оказывается уже успел устареть

// Собирает http роутинг приложения
func NewRouter() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RequestID, middleware.RealIP, middleware.Recoverer)

	r.NotFound(handleNotFound)
	r.MethodNotAllowed(handleMethodNotAllowed)

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/me", handleMe) // mock 
	})
	
	return r
}
