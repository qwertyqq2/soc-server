package hadlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/julienschmidt/httprouter"
	"github.com/qwertyqq2/soc-server/internal/app/apiserver/hub"
	"github.com/qwertyqq2/soc-server/internal/listner/provider"
)

const (
	LoadURL = "/connect"
)

type handler struct {
	provider *provider.Provider
	hub      *hub.Hub
}

func NewHander(p *provider.Provider, hub *hub.Hub) *handler {
	return &handler{
		provider: p,
		hub:      hub,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET(LoadURL, h.Connect)
}

func sendStuff(ws *websocket.Conn) {
	ws.WriteMessage(websocket.TextMessage, []byte("qwe"))
}

var upgrader = websocket.Upgrader{}

type Message struct {
	Type uint        `json:"type"`
	Data interface{} `json:"data"`
}

func (h *handler) Connect(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	conn, err := upgrader.Upgrade(w, r, nil)
	senderAddr := r.URL.Query()["addr"][0]
	h.hub.NewClient(senderAddr)
	log.Printf("connected '%s'", senderAddr)
	if err != nil {
		log.Println("upgrade:", err)
		return
	}
	defer conn.Close()
	doneChan := make(chan bool)
	//doneInit := make(chan bool)
	go h.initLoad(conn, senderAddr, doneChan)
	go h.listenHub(conn, senderAddr, doneChan)
	for {
	}
}

func (h *handler) initLoad(conn *websocket.Conn, senderAddr string, doneChan chan bool) {
	resp, err := h.provider.InitData()
	if err != nil {
		log.Println(err)
	}
	jsonData, err := resp.Marshal()
	if err != nil {
		log.Println(err)
	}
	err = conn.WriteMessage(websocket.TextMessage, jsonData)
	if err != nil {
		log.Println(err)
	}
	doneChan <- true
}

func (h *handler) listenHub(conn *websocket.Conn, senderAddr string, doneChan chan bool) {
	select {
	case ok := <-doneChan:
		if ok {
			log.Println("Okey")
			for {
				select {
				case data := <-h.hub.Broadcasters[senderAddr]:
					jsonData, err := json.Marshal(data)
					if err != nil {
						return
					}
					err = conn.WriteMessage(websocket.TextMessage, jsonData)
					if err != nil {
						return
					}
				}
			}
		}
	}
}
