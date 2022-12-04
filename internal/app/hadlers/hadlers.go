package hadlers

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/julienschmidt/httprouter"
)

const (
	initLoadURL = "/connect"
)

type handler struct {
}

func NewHander() *handler {
	return &handler{}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET("/ws", Connect)
}

func sendStuff(ws *websocket.Conn /* stuff */) {
	ws.WriteMessage(websocket.TextMessage, []byte("qwe"))
}

var upgrader = websocket.Upgrader{}

func Connect(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade:", err)
		return
	}
	defer conn.Close()
	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = conn.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}
