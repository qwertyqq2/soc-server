package apiserver

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/qwertyqq2/soc-server/internal/app/hadlers"
)

type Server struct {
	router *httprouter.Router
}

func NewServer(r *httprouter.Router) *Server {
	s := &Server{
		router: httprouter.New(),
	}
	s.conifigureServer()
	return s
}

func (s *Server) conifigureServer() {
	h := hadlers.NewHander()
	h.Register(s.router)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
