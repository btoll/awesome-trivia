package main

import (
	"flag"
	"log"
	"net/url"
	"strconv"

	"github.com/btoll/trivial"
	"golang.org/x/net/websocket"
)

var (
	serverURI       = flag.String("u", "wss://192.168.1.96:3000", "URL of game websocket server")
	gameName        = flag.String("n", "default", "Name of game")
	tokenExpiration = flag.Float64("e", 3600, "Token expiration (in seconds)")
)

func main() {
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
	}, trivial.TLSCert{
		EcdsaCurve: "P384",
		Host:       "127.0.0.1,192.168.1.96",
		IsCA:       true,
		RsaBits:    3072,
	})

	game := trivial.NewGame(*gameName, *tokenExpiration)
	server.RegisterAndStartGame(game)
}
