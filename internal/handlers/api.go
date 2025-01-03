package handlers

import (
	"nilomiranda/web-alerts/internal/routes"

	"github.com/go-chi/chi"
	// middleware "github.com/go-chi/chi/middleware"
	// "danilo/web-alerts/internal/middlware"
)

func Handler(r *chi.Mux) {
  // r.Use(middleware.StripSlashes)
  
  r.Route("/api", func(r chi.Router) {
    r.Get("/health", routes.GetHealth)
  })
}
