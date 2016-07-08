package daemon

import (
	"log"
	"net/http"

	"github.com/HBM/go-jet/peer"

	"golang.org/x/net/websocket"
)

type Daemon struct {
	WebsocketServer websocket.Server
}

func NewDaemon() Daemon {
	return Daemon{
		WebsocketServer: websocket.Server{
			// use custom handshake to fix Origin not set when using websockets not in browser
			Handshake: func(c *websocket.Config, r *http.Request) error {
				return nil
			},
			Handler: Handler,
		},
	}
}

func Handler(ws *websocket.Conn) {

	peer := peer.NewPeer(ws)
	log.Printf("%+v", peer)

	// add peer to peers
	// peers[peer.ID] = peer

	log.Printf("%#v", ws)
	// show received message
	var message string
	if err := websocket.Message.Receive(ws, &message); err != nil {
		panic(err)
	}
	log.Println(message)

	// send message
	websocket.Message.Send(ws, "awesome")
}
