package hub

import "github.com/qwertyqq2/soc-server/internal/listner/provider/model"

type Hub struct {
	Clients map[string]bool

	Broadcasters map[string]chan *model.Resp
}

func NewHub() *Hub {
	return &Hub{
		Clients:      map[string]bool{},
		Broadcasters: map[string]chan *model.Resp{},
	}
}

func (hub *Hub) NewClient(addr string) {
	hub.Clients[addr] = true
}

func (hub *Hub) Broadcast(data *model.Resp) {
	for cli, ok := range hub.Clients {
		if ok {
			hub.Broadcasters[cli] <- data
		}
	}
}
