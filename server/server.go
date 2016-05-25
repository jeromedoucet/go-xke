package server

import (
	"net/http"
	"log"
	"github.com/vil-coyote-acme/go-concurrency/commons"
	"io"
	"encoding/json"
	"bytes"
)

const bartenderPath string = "/orders/"

func NewServer(playerId string, bartenderUrl string) (s *Server) {
	s = new(Server)
	s.bartenderUrl = bartenderUrl
	s.playerId = playerId
	s.mux = http.NewServeMux()
	s.mux.HandleFunc("/orders", s.orderHandler)//mux server. only listen on /order request !
	return s
}

type Server struct {
	playerId     string
	bartenderUrl string
	mux          *http.ServeMux
}

// start the server. Beware ! this call is blocking !
func (s *Server) Start() {
	err := http.ListenAndServe(":4242", s.mux)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func (s *Server) orderHandler(w http.ResponseWriter, r *http.Request) {
	// first step : unmarshal the incoming order
	var order commons.Order
	buf, unMarshallErr := unmarshallOrder(r, &order)
	if unMarshallErr != nil {
		log.Println(unMarshallErr.Error())
		return
	}
	log.Printf("receive one order : %s", order)

	// second step, send the order to the bartender
	res, err := http.Post(s.bartenderUrl + bartenderPath, "application/json", bytes.NewBuffer(buf))
	if err != nil {
		log.Printf("error when calling bartender api : %s", err)
		return
	}
	if res.StatusCode != 200 {
		log.Printf("get a non 200 response when calling bartender api : %s", res.Status)
		return
	}

	// third step, if all is right, get your money back !
	http.Get(order.CallBackUrl)
	w.WriteHeader(200)
}

func unmarshallOrder(r *http.Request, order *commons.Order) (buf []byte, err error) {
	buf = make([]byte, r.ContentLength)
	io.ReadFull(r.Body, buf)
	err = json.Unmarshal(buf, &order)
	return
}
