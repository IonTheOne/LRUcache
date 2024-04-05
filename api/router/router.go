package router

import (
    "github.com/go-chi/chi/v5"

    "github.com/Mlstermass/LRUcache/api/controller"
    "github.com/Mlstermass/LRUcache/pkg/env"
)

func New(ctl controller.App, conf env.Config) *chi.Mux {
    r := chi.NewRouter()

    r.Route("/", func(r chi.Router) {
        r.Get("/health", ctl.HealthCheck)
        r.Get("/news", ctl.GetNews)
        r.Get("/news/{newsItemId}", ctl.GetNewsByID)
    })

    return r
}