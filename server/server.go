package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"github.com/vil-coyote-acme/go-concurrency/commons"
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
	/** TODO 1. complete the method unmarshalOrderFromHttp() (at the end), to unmarshall
	  the http entry of message**/
	unMarshallErr := unmarshalOrderFromHttp(r, &order)
	if unMarshallErr != nil {
		log.Println(unMarshallErr.Error())
		return
	}
	log.Printf("receive one order : %s", order)

	// second step, send the order to the bartender
	/** TODO 2. assign the playerId from the server to the order and marshall order to send
	to the bartender. complete the method postOrder**/
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
	/**TODO 3. Get payment information from the CallBackUrl in the order
	Complete the method*/
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
	/** TODO  1.*/
	buf := make([]byte, r.ContentLength)
	// a. use io library to read message from the body of the request
	io.ReadFull(r.Body, buf)
	// b. use json library to unmarsall the message in a "commons.Order"
	err = json.Unmarshal(buf, &order)
	// c. return
	return
}

func (s *Server) postOrder(order commons.Order) (r *http.Response, err error) {
	/** TODO 2.*/
	//a. assing playerId
	order.PlayerId = s.playerId
	//b. import and use json library to marshall order in the "buf" variable.
	buf, marshalErr := json.Marshal(order)
	//c. if error in marshalling, log it and return
	if marshalErr != nil {
		log.Println(marshalErr.Error())
		return
	}
	bartenderUrl := s.bartenderUrl + bartenderPath
	// d. use http post to send to the bartenderUrl,as json the marshalled order
	// you may want to explore la function bytes.NewBuffer
	return http.Post(bartenderUrl, "application/json", bytes.NewBuffer(buf))
}

func getDataFromCallback(order commons.Order) (err error) {
	/**TODO 3.*/
	//a. use the method Get from http to get my payment status
	paymentRes, paymentErr := http.Get(order.CallBackUrl)
	//b. if error, log it and return
	if paymentErr != nil {
		log.Printf("get an error when calling payment api : %s", paymentErr.Error())
		return paymentErr
	}
	//c. if StatusCode from the response of the server != 200, log the error and return
	if paymentRes != nil && paymentRes.StatusCode != 200 {
		log.Printf("get a non 200 response when calling payment api : %s", paymentRes.Status)
		return fmt.Errorf("get a non 200 response when calling payment api : %s", paymentRes.Status)
	}
	return nil
}
