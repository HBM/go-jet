package peer

import "golang.org/x/net/websocket"

type Peer struct {
	ID       string
	Fetchers map[string]*websocket.Conn
	WS       *websocket.Conn
}

func NewPeer(ws *websocket.Conn) Peer {
	return Peer{
		ID:       ws.Request().RemoteAddr,
		Fetchers: make(map[string]*websocket.Conn),
		WS:       ws,
	}
}

func (p *Peer) sendMessage(message string) {
	websocket.Message.Send(p.WS, "awesome")
}

func (p *Peer) addFetcher(id string, ws *websocket.Conn) {
	p.Fetchers[id] = ws
}

func (p *Peer) removeFetcher(id string) {
	delete(p.Fetchers, id)
}

func (p *Peer) hasFetchers() bool {
	return len(p.Fetchers) != 0
}
