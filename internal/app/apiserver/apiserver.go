package apiserver

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"github.com/qwertyqq2/soc-server/internal/configs"
	"github.com/qwertyqq2/soc-server/internal/listner/provider"
)

func Start(conf *configs.Config, p *provider.Provider) error {
	router := httprouter.New()
	srv := NewServer(router, p)

	log.Println("Listen: ", conf.Server.BindIp+":"+conf.Server.Port)
	return http.ListenAndServe(conf.Server.BindIp+":"+conf.Server.Port, srv)
}
