package rest

import (
	"github.com/error2215/url_shortener/server/config"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	log "github.com/sirupsen/logrus"

	"net/http"
)

type Server struct{}

func (s *Server) Start() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Post("/url_shortener", s.CreateURLHandler)
		})
	})

	r.Get("/rd/{url}", s.RedirectHandler)

	log.Info("Api Server started on port: " + config.GlobalConfig.ApiPort)
	err := http.ListenAndServe(":"+config.GlobalConfig.ApiPort, r)
	if err != nil {
		log.WithField("ERROR", "Cannot start API Server").Fatal(err)
	}
}
