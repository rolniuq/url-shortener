package app

import (
	"context"
	"fmt"
	"net/http"
	"urlshorter/internal/config"
	"urlshorter/internal/service/shorten"
)

type Server struct {
	config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

func (s *Server) Start(ctx context.Context) {
	fmt.Println("server is running")
	mux := http.NewServeMux()

	mux.Handle("/health", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}))

	mux.Handle("/shorten", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Query().Get("url")
		if url == "" {
			http.Error(w, "Missing 'url' query parameter", http.StatusBadRequest)
			return
		}

		srv := shorten.NewShortenService(s.config)
		code, err := srv.ShortenURL(ctx, url)
		if err != nil {
			http.Error(w, "Failed to shorten URL", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(code))
	}))

	http.ListenAndServe(fmt.Sprintf(":%s", s.config.Port), mux)
}
