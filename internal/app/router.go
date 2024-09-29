package app

import (
	"net/http"

	"github.com/BrancheDeboua/url-shortener/internal/controller"
)

func router(s *controller.Shortener) *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("POST /{$}", s.HandleShorten)
	router.HandleFunc("GET /{$}", s.HandleIndex)
	router.HandleFunc("GET /404", s.HandleError)
	router.HandleFunc("GET /{id}", s.HandleRedirect)

	return router
}
