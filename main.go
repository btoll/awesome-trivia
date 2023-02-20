package main

import (
	"flag"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/btoll/trivial"
	"golang.org/x/net/websocket"
)

func main() {
	serverURI := flag.String("u", "ws://192.168.1.96:3000", "URL of game websocket server")
	gameName := flag.String("n", "default", "Name of game")
	tokenExpiration := flag.Float64("e", 3600, "Token expiration (in seconds)")
	keyLength := flag.Int("k", 25, "Key length")
	flag.Parse()

	parsedUrl, err := url.Parse(*serverURI)
	if err != nil {
		log.Fatalln("game server uri could not be parsed")
	}
	port, err := strconv.Atoi(parsedUrl.Port())
	if err != nil {
		log.Fatalln("port could not be parsed")
	}
	server := trivial.NewSocketServer(trivial.URI{
		Sock: trivial.Socket{
			Protocol: parsedUrl.Scheme,
			Domain:   parsedUrl.Hostname(),
			Port:     port,
		},
		Path: "ws",
	})

	game := trivial.NewGame(*gameName, *keyLength, *tokenExpiration)
	server.RegisterGame(game)

	http.Handle("/ws", websocket.Handler(server.DefaultHandler))
	http.HandleFunc("/", server.BaseHandler)
	http.HandleFunc("/kill", server.KillHandler)
	http.HandleFunc("/message", server.MessageHandler)
	http.HandleFunc("/notify", server.NotifyHandler)
	http.HandleFunc("/query", server.QueryHandler)
	http.HandleFunc("/reset", server.ResetHandler)
	http.HandleFunc("/scoreboard", server.ScoreboardHandler)
	http.ListenAndServe(":3000", nil)
}
