package app

import (
	"context"
	"fmt"
	"net/http"
	"urlshorter/config"
)

type Server struct {
	port string
}

func NewServer(config *config.Config) *Server {
	return &Server{
		port: config.Port,
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
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Shorten URL endpoint"))
	}))
	http.ListenAndServe(fmt.Sprintf(":%s", s.port), mux)
}
