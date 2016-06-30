package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/vil-coyote-acme/go-concurrency/commons"
	"io"
	"encoding/json"
)

const bartenderPath string = "/orders"

type Server struct {
	playerId     string
	bartenderUrl string
	mux          *http.ServeMux
}

func (s *Server) handleOrder(w http.ResponseWriter, r *http.Request) {
	// first step : unmarshal the incoming order
	var order commons.Order
	unMarshallErr := unmarshalOrderFromHttp(r, &order)
	if unMarshallErr != nil {
		log.Println(unMarshallErr.Error())
		return
	}
	log.Printf("receive one order : %s", order)

	// second step, send the order to the bartender
	// TODO : complete the postOrder function
	res, err := s.postOrder(order)

	if err != nil {
		log.Printf("error when calling bartender api : %s", err)
		return
	}
	if res.StatusCode != 200 {
		log.Printf("get a non 200 response when calling bartender api : %s", res.Status)
		return
	}

	// third step, if all is right, get your money back !
	//TODO : complete the getDataFromCallback function
	getDataFromCallback(order)
	w.WriteHeader(200)
}

func (s *Server) Start(url string) {
	err := http.ListenAndServe(url, s.mux)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func NewServer(playerId string, bartenderUrl string) (s *Server) {
	s = new(Server)
	s.bartenderUrl = bartenderUrl
	s.playerId = playerId
	s.mux = http.NewServeMux()
	s.mux.HandleFunc("/orders", s.handleOrder)
	return s
}

func unmarshalOrderFromHttp(r *http.Request, order *commons.Order) (err error) {
	buf := make([]byte, r.ContentLength)
	io.ReadFull(r.Body, buf)
	err = json.Unmarshal(buf, &order)
	return
}

func (s *Server) postOrder(order commons.Order) (r *http.Response, err error) {
	order.PlayerId = s.playerId
	buf, _ := json.Marshal(order)
	fmt.Println(buf)
	bartenderUrl := s.bartenderUrl + bartenderPath
	fmt.Println(bartenderUrl)
	// d. use http post to send to the bartenderUrl, as application/json, the marshalled order
	// you may want to explore la function bytes.NewBuffer. Return the response and the error if it exists
	return nil, nil
}

func getDataFromCallback(order commons.Order) (err error) {
	/**TODO 3.*/
	//a. use the method Get from http to get my payment status

	//b. if error, log it and return

	//c. if StatusCode from the response of the server != 200, log it  and return an error
	// hint: explore fmt.Errorf

	return nil
}
