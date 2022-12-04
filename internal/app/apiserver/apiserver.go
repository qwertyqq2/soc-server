package apiserver

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"github.com/qwertyqq2/soc-server/internal/configs"
)

func Start(conf *configs.Config) error {
	router := httprouter.New()
	srv := NewServer(router)

	log.Println("listen...")
	return http.ListenAndServe(conf.Server.BindIp+":"+conf.Server.Port, srv)
}
