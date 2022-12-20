package hub

type Hub struct {
	Clients map[string]bool

	Broadcasters map[string]chan []interface{}
}

func NewHub() *Hub {
	clients := make(map[string]bool, 10)
	broadcasters := make(map[string]chan []interface{}, 10)

	return &Hub{
		Clients:      clients,
		Broadcasters: broadcasters,
	}
}

func (hub *Hub) NewClient(addr string) {
	hub.Clients[addr] = true
	hub.Broadcasters[addr] = make(chan []interface{}, 10)
}

func (hub *Hub) Broadcast(data []interface{}) {
	for cli, ok := range hub.Clients {
		if ok {
			hub.Broadcasters[cli] <- data
		}
	}
}
