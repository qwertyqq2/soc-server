package hadlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/julienschmidt/httprouter"
	"github.com/qwertyqq2/soc-server/internal/listner/provider"
)

const (
	LoadURL = "/connect"
)

type HandlerStorage struct {
	clients map[string]bool
}

func NewHandlerStorage() *HandlerStorage {
	return &HandlerStorage{
		clients: make(map[string]bool),
	}
}

type handler struct {
	provider *provider.Provider
	storage  HandlerStorage
}

func NewHander(p *provider.Provider) *handler {
	return &handler{
		provider: p,
		storage:  *NewHandlerStorage(),
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
	h.storage.clients[senderAddr] = true
	log.Printf("connected '%s'", senderAddr)
	defer delete(h.storage.clients, senderAddr)
	if err != nil {
		log.Println("upgrade:", err)
		return
	}
	defer conn.Close()
	rounds, err := h.provider.Store.Repository().AllRounds()
	if err != nil {
		log.Fatal(err)
	}
	for _, round := range rounds {
		msg := Message{
			Type: 101,
			Data: *round,
		}
		jsr, err := json.Marshal(msg)
		if err != nil {
			log.Fatal(err)
		}
		err = conn.WriteMessage(websocket.TextMessage, jsr)
		if err != nil {
			log.Fatal(err)
		}
	}
	players, err := h.provider.Store.Repository().AllPlayers()
	if err != nil {
		log.Fatal(err)
	}
	for _, player := range players {
		msg := Message{
			Type: 102,
			Data: *player,
		}
		jpl, err := json.Marshal(msg)
		if err != nil {
			log.Fatal(err)
		}
		err = conn.WriteMessage(websocket.TextMessage, jpl)
		if err != nil {
			log.Fatal(err)
		}
	}
	lots, err := h.provider.Store.Repository().AllLots()
	if err != nil {
		log.Fatal(err)
	}
	for _, lot := range lots {
		msg := Message{
			Type: 103,
			Data: *lot,
		}
		jlot, err := json.Marshal(msg)
		if err != nil {
			log.Fatal(err)
		}
		err = conn.WriteMessage(websocket.TextMessage, jlot)
		if err != nil {
			log.Fatal(err)
		}

	}
	msg := Message{
		Type: 100,
		Data: nil,
	}
	jall, err := json.Marshal(msg)
	if err != nil {
		log.Fatal(err)
	}
	err = conn.WriteMessage(websocket.TextMessage, jall)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case lot := <-h.provider.Store.LotsChan:
			msg := Message{
				Type: 103,
				Data: *lot,
			}
			jlot, err := json.Marshal(msg)
			if err != nil {
				log.Fatal(err)
			}
			err = conn.WriteMessage(websocket.TextMessage, jlot)
			if err != nil {
				log.Fatal(err)
			}
		case player := <-h.provider.Store.PlayersChan:
			msg := Message{
				Type: 102,
				Data: *player,
			}
			jpl, err := json.Marshal(msg)
			if err != nil {
				log.Fatal(err)
			}
			err = conn.WriteMessage(websocket.TextMessage, jpl)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

}
