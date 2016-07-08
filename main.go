package main

import (
	"log"
	"net/http"

	"github.com/HBM/go-jet/daemon"
)

func main() {
	d := daemon.NewDaemon()
	http.Handle("/", http.Handler(d.WebsocketServer))
	log.Fatal(http.ListenAndServe(":12345", nil))
}
