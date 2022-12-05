package apiserver

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/qwertyqq2/soc-server/internal/app/hadlers"
	"github.com/qwertyqq2/soc-server/internal/listner/provider"
)

type Server struct {
	router   *httprouter.Router
	provider *provider.Provider
}

func NewServer(r *httprouter.Router, p *provider.Provider) *Server {
	s := &Server{
		router:   httprouter.New(),
		provider: p,
	}
	s.conifigureServer(p)
	return s
}

func (s *Server) conifigureServer(p *provider.Provider) {
	h := hadlers.NewHander(p)
	h.Register(s.router)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
