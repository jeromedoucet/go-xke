package main

import (
	"flag"
	"github.com/vil-coyote-acme/go-xke/registration"
	"github.com/vil-coyote-acme/go-xke/server"
	"log"
	"fmt"
)

var (
	clientIp    string
	ourIp       string
	bartenderIp string
	playerId    string
)

func main() {
	flag.StringVar(&clientIp, "clientIp", "127.0.0.1:4444", "ip to the client component. Used for registration")
	flag.StringVar(&ourIp, "ourIp", "127.0.0.1:4242", "ip that must be used by the client to call us")
	flag.StringVar(&bartenderIp, "bartenderIp", "127.0.0.1:4343", "ip to the bartender component")
	flag.StringVar(&playerId, "playerId", "playerId", "our id !")
	flag.Parse()
	registrationErr := registration.Register("http://"+clientIp, "http://"+ourIp, playerId)
	if registrationErr == nil {
		server.NewServer(playerId, "http://"+bartenderIp).Start()
	} else {
		log.Panic(fmt.Sprintf("Erreur when trying to register the server : %s", registrationErr.Error()))
	}
}
