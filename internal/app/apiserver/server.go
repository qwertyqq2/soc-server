package apiserver

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/qwertyqq2/soc-server/internal/app/apiserver/hub"
	"github.com/qwertyqq2/soc-server/internal/app/hadlers"
	"github.com/qwertyqq2/soc-server/internal/listner/provider"
)

type Server struct {
	router   *httprouter.Router
	provider *provider.Provider
	hub      *hub.Hub
}

func NewServer(r *httprouter.Router) *Server {
	hub := hub.NewHub()
	provider, err := provider.New(hub)
	if err != nil {
		log.Fatal(err)
	}
	s := &Server{
		router:   httprouter.New(),
		provider: provider,
		hub:      hub,
	}
	s.conifigureServer()
	return s
}

func (s *Server) conifigureServer() {
	h := hadlers.NewHander(s.provider, s.hub)
	h.Register(s.router)
}

func (s *Server) ListenEvents() {
	err := s.provider.ListenEvent()
	log.Fatal(err)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
